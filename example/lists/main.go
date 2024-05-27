// The simple command demonstrates the functionality that
// prompts the user for a Clickup list and lists all the entities
// that are related to the specified folder.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gassara-kys/go-clickup/clickup"
)

func fetchLists(folderId string) ([]clickup.List, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	lists, _, err := client.Lists.GetLists(context.Background(), folderId, false)
	return lists, err
}

func main() {
	var folderId string
	fmt.Print("Enter clickup folderId: ")
	fmt.Scanf("%s", &folderId)

	lists, err := fetchLists(folderId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, list := range lists {
		fmt.Println(list.Name)
	}
}
