package wandel

import (
	"context"
	"fmt"
)

type menuNutritionFacts struct {
	KCalPerServing int16 `json:"KCalPerServing"`
}

type menuPrice struct {
	Price float32 `json:"Price"`
}

type menuProductFull struct {
	Name           string `json:"Name"`
	Description    string `json:"Description,omitempty"`
	Prices         []menuPrice
	NutritionFacts menuNutritionFacts
}

type menuCategory struct {
	Products []menuProductFull
}

type menuContent struct {
	Categories []menuCategory
}

type menuResponseFull struct {
	Content menuContent
}

type MenuProduct struct {
	Name        string
	Description string
	Price       string
	KCal        string
}

func (client *Client) GetTodayMenu() *[]MenuProduct {
	return client.GetTodayMenuContext(context.Background())
}

func (client *Client) GetTodayMenuContext(ctx context.Context) *[]MenuProduct {
	fullResponse := &menuResponseFull{}
	headers := headersAuthApp(client.app, client.authorization)
	url := client.endpoint + "Menus/" + client.restaraunt

	err := getResource(ctx, client.httpclient, headers, url, fullResponse, client)

	if err != nil {
		fmt.Println(err)
	}

	categories := fullResponse.Content.Categories

	res := make([]MenuProduct, len(categories))

	for i, category := range categories {
		for _, product := range category.Products {
			res[i] = MenuProduct{
				Name:        product.Name,
				Description: product.Description,
				Price:       fmt.Sprintf("%.2f", product.Prices[0].Price),
				KCal:        fmt.Sprintf("%d", product.NutritionFacts.KCalPerServing),
			}
		}
	}

	return &res
}
