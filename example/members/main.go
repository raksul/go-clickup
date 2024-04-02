// The simple command demonstrates the functionality that
// prompts the user for a Clickup member and lists all the entities
// that are related to the specified member.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dtylman/go-clickup/clickup"
)

func fetchListMembers(listId string) ([]clickup.Member, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	members, _, err := client.Members.GetListMembers(context.Background(), listId)
	return members, err
}

func main() {
	var listId string
	fmt.Print("Enter clickup listId: ")
	fmt.Scanf("%s", &listId)

	members, err := fetchListMembers(listId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, member := range members {
		fmt.Println(member.Username)
	}
}
