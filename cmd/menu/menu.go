package main

import (
	"fmt"
	"os"

	"github.com/zloesabo/kantinebot/wandel"
)

func main() {
	// client := wandel.NewClient(wandel.OptionDebug(true))
	client := wandel.NewClient(
		wandel.OptionDebug(true),
		wandel.OptionAuthorization(os.Getenv("AUTHORIZATION")),
	)

	products := client.GetTodayMenu()

	for _, product := range *products {
		fmt.Printf("Product: %s\nDescription: %s\nPrice: %s KCal: %s\n\n", product.Name, product.Description, product.Price, product.KCal)
	}
}
