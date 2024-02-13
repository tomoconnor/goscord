package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type AutoModerationRuleUpdate struct {
	Data *discord.AutoModerationRule `json:"d"`
}

func NewAutoModerationRuleUpdate(rest *rest.Client, data []byte) (*AutoModerationRuleUpdate, error) {
	pk := new(AutoModerationRuleUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
