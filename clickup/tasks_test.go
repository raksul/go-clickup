package clickup

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTasksService_GetTask(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/task/9hv/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"id": "9hx",
    		    "custom_id":null,
    		    "name": "Task Name",
    		    "text_content": "New Task Description",
    		    "description": "New Task Description",
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
    		    "custom_fields": [
    		        {
    		        	"id": "0a52c486-5f05-403b-b4fd-c512ff05131c",
    		        	"name": "My Number field",
    		        	"type": "checkbox",
    		        	"type_config": {},
    		        	"date_created": "1622176979540",
    		        	"hide_from_guests": false,
    		        	"value": "23",
    		        	"required": true
    		        }
    		    ],
    		    "list": {
    		        "id": "123"
    		    },
    		    "folder": {
    		        "id": "456"
    		    },
    		    "space": {
    		        "id": "789"
    		    },
    		    "url": "https://app.clickup.com/t/9hx"
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Tasks.GetTask(ctx, "9hv", nil)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &Task{
		ID:          "9hx",
		Name:        "Task Name",
		TextContent: "New Task Description",
		Description: "New Task Description",
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
		Assignees:  []User{},
		Checklists: []Checklist{},
		Tags:       []Tag{},
		CustomFields: []CustomField{
			{
				ID:             "0a52c486-5f05-403b-b4fd-c512ff05131c",
				Name:           "My Number field",
				Type:           "checkbox",
				DateCreated:    "1622176979540",
				HideFromGuests: false,
				TypeConfig:     map[string]interface{}{},
				Value:          "23",
			},
		},
		List:   ListOfTaskBelonging{ID: "123"},
		Folder: FolderOftaskBelonging{ID: "456"},
		Space:  SpaceOfTaskBelonging{ID: "789"},
		URL:    "https://app.clickup.com/t/9hx",
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
