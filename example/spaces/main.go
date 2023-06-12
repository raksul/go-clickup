// The simple command demonstrates the functionality that
// prompts the user for a Clickup space and lists all the entities
// that are related to the specified space.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/raksul/go-clickup/clickup"
)

func fetchSpaces(teamId string) ([]clickup.Space, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	spaces, _, err := client.Spaces.GetSpaces(context.Background(), teamId, false)
	return spaces, err
}

func main() {
	var teamId string
	fmt.Print("Enter clickup teamId: ")
	fmt.Scanf("%s", &teamId)

	spaces, err := fetchSpaces(teamId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, space := range spaces {
		fmt.Println(space.Name)
	}
}
