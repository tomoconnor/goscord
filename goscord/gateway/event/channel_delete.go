package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type ChannelDelete struct {
	Data *discord.Channel `json:"d"`
}

func NewChannelDelete(rest *rest.Client, data []byte) (*ChannelDelete, error) {
	pk := new(ChannelDelete)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
