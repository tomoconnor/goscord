package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type ChannelUpdate struct {
	Data *discord.Channel `json:"d"`
}

func NewChannelUpdate(rest *rest.Client, data []byte) (*ChannelUpdate, error) {
	pk := new(ChannelUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
