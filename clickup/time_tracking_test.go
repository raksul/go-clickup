package clickup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTimeTrackingService_CreateTimeTracking(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &TimeTrackingRequest{
		Description: "description",
		Start:       1719595398,
		Duration:    120000,
		Assignee:    99999999,
		Tid:         "9hz",
		Billable:    true,
	}

	mux.HandleFunc("/team/123/time_entries", func(w http.ResponseWriter, r *http.Request) {
		v := new(TimeTrackingRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w,
			`{
				"data": {
					"id": "4090130922962924695",
					"task": {
						"id": "9hz",
						"name": "Test",
						"status": {
							"status": "in progress",
							"id": "p99999999999_99asdhAS",
							"color": "#5f55ee",
							"type": "custom",
							"orderindex": 1
						}
					},
					"wid": "9999999999",
					"user": {
						"id": 99999999,
						"username": "John",
						"email": "john@mail.com",
						"color": "#afb42b",
						"initials": "J",
						"profilePicture": "https://attachments.clickup.com/profilePictures/99999999_tX9.jpg"
					},
					"billable": true,
					"start": 1719595398,
					"end": "1719715398",
					"duration": 120000,
					"description": "description",
					"tags": [],
					"at": 1719586940375,
					"is_locked": false,
					"task_location": {}
				}
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.TimeTrackings.CreateTimeTracking(ctx, "123", nil, input)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &CreateTimeTrackingResponse{Data: TimeTrackingData{
		ID:           "4090130922962924695",
		Wid:          "9999999999",
		User:         User{ID: 99999999, Username: "John", Email: "john@mail.com", Color: "#afb42b", Initials: "J", ProfilePicture: "https://attachments.clickup.com/profilePictures/99999999_tX9.jpg"},
		Billable:     true,
		Start:        1719595398,
		End:          "1719715398",
		Duration:     120000,
		Description:  "description",
		Tags:         []TimeTrackingTag{},
		At:           1719586940375,
		IsLocked:     false,
		TaskLocation: TaskLocation{},
		Task:         Task{ID: "9hz", Name: "Test", Status: TaskStatus{ID: "p99999999999_99asdhAS", Status: "in progress", Color: "#5f55ee", Type: "custom", Orderindex: json.Number("1")}},
	}}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}

}

func TestTimeTrackingService_GetSingularTimeEntry(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/team/123/time_entries/456", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"data": {
					"id": "456",
					"task": {
						"id": "99a999qj9",
						"name": "test",
						"status": {
							"status": "Closed",
							"color": "#008844",
							"type": "closed",
							"orderindex": 2
						}
					},
					"wid": "9997138699",
					"user": {
						"id": 99999999,
						"username": "John",
						"email": "john@mail.com",
						"color": "#afb42b",
						"initials": "J",
						"profilePicture": "https://attachments.clickup.com/profilePictures/99999999_tX9.jpg"
					},
					"billable": true,
					"start": "1720012882348",
					"end": "1720013482348",
					"duration": "600000",
					"description": "",
					"tags": [],
					"source": "clickup",
					"at": "1720013485531",
					"is_locked": false,
					"task_location": {
						"list_id": "999904299397",
						"folder_id": "99992399554",
						"space_id": "99990359986"
					},
					"task_url": "https://app.clickup.com/t/99a999qj9"
				}
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.TimeTrackings.GetSingularTimeEntry(ctx, "123", "456", nil)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &GetTimeTrackingResponse{Data: GetTimeTrackingData{
		ID:           "456",
		Wid:          "9997138699",
		User:         User{ID: 99999999, Username: "John", Email: "john@mail.com", Color: "#afb42b", Initials: "J", ProfilePicture: "https://attachments.clickup.com/profilePictures/99999999_tX9.jpg"},
		Billable:     true,
		Start:        "1720012882348",
		End:          "1720013482348",
		Duration:     "600000",
		Description:  "",
		Tags:         []TimeTrackingTag{},
		Source:       "clickup",
		At:           "1720013485531",
		IsLocked:     false,
		TaskLocation: GetTaskLocation{ListID: "999904299397", FolderID: "99992399554", SpaceID: "99990359986"},
		Task:         Task{ID: "99a999qj9", Name: "test", Status: TaskStatus{Status: "Closed", Color: "#008844", Type: "closed", Orderindex: json.Number("2")}},
		TaskURL:      "https://app.clickup.com/t/99a999qj9",
	}}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
