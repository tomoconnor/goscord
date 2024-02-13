package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type GuildStickersUpdate struct {
	Data *discord.GuildStickersUpdateEventFields `json:"d"`
}

func NewGuildStickersUpdate(rest *rest.Client, data []byte) (*GuildStickersUpdate, error) {
	pk := new(GuildStickersUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
