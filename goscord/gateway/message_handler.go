package gateway

import "github.com/tomoconnor/goscord/goscord/gateway/event"

type MessageCreateHandler struct{}

func (_ *MessageCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageCreate(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventMessageCreate, ev.Data)
}
