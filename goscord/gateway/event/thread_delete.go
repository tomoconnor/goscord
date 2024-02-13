package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type ThreadDelete struct {
	Data *discord.Channel `json:"d"`
}

func NewThreadDelete(rest *rest.Client, data []byte) (*ThreadDelete, error) {
	pk := new(ThreadDelete)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
