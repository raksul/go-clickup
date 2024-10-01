// The simple command demonstrates the functionality that
// prompts the user for a Clickup list and lists which are not belonging to any folder inside a space.
// that are related to the specified space.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/raksul/go-clickup/clickup"
)

func fetchLists(spaceId string) ([]clickup.List, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	lists, _, err := client.Lists.GetFolderlessLists(context.Background(), spaceId, false)
	return lists, err
}

func main() {
	var spaceId string
	fmt.Print("Enter clickup spaceId: ")
	fmt.Scanf("%s", &spaceId)

	lists, err := fetchLists(spaceId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, list := range lists {
		fmt.Println(list.Name)
	}
}
