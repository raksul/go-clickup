// The simple command demonstrates the functionality that
// prompts the user for a Clickup goal and lists all the entities
// that are related to the specified goal.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/raksul/go-clickup/clickup"
)

func fetchGoals(teamId string) ([]clickup.Goal, []clickup.GoalFolder, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	goals, folders, _, err := client.Goals.GetGoals(context.Background(), teamId, false)
	return goals, folders, err
}

func main() {
	var teamId string
	fmt.Print("Enter clickup teamId: ")
	fmt.Scanf("%s", &teamId)

	goals, _, err := fetchGoals(teamId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, goal := range goals {
		fmt.Println(goal.Name)
	}
}
