package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type GuildCreate struct {
	Data *discord.Guild `json:"d"`
}

func NewGuildCreate(rest *rest.Client, data []byte) (*GuildCreate, error) {
	pk := new(GuildCreate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
