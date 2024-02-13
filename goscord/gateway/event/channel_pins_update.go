package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type ChannelPinsUpdate struct {
	Data *discord.ChannelPinsUpdateEventFields `json:"d"`
}

func NewChannelPinsUpdate(rest *rest.Client, data []byte) (*ChannelPinsUpdate, error) {
	pk := new(ChannelPinsUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
