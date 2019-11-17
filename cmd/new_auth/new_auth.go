package main

import (
	"fmt"

	"github.com/zloesabo/kantinebot/wandel"
)

func main() {
	// client := wandel.NewClient(wandel.OptionDebug(true))
	client := wandel.NewClient()

	authorization := client.NewAuthorization()

	fmt.Printf("Authorization: %s\n", authorization)
}
