package clickup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDependenciesService_AddDependency(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &AddDependencyRequest{
		DependsOn:    "abc",
		DependencyOf: "def",
	}

	mux.HandleFunc("/task/9hz/dependency", func(w http.ResponseWriter, r *http.Request) {
		v := new(AddDependencyRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.WriteHeader(http.StatusOK)
	})

	ctx := context.Background()
	_, err := client.Dependencies.AddDependency(ctx, "9hz", input, nil)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}
}

func TestDependenciesService_DeleteDependency(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &DeleteDependencyOptions{
		DependsOn: "abc",
	}

	mux.HandleFunc("/task/9hz/dependency", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusOK)
	})

	ctx := context.Background()
	_, err := client.Dependencies.DeleteDependency(ctx, "9hz", input)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}
}

func TestDependenciesService_AddTaskLink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/task/9hv/link/9hz", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w,
			`{
				"task": {
					"id": "9hv",
					"name": "Task Name",
					"status": {
					  "status": "in progress",
					  "color": "#d3d3d3",
					  "orderindex": 1,
					  "type": "custom"
					},
					"orderindex": "1.00000000000000000000000000000000",
					"date_created": "1567780450202",
					"date_updated": "1567780450202",
					"date_closed": null,
					"creator": {
					  "id": 183,
					  "username": "John Doe",
					  "color": "#827718",
					  "profilePicture": "https://attachments-public.clickup.com/profilePictures/183_abc.jpg"
					},
					"assignees": [],
					"checklists": [],
					"tags": [],
					"parent": null,
					"priority": null,
					"due_date": null,
					"start_date": null,
					"time_estimate": null,
					"time_spent": null,
					"list": {
					  "id": "123"
					},
					"folder": {
					  "id": "456"
					},
					"space": {
					  "id": "789"
					},
					"linked_tasks": [
					  {
						"task_id": "9hv",
						"link_id": "9hz",
						"date_created": "1587571108988",
						"userid": "183"
					  }
					],
					"url": "https://app.clickup.com/t/9hx"
				}
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Dependencies.AddTaskLink(ctx, "9hv", "9hz", nil)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &Task{
		ID:   "9hv",
		Name: "Task Name",
		Status: TaskStatus{
			Status:     "in progress",
			Color:      "#d3d3d3",
			Orderindex: "1",
			Type:       "custom",
		},
		Orderindex:  "1.00000000000000000000000000000000",
		DateCreated: "1567780450202",
		DateUpdated: "1567780450202",
		Creator: User{
			ID:             183,
			Username:       "John Doe",
			Color:          "#827718",
			ProfilePicture: "https://attachments-public.clickup.com/profilePictures/183_abc.jpg",
		},
		Assignees:   []User{},
		Checklists:  []Checklist{},
		Tags:        []Tag{},
		List:        ListOfTaskBelonging{ID: "123"},
		Folder:      FolderOftaskBelonging{ID: "456"},
		Space:       SpaceOfTaskBelonging{ID: "789"},
		LinkedTasks: []LinkedTask{{TaskID: "9hv", LinkID: "9hz", DateCreated: "1587571108988", Userid: "183"}},
		URL:         "https://app.clickup.com/t/9hx",
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestDependenciesService_DeleteTaskLink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/task/9hv/link/9hz", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w,
			`{
				"task": {
					"id": "9hv",
					"name": "Task Name",
					"status": {
					  "status": "in progress",
					  "color": "#d3d3d3",
					  "orderindex": 1,
					  "type": "custom"
					},
					"orderindex": "1.00000000000000000000000000000000",
					"date_created": "1567780450202",
					"date_updated": "1567780450202",
					"date_closed": null,
					"creator": {
					  "id": 183,
					  "username": "John Doe",
					  "color": "#827718",
					  "profilePicture": "https://attachments-public.clickup.com/profilePictures/183_abc.jpg"
					},
					"assignees": [],
					"checklists": [],
					"tags": [],
					"parent": null,
					"priority": null,
					"due_date": null,
					"start_date": null,
					"time_estimate": null,
					"time_spent": null,
					"list": {
					  "id": "123"
					},
					"folder": {
					  "id": "456"
					},
					"space": {
					  "id": "789"
					},
					"linked_tasks": [],
					"url": "https://app.clickup.com/t/9hx"
				}
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Dependencies.DeleteTaskLink(ctx, "9hv", "9hz", nil)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &Task{
		ID:   "9hv",
		Name: "Task Name",
		Status: TaskStatus{
			Status:     "in progress",
			Color:      "#d3d3d3",
			Orderindex: "1",
			Type:       "custom",
		},
		Orderindex:  "1.00000000000000000000000000000000",
		DateCreated: "1567780450202",
		DateUpdated: "1567780450202",
		Creator: User{
			ID:             183,
			Username:       "John Doe",
			Color:          "#827718",
			ProfilePicture: "https://attachments-public.clickup.com/profilePictures/183_abc.jpg",
		},
		Assignees:   []User{},
		Checklists:  []Checklist{},
		Tags:        []Tag{},
		List:        ListOfTaskBelonging{ID: "123"},
		Folder:      FolderOftaskBelonging{ID: "456"},
		Space:       SpaceOfTaskBelonging{ID: "789"},
		LinkedTasks: []LinkedTask{},
		URL:         "https://app.clickup.com/t/9hx",
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
