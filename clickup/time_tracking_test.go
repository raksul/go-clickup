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
