package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type GuildMemberAdd struct {
	Data *discord.GuildMember `json:"d"`
}

func NewGuildMemberAdd(rest *rest.Client, data []byte) (*GuildMemberAdd, error) {
	pk := new(GuildMemberAdd)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
