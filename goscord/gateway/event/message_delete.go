package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type MessageDelete struct {
	Data struct {
		Id        string `json:"id"`
		ChannelId string `json:"channel_id"`
		GuildId   string `json:"guild_id"`
	} `json:"d"`
}

func NewMessageDelete(_ *rest.Client, data []byte) (*MessageDelete, error) {
	pk := new(MessageDelete)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
