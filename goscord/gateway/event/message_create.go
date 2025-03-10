package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type MessageCreate struct {
	Data *discord.Message `json:"d"`
}

func NewMessageCreate(rest *rest.Client, data []byte) (*MessageCreate, error) {
	pk := new(MessageCreate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
