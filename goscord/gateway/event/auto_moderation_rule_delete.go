package event

import (
	"github.com/bytedance/sonic"
	"github.com/tomoconnor/goscord/goscord/discord"
	"github.com/tomoconnor/goscord/goscord/rest"
)

type AutoModerationRuleDelete struct {
	Data *discord.AutoModerationRule `json:"d"`
}

func NewAutoModerationRuleDelete(rest *rest.Client, data []byte) (*AutoModerationRuleDelete, error) {
	pk := new(AutoModerationRuleDelete)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
