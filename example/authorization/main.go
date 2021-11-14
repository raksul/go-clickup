// The simple command demonstrates the functionality that
// prompts the user for a Clickup authorization and lists all the entities
// that are related to the specified authorization.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/raksul/go-clickup/clickup"
)

func fetchAuthorizedUser() (clickup.User, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	user, _, err := client.Authorization.GetAuthorizedUser(context.Background())
	return *user, err
}

func main() {
	user, err := fetchAuthorizedUser()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(user.Username)
}
