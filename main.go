package main

import (
	"fmt"

	"github.com/BSick7/okta-sdk-go/api"
)

func main() {
	session := api.DefaultSession()
	client := api.NewClient(session)

	apps, err := client.Apps().List()
	fmt.Println(apps, err)
}
