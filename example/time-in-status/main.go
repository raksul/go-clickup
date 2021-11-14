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

func fetchTimeInStatus(taskId string, options *clickup.GetTaskOptions) (*clickup.TasksInStatus, error) {
	api_key := os.Getenv("CLICKUP_API_KEY")
	client := clickup.NewClient(nil, api_key)

	timeInStatus, _, err := client.Tasks.GetTasksTimeInStatus(context.Background(), taskId, options)
	return timeInStatus, err
}

func main() {
	var taskId string
	fmt.Print("Enter clickup TaskId: ")
	fmt.Scanf("%s", &taskId)

	options := clickup.GetTaskOptions{}
	timeInStatus, err := fetchTimeInStatus(taskId, &options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, v := range timeInStatus.StatusHistory {
		fmt.Println(v.Status, v.TotalTime.ByMinute)
	}
}
