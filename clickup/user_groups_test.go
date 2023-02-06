package clickup

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUserGroupsService_GetUserGroups(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/group", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"groups": [
					{
						"id": "d899ba3a-1cbf-4188-ae39-11e6acc5f93d",
						"team_id": "123",
						"userid": 321,
						"name": "Test Team",
						"handle": "test-team",
						"date_created": "1675407473787",
						"initials": "TT",
						"members": [
							{
								"id": 321,
								"username": "Test User",
								"email": "testuser@test.com",
								"color": "#7b68ee",
								"initials": "TU",
								"profilePicture": "https://attachments-public.clickup.com/profilePictures/812_nx1.jpg"
							}
						]
					}
				]
			}`,
		)
	})

	ctx := context.Background()
	opts := &GetUserGroupsOptions{
		TeamID: "123",
	}
	artifacts, _, err := client.UserGroups.GetUserGroups(ctx, opts)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	m := GroupMember{
		ID:             321,
		Username:       "Test User",
		Email:          "testuser@test.com",
		Color:          "#7b68ee",
		Initials:       "TU",
		ProfilePicture: "https://attachments-public.clickup.com/profilePictures/812_nx1.jpg",
	}
	ug := UserGroup{
		ID:          "d899ba3a-1cbf-4188-ae39-11e6acc5f93d",
		TeamID:      "123",
		UserID:      321,
		Name:        "Test Team",
		Handle:      "test-team",
		DateCreated: "1675407473787",
		Initials:    "TT",
		Members:     []GroupMember{m},
	}
	want := []UserGroup{ug}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
