package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type GuildIntegrationsUpdate struct {
	Data *discord.GuildIntegrationsUpdateEventFields `json:"d"`
}

func NewGuildIntegrationsUpdate(rest *rest.Client, data []byte) (*GuildIntegrationsUpdate, error) {
	pk := new(GuildIntegrationsUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
