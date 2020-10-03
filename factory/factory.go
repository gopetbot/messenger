package factory

import (
	"errors"
	"github.com/gopetbot/messenger/webhook"
)

const (
	facebook = "facebook"
	telegram = "telegram"
)

func SamplePlatform(platform string) (webhook.Platform, error) {
	switch platform {

	case facebook:
		return webhook.NewFacebook(), nil
	case telegram:
		return webhook.NewTelegram(), nil
	}
	return nil, errors.New("was not possible found platform")

}
