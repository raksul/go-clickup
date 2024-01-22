// The simple command demonstrates the functionality that
// prompts the user for a Clickup User Groups and lists all the entities
// that are related to the specified group.
package main

import (
	"context"
	"fmt"

	"github.com/raksul/go-clickup/clickup"
)

func fetchUserGroups() ([]clickup.UserGroup, error) {
	api_key := "CLICKUP_API_KEY"
	team_id := "TEAM_ID"
	client := clickup.NewClient(nil, api_key)

	opts := &clickup.GetUserGroupsOptions{
		TeamID:   team_id,
		GroupIDs: []string{"GROUP_ID"}, // optional parameter
	}
	groups, _, err := client.UserGroups.GetUserGroups(context.Background(), opts)

	return groups, err
}

func createUserGroup() (*clickup.UserGroup, error) {
	api_key := "CLICKUP_API_KEY"
	team_id := "TEAM_ID"
	member_id_to_add := 0
	client := clickup.NewClient(nil, api_key)

	opts := clickup.CreateUserGroupRequest{
		Name:    "test",
		Members: []int{member_id_to_add}, // optional parameter
	}
	group, _, err := client.UserGroups.CreateUserGroup(context.Background(), team_id, &opts)

	return group, err
}

func updateUserGroup() (*clickup.UserGroup, error) {
	api_key := "CLICKUP_API_KEY"
	group_id := "GROUP_ID"
	client := clickup.NewClient(nil, api_key)

	opts := clickup.UpdateUserGroupRequest{
		// Name: "new name", // optional parameter
		Members: clickup.UpdateUserGroupMember{
			Add: []int{0}, // optional parameter
			// Remove: []int{0}, // optional parameter
		},
	}
	group, response, err := client.UserGroups.UpdateUserGroup(context.Background(), group_id, &opts)
	fmt.Println(response)
	return group, err
}

func deleteUserGroup() error {
	api_key := "CLICKUP_API_KEY"
	group_id := "GROUP_ID"
	client := clickup.NewClient(nil, api_key)

	_, err := client.UserGroups.DeleteUserGroup(context.Background(), group_id)

	return err
}

func main() {
	groups, err := fetchUserGroups()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, group := range groups {
		fmt.Println(group.Name, group.ID, group.Members)
	}
}
