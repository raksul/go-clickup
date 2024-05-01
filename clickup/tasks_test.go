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
    		    "custom_item_id": null,
    		    "name": "Task Name",
    		    "text_content": "New Task Description",
    		    "description": "New Task Description",
				"markdown_description": "## New Task Description",
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
		ID:           "9hx",
		CustomItemId: 0,
		Name:         "Task Name",
		TextContent:  "New Task Description",
		Description:  "New Task Description",
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

func TestUrlEncodeCustomFieldsInGetTasksRequest(t *testing.T) {

	var gtOpts = GetTasksOptions{
		Archived: false,
		Page:     3,
		Reverse:  false,
		CustomFields: CustomFieldsInGetTasksRequest{
			{
				FieldId:  "de761538-8ae0-42e8-91d9-f1a0cdfbd8b5",
				Operator: GreaterThan,
				Value:    []string{"2"},
			},
		},
	}
	options, err := addOptions("https://www.example.org/", gtOpts)
	if err != nil {
		t.Errorf("expected no error but got error: %v", err)
	}

	want := "https://www.example.org/?custom_fields=%5B%7B%22field_id%22%3A%22de761538-8ae0-42e8-91d9-f1a0cdfbd8b5%22%2C%22operator%22%3A%22%3E%22%2C%22value%22%3A%222%22%7D%5D&page=3"

	if !cmp.Equal(options, want) {

		t.Errorf("addOptions returned %+v, want %+v", options, want)
	}

	gtOpts = GetTasksOptions{
		Archived: false,
		Page:     3,
		Reverse:  false,
		CustomFields: CustomFieldsInGetTasksRequest{
			{
				FieldId:  "de761538-8ae0-42e8-91d9-f1a0cdfbd8b5",
				Operator: GreaterThan,
				Value:    []string{"2"},
			},
			{
				FieldId:  "4223cfb4-b14b-4bd4-aa35-81ae29c62f4d",
				Operator: IsNotNull,
			},
			{
				FieldId:  "4d4044e9-4819-4819-af2a-34fde3a41903",
				Operator: Range,
				Value:    []string{"5", "10"},
			},
			{
				FieldId:  "1bea591e-22a6-4485-8614-88eb3d21c188",
				Operator: Any,
				Value:    []string{"5", "10", "15", "20"},
			},
		},
	}
	options, err = addOptions("https://www.example.org/", gtOpts)
	if err != nil {
		t.Errorf("expected no error but got error: %v", err)
	}

	want = "https://www.example.org/?custom_fields=%5B%7B%22field_id%22%3A%22de761538-8ae0-42e8-91d9-f1a0cdfbd8b5%22%2C%22operator%22%3A%22%3E%22%2C%22value%22%3A%222%22%7D%2C%7B%22field_id%22%3A%224223cfb4-b14b-4bd4-aa35-81ae29c62f4d%22%2C%22operator%22%3A%22IS+NOT+NULL%22%7D%2C%7B%22field_id%22%3A%224d4044e9-4819-4819-af2a-34fde3a41903%22%2C%22operator%22%3A%22RANGE%22%2C%22value%22%3A%5B%225%22%2C%2210%22%5D%7D%2C%7B%22field_id%22%3A%221bea591e-22a6-4485-8614-88eb3d21c188%22%2C%22operator%22%3A%22ANY%22%2C%22value%22%3A%5B%225%22%2C%2210%22%2C%2215%22%2C%2220%22%5D%7D%5D&page=3"

	if !cmp.Equal(options, want) { //"https://www.example.org/?custom_fields=[{\"field_id\":\"de761538-8ae0-42e8-91d9-f1a0cdfbd8b5\",\"operator\":\">\",\"value\":"2"}]&page=3",

		t.Errorf("addOptions returned %+v, want %+v", options, want)
	}
}
