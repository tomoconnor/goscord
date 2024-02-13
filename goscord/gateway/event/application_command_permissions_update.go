package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type ApplicationCommandPermissionsUpdate struct {
	Data *discord.GuildApplicationCommandPermissions `json:"d"`
}

func NewApplicationCommandPermissionsUpdate(rest *rest.Client, data []byte) (*ApplicationCommandPermissionsUpdate, error) {
	pk := new(ApplicationCommandPermissionsUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
