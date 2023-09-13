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
	artifacts, _, err := client.Teams.GetTeams(ctx)
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

func TestTeamsService_GetSeats(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/team/123/seats", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(
			w,
			`{
                "members": {
                  "filled_members_seats": 9,
                  "total_member_seats": 9,
                  "empty_member_seats": 0
                },
                "guests": {
                  "filled_guest_seats": 2,
                  "total_guest_seats": 50,
                  "empty_guest_seats": 48
                }
            }`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Teams.GetSeats(ctx, "123")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	var want Seats
	members := Members{
		FilledMembersSeats: 9,
		TotalMemberSeats:   9,
		EmptyMemberSeats:   0,
	}
	guests := Guests{
		FilledGuestSeats: 2,
		TotalGuestSeats:  50,
		EmptyGuestSeats:  48,
	}
	want.Members = members
	want.Guests = guests
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestTeamsService_GetPlan(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/team/123/plan", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(
			w,
			`{
                "plan_name": "Free Forever",
                "plan_id": 1
            }`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Teams.GetPlan(ctx, "123")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := Plan{
		Id:   1,
		Name: "Free Forever",
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
