// The simple command demonstrates the functionality that
// prompts the user for a Clickup team and lists all the entities
// that are related to the specified team.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/raksul/go-clickup/clickup"
)

func fetchTeams() ([]clickup.Team, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	teams, _, err := client.Teams.GetTeams(context.Background())
	return teams, err
}

func main() {
	teams, err := fetchTeams()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, team := range teams {
		fmt.Println(team.Name)
	}
}
