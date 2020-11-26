package main

import (
	"fmt"

	"nanago/nana"
)

const (
	// YourUserID your Nana ID
	YourUserID = "XXXXXXXXXX"
)

func main() {
	client, err := nana.CreateAccount("", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[ID]", client.ID, "[NAME]", client.Name, "[TOKEN]", client.Token)
	if _, err := client.Follow(YourUserID); err != nil {
		fmt.Println(err)
		return
	}
}
