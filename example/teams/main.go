// The simple command demonstrates the functionality that
// prompts the user for a Clickup team and lists all the entities
// that are related to the specified team.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gassara-kys/go-clickup/clickup"
)

func fetchTeams() ([]clickup.Team, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	teams, _, err := client.Teams.GetTeams(context.Background())
	return teams, err
}

func fetchSeats(teamId string) (clickup.Seats, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	seats, _, err := client.Teams.GetSeats(context.Background(), teamId)
	return seats, err
}

func fetchPlan(teamId string) (clickup.Plan, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	plan, _, err := client.Teams.GetPlan(context.Background(), teamId)
	return plan, err
}

func main() {
	var teamId string
	fmt.Print("Enter clickup workspaceId (previously known as a teamId): ")
	fmt.Scanf("%s", &teamId)
	teams, err := fetchTeams()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for _, team := range teams {
			fmt.Println(team.Name)
		}
	}

	seats, err := fetchSeats(teamId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Seats: %+v\n", seats)
	}

	plan, err := fetchPlan(teamId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Plan: %+v\n", plan)
	}
}
