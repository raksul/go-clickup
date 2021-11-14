// The simple command demonstrates the functionality that
// prompts the user for a Clickup tag and lists all the entities
// that are related to the specified tag.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/raksul/go-clickup/clickup"
)

func fetchTags(spaceId string) ([]clickup.Tag, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	tags, _, err := client.Tags.GetTags(context.Background(), spaceId)
	return tags, err
}

func main() {
	var spaceId string
	fmt.Print("Enter clickup spaceId: ")
	fmt.Scanf("%s", &spaceId)

	tags, err := fetchTags(spaceId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, tag := range tags {
		fmt.Println(tag.Name)
	}
}
