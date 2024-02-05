package clickup

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
						"id": "a-b-c-d-e",
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
								"profilePicture": "test.jpg"
							}
						],
						"avatar": {
							"attachment_id": null,
							"color": null,
							"source": null,
							"icon": null
						}						
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
		ProfilePicture: "test.jpg",
	}
	ug := UserGroup{
		ID:          "a-b-c-d-e",
		TeamID:      "123",
		UserID:      321,
		Name:        "Test Team",
		Handle:      "test-team",
		DateCreated: "1675407473787",
		Initials:    "TT",
		Members:     []GroupMember{m},
		Avatar: UserGroupAvatar{
			AttachmentId: nil,
			Color:        nil,
			Source:       nil,
			Icon:         nil,
		},
	}
	want := []UserGroup{ug}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestUserGroupsService_CreateUserGroup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/team/123/group", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w,
			`{
				"id": "a-b-c-d-e",
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
						"profilePicture": "test.jpg"
					}
				],
				"avatar": {
					"attachment_id": null,
					"color": null,
					"source": null,
					"icon": null
				}
			}`,
		)
	})
	ctx := context.Background()
	artifacts, _, err := client.UserGroups.CreateUserGroup(ctx, "123", &CreateUserGroupRequest{
		Name:    "Test Team",
		Members: []int{321},
	})
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	m := GroupMember{
		ID:             321,
		Username:       "Test User",
		Email:          "testuser@test.com",
		Color:          "#7b68ee",
		Initials:       "TU",
		ProfilePicture: "test.jpg",
	}
	want := &UserGroup{
		ID:          "a-b-c-d-e",
		TeamID:      "123",
		UserID:      321,
		Name:        "Test Team",
		Handle:      "test-team",
		DateCreated: "1675407473787",
		Initials:    "TT",
		Members:     []GroupMember{m},
		Avatar: UserGroupAvatar{
			AttachmentId: nil,
			Color:        nil,
			Source:       nil,
			Icon:         nil,
		},
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestUserGroupsService_UpdateUserGroup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/group/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var requestBody UpdateUserGroupRequest
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		userIds := requestBody.Members.Add

		members := make([]map[string]interface{}, len(userIds))
		for i, id := range userIds {
			members[i] = map[string]interface{}{
				"id":             id,
				"username":       fmt.Sprintf("Test User %d", id),
				"email":          fmt.Sprintf("testuser%d@test.com", id),
				"color":          "#7b68ee",
				"initials":       "TU",
				"profilePicture": "test.jpg",
			}
		}

		// Generate the response
		response := map[string]interface{}{
			"id":           "a-b-c-d-e",
			"team_id":      "123",
			"userid":       321,
			"name":         requestBody.Name,
			"handle":       requestBody.Handle,
			"date_created": "1675407473787",
			"initials":     "TT",
			"members":      members,
			"avatar": map[string]interface{}{
				"attachment_id": nil,
				"color":         nil,
				"source":        nil,
				"icon":          nil,
			},
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error generating response", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(responseJSON))
	})
	ctx := context.Background()
	artifacts, _, err := client.UserGroups.UpdateUserGroup(ctx, "123", &UpdateUserGroupRequest{
		Name:   "Test Team Updated",
		Handle: "New Handle",
		Members: UpdateUserGroupMember{
			Add: []int{1, 2},
		},
	})
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	m1 := GroupMember{
		ID:             1,
		Username:       "Test User 1",
		Email:          "testuser1@test.com",
		Color:          "#7b68ee",
		Initials:       "TU",
		ProfilePicture: "test.jpg",
	}
	m2 := GroupMember{
		ID:             2,
		Username:       "Test User 2",
		Email:          "testuser2@test.com",
		Color:          "#7b68ee",
		Initials:       "TU",
		ProfilePicture: "test.jpg",
	}
	want := &UserGroup{
		ID:          "a-b-c-d-e",
		TeamID:      "123",
		UserID:      321,
		Name:        "Test Team Updated",
		Handle:      "New Handle",
		DateCreated: "1675407473787",
		Initials:    "TT",
		Members:     []GroupMember{m1, m2},
		Avatar: UserGroupAvatar{
			AttachmentId: nil,
			Color:        nil,
			Source:       nil,
			Icon:         nil,
		},
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestUserGroupsService_DeleteUserGroup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/group/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, nil)
	})
	ctx := context.Background()
	_, err := client.UserGroups.DeleteUserGroup(ctx, "123")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}
}

func TestUserGroupsService_DeleteUserGroupError(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/group/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal server error")
	})
	ctx := context.Background()
	response, err := client.UserGroups.DeleteUserGroup(ctx, "123")

	if err == nil {
		t.Errorf("Actions.ListArtifacts did not return error: %v", response)
	}

	expectedError := "api/v2/group/123: 500"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("UserGroups.DeleteUserGroup returned error: %v, want %v", err, expectedError)
	}
}
