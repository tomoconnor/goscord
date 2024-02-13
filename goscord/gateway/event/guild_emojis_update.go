package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type GuildEmojisUpdate struct {
	Data *discord.GuildEmojisUpdateEventFields `json:"d"`
}

func NewGuildEmojisUpdate(rest *rest.Client, data []byte) (*GuildEmojisUpdate, error) {
	pk := new(GuildEmojisUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
