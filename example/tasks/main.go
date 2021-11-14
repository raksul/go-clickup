// The simple command demonstrates the functionality that
// prompts the user for a Clickup task and lists all the entities
// that are related to the specified task.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/raksul/go-clickup/clickup"
)

func fetchTasks(listId string, options *clickup.GetTasksOptions) ([]clickup.Task, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	tasks, _, err := client.Tasks.GetTasks(context.Background(), listId, options)
	return tasks, err
}

func main() {
	var listId string
	fmt.Print("Enter clickup listId: ")
	fmt.Scanf("%s", &listId)

	options := clickup.GetTasksOptions{Statuses: []string{"to do", "in progress"}}
	tasks, err := fetchTasks(listId, &options)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, task := range tasks {
		fmt.Println(task.Name)
	}
}
