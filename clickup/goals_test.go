package clickup

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGoalsService_GetGoals(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/team/789/goal", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"goals": [
					{
					    "id": "e53a033c-900e-462d-a849-4a216b06d930",
					    "name": "Updated Goal Name",
					    "team_id": "512",
					    "date_created": "1568044355026",
					    "due_date": "1568036964079",
					    "description": "Updated Goal Description",
					    "private": false,
					    "archived": false,
					    "creator": 183,
					    "color": "#32a852",
					    "pretty_id": "6",
					    "multiple_owners": true,
						"folder_id": null,
					    "members": [],
					    "owners": [
						    {
						      "id": 182,
						      "username": "Jane Doe",
						      "initials": "JD",
						      "email": "janedoe@gmail.com",
						      "color": "#827718",
						      "profilePicture": "https://attachments-public.clickup.com/profilePictures/182_abc.jpg"
						    }
					    ],
					    "key_results": [],
					    "percent_completed": 0,
					    "history": [],
					    "pretty_url": "https://app.clickup.com/512/goals/6"
					}
				],
				"folders": [
					{
					    "id": "05921253-7737-44af-a1aa-36fd11244e6f",
					    "name": "Goal Folder",
					    "team_id": "512",
					    "private": true,
					    "date_created": "1548802674671",
					    "creator": 182,
					    "goal_count": 0,
					    "group_members": [
							{
							    "id": 182,
							    "username": "Jane Doe",
							    "initials": "JD",
							    "email": "janedoe@gmail.com",
							    "color": "#827718",
							    "profilePicture": "https://attachments-public.clickup.com/profilePictures/182_abc.jpg"
							}
					    ],
					    "goals": []
					}
				]
			}`,
		)
	})

	ctx := context.Background()
	goals, folders, _, err := client.Goals.GetGoals(ctx, "789", true)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	wantGoals := []Goal{
		{
			ID:             "e53a033c-900e-462d-a849-4a216b06d930",
			Name:           "Updated Goal Name",
			TeamID:         "512",
			DateCreated:    "1568044355026",
			DueDate:        "1568036964079",
			Description:    "Updated Goal Description",
			Private:        false,
			Archived:       false,
			Creator:        183,
			Color:          "#32a852",
			PrettyID:       "6",
			MultipleOwners: true,
			Owners: []GoalOwner{
				{
					ID:             182,
					Username:       "Jane Doe",
					Initials:       "JD",
					Email:          "janedoe@gmail.com",
					Color:          "#827718",
					ProfilePicture: "https://attachments-public.clickup.com/profilePictures/182_abc.jpg",
				},
			},
			Members:          []GoalMember{},
			PercentCompleted: 0,
		},
	}
	wantFolders := []GoalFolder{
		{
			ID:          "05921253-7737-44af-a1aa-36fd11244e6f",
			Name:        "Goal Folder",
			TeamID:      "512",
			Private:     true,
			DateCreated: "1548802674671",
			Creator:     182,
			GoalCount:   0,
			GroupMembers: []GoalMember{
				{
					ID:             182,
					Username:       "Jane Doe",
					Initials:       "JD",
					Email:          "janedoe@gmail.com",
					Color:          "#827718",
					ProfilePicture: "https://attachments-public.clickup.com/profilePictures/182_abc.jpg",
				},
			},
			Goals: []Goal{},
		},
	}

	if !cmp.Equal(goals, wantGoals) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", goals, wantGoals)
	}
	if !cmp.Equal(folders, wantFolders) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", folders, wantFolders)
	}
}

func TestGoalsService_GetGoal(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/goal/e53a033c-900e-462d-a849-4a216b06d930", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"goal": {
					"id": "e53a033c-900e-462d-a849-4a216b06d930",
					"name": "Updated Goal Name",
					"team_id": "512",
					"date_created": "1568044355026",
					"due_date": "1568036964079",
					"description": "Updated Goal Description",
					"private": false,
					"archived": false,
					"creator": 183,
					"color": "#32a852",
					"pretty_id": "6",
					"multiple_owners": true,
					"members": [],
					"owners": [
						 {
						   "id": 182,
						   "username": "Jane Doe",
						   "initials": "JD",
						   "email": "janedoe@gmail.com",
						   "color": "#827718",
						   "profilePicture": "https://attachments-public.clickup.com/profilePictures/182_abc.jpg"
						 }
					],
					"key_results": [],
					"percent_completed": 0,
					"history": [],
					"pretty_url": "https://app.clickup.com/512/goals/6"
				}
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Goals.GetGoal(ctx, "e53a033c-900e-462d-a849-4a216b06d930")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &Goal{
		ID:             "e53a033c-900e-462d-a849-4a216b06d930",
		Name:           "Updated Goal Name",
		TeamID:         "512",
		DateCreated:    "1568044355026",
		DueDate:        "1568036964079",
		Description:    "Updated Goal Description",
		Private:        false,
		Archived:       false,
		Creator:        183,
		Color:          "#32a852",
		PrettyID:       "6",
		MultipleOwners: true,
		Owners: []GoalOwner{
			{
				ID:             182,
				Username:       "Jane Doe",
				Initials:       "JD",
				Email:          "janedoe@gmail.com",
				Color:          "#827718",
				ProfilePicture: "https://attachments-public.clickup.com/profilePictures/182_abc.jpg",
			},
		},
		Members:          []GoalMember{},
		PercentCompleted: 0,
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
