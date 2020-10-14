package yalis

import (
	"fmt"
	"sync"
	"time"
	"github.com/Seyz123/yalis/rest"
	"github.com/Seyz123/yalis/ws/packet"
	"github.com/Seyz123/yalis/ws/handler"
	"github.com/Seyz123/yalis/ws/event"
	"github.com/gorilla/websocket"
	ev "github.com/asaskevich/EventBus"
)

type Session struct {
	sync.Mutex
	client *Client
	bus *ev.EventBus
	connMu sync.Mutex
	conn *websocket.Conn
	sessionID string
	heartbeatInterval time.Duration
	lastSequence int
	handlers map[string]handler.EventHandler
	close chan bool
}

func NewSession(client *Client) *Session {
	s := &Session{}

	s.client = client
	s.bus = client.Bus()
	s.close = make(chan bool)
	
	s.registerHandlers()

	return s
}

func (s *Session) registerHandlers() {
    s.handlers = map[string]handler.EventHandler{
        event.EventReady: handler.NewReady(s.bus),
    }
}

func (s *Session) Login() error {
    s.Lock()
	defer s.Unlock()
	
    conn, _, err := websocket.DefaultDialer.Dial(rest.GatewayUrl, nil)
	if err != nil {
		return err
	}
	
	conn.SetCloseHandler(func (code int, text string) error {
	    return nil
	})

	s.conn = conn
	
	go func() {
		for {
			select {
			case <-s.close:
				return
			break

			default:
				_, msg, err := s.conn.ReadMessage()

				if err != nil {
					return
				}

				s.onMessage(msg)
			break
			}
		}
	}()

	return nil
}

func (s *Session) onMessage(msg []byte) {
	pk, err := packet.NewPacket(msg)

	if err != nil {
		panic(err)
	}

	opcode, event := pk.Opcode, pk.Event
	
	s.Lock()
	s.lastSequence = pk.Sequence
	s.Unlock()

	switch opcode {
	case packet.OpHello:
		hello, err := packet.NewHello(msg)

		if err != nil {
			panic(err)
		}
		
		s.Lock()
		s.heartbeatInterval = hello.Data.HeartbeatInterval
		s.Unlock()
		
		go s.startHeartbeat()
		
		identify := packet.NewIdentify(s.client.Token())
		
		if err := s.Send(identify); err != nil {
		    panic("Cannot identify")
		}

		break
	}

	if event != "" {
		handler, exists := s.handlers[event]
		
		if exists {
		    handler.Handle(msg)
		} else {
		    fmt.Println("Unhandled event : " + event)
		}
	}
}

func (s *Session) startHeartbeat() {
    for {
        s.Lock()
        ticker := time.NewTicker(s.heartbeatInterval)
        s.Unlock()
        
        defer ticker.Stop()
        
        heartbeat := packet.NewHeartbeat(s.lastSequence)
        
        err := s.Send(heartbeat)
        
        if err != nil {
            // ToDo : Try resume session
        }
        
        select {
            case <-ticker.C:
                // loop
            break
            
            case <-s.close:
                return
            break
        }
    }
}

func (s *Session) Send(v interface{}) error {
    s.connMu.Lock()
    defer s.connMu.Unlock()
    
    return s.conn.WriteJSON(v)
}

func (s *Session) Close() {
	s.conn.Close()
	s.close <- true

	fmt.Println("Connection closed")
}