package clickup

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTeamsService_GetTeams(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/team", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
			  "teams": [
				{
				  "id": "1234",
				  "name": "My ClickUp Team",
				  "color": "#000000",
				  "avatar": "https://clickup.com/avatar.jpg",
				  "members": [
					{
					  "user": {
						"id": 123,
						"username": "John Doe",
						"color": "#000000",
						"profilePicture": "https://clickup.com/avatar.jpg"
					  }
					}
				  ]
				}
			  ]
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Authorization.GetAuthorizedTeams(ctx)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	var want []Team
	user := TeamUser{
		ID:             123,
		Username:       "John Doe",
		Color:          "#000000",
		ProfilePicture: "https://clickup.com/avatar.jpg",
	}
	m := TeamMember{
		User: user,
	}
	team := Team{
		ID:      "1234",
		Name:    "My ClickUp Team",
		Color:   "#000000",
		Avatar:  "https://clickup.com/avatar.jpg",
		Members: []TeamMember{m},
	}
	want = append(want, team)
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
