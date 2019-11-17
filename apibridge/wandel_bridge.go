package apibridge

import (
	"github.com/zloesabo/kantinebot"
	"github.com/zloesabo/kantinebot/wandel"
)

type WandelBridge struct {
	client *wandel.Client
}

func (b *WandelBridge) Can(_ string) bool {
	return true
}

func (b *WandelBridge) Menu() *[]kantinebot.BotProduct {
	products := b.client.GetTodayMenu()

	res := make([]kantinebot.BotProduct, len(*products))
	for i, product := range *products {
		res[i] = kantinebot.BotProduct{
			Name: product.Name,
		}
	}

	return &res
}

func NewWandelBridge(client *wandel.Client) *WandelBridge {
	return &WandelBridge{
		client,
	}
}
