// The simple command demonstrates the functionality that
// prompts the user for a Clickup User Groups and lists all the entities
// that are related to the specified group.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gassara-kys/go-clickup/clickup"
)

func fetchUserGroups() ([]clickup.UserGroup, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	opts := &clickup.GetUserGroupsOptions{
		TeamID:   "123",
		GroupIDs: []string{"321", "456"}, // optional parameter
	}
	groups, _, err := client.UserGroups.GetUserGroups(context.Background(), opts)

	return groups, err
}

func main() {
	groups, err := fetchUserGroups()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, group := range groups {
		fmt.Println(group.Name)
	}
}
