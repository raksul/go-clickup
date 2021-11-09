package clickup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestChecklistsService_CreateChecklist(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &ChecklistRequest{
		Name:     "test",
		Position: 2,
	}

	mux.HandleFunc("/task/9hz/checklist/", func(w http.ResponseWriter, r *http.Request) {
		v := new(ChecklistRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w,
			`{
				"checklist": {
					"id": "b955c4dc-b8a8-48d8-a0c6-b4200788a683",
					"task_id": "9hz",
					"name": "Checklist",
					"orderindex": 0,
					"resolved": 0,
					"unresolved": 0,
					"items": []
				}
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Checklists.CreateChecklist(ctx, "9hz", nil, input)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &Checklist{ID: "b955c4dc-b8a8-48d8-a0c6-b4200788a683", Name: "Checklist", TaskID: "9hz", Orderindex: 0, Resolved: 0, Unresolved: 0, Items: []Item{}}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestChecklistsService_EditChecklist(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &ChecklistRequest{
		Name:     "test",
		Position: 2,
	}

	mux.HandleFunc("/checklist/b955c4dc-b8a8-48d8-a0c6-b4200788a683", func(w http.ResponseWriter, r *http.Request) {
		v := new(ChecklistRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w,
			`{
				"checklist": {
					"id": "b955c4dc-b8a8-48d8-a0c6-b4200788a683",
           			"task_id": "9hz",
           			"name": "Updated Checklist",
           			"orderindex": 5,
           			"resolved": 0,
           			"unresolved": 0,
           			"items": []
				}
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Checklists.EditChecklist(ctx, "b955c4dc-b8a8-48d8-a0c6-b4200788a683", input)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &Checklist{ID: "b955c4dc-b8a8-48d8-a0c6-b4200788a683", Name: "Updated Checklist", TaskID: "9hz", Orderindex: 5, Resolved: 0, Unresolved: 0, Items: []Item{}}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestChecklistsService_DeleteChecklist(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/checklist/b955c4dc-b8a8-48d8-a0c6-b4200788a683", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusOK)
	})

	ctx := context.Background()
	_, err := client.Checklists.DeleteChecklist(ctx, "b955c4dc-b8a8-48d8-a0c6-b4200788a683")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}
}
