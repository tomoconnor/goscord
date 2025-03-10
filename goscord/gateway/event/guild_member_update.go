package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type GuildMemberUpdate struct {
	Data *discord.GuildMember `json:"d"`
}

func NewGuildMemberUpdate(rest *rest.Client, data []byte) (*GuildMemberUpdate, error) {
	pk := new(GuildMemberUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
