// The simple command demonstrates the functionality that
// prompts the user for a Clickup folder and lists all the entities
// that are related to the specified space.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gassara-kys/go-clickup/clickup"
)

func fetchFolders(spaceId string) ([]clickup.Folder, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	folders, _, err := client.Folders.GetFolders(context.Background(), spaceId, false)
	return folders, err
}

func main() {
	var spaceId string
	fmt.Print("Enter clickup spaceId: ")
	fmt.Scanf("%s", &spaceId)

	folders, err := fetchFolders(spaceId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, folder := range folders {
		fmt.Println(folder.Name)
	}
}
