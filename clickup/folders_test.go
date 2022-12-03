package clickup

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoldersService_GetFolders(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/789/folder", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"folders": [
					{
						"id": "457",
						"name": "Updated Folder Name",
						"orderindex": 0,
						"override_statuses": false,
						"hidden": false,
						"space": {
						  "id": "789",
						  "name": "Space Name",
						  "access": true
						},
						"task_count": "0",
						"lists": []
					},
					{
						"id": "458",
						"name": "Second Folder Name",
						"orderindex": 1,
						"override_statuses": false,
						"hidden": false,
						"space": {
						  "id": "789",
						  "name": "Space Name",
						  "access": true
						},
						"task_count": "0",
						"lists": []
					}
				]
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Folders.GetFolders(ctx, "789", true)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	var want []Folder
	folder1 := Folder{
		ID:               "457",
		Name:             "Updated Folder Name",
		Orderindex:       "0",
		OverrideStatuses: false,
		Hidden:           false,
		Space:            SpaceOfFolderBelonging{ID: "789", Name: "Space Name", Access: true},
		TaskCount:        "0",
		Lists:            []ListOfFolderBelonging{},
	}
	folder2 := Folder{
		ID:               "458",
		Name:             "Second Folder Name",
		Orderindex:       "1",
		OverrideStatuses: false,
		Hidden:           false,
		Space:            SpaceOfFolderBelonging{ID: "789", Name: "Space Name", Access: true},
		TaskCount:        "0",
		Lists:            []ListOfFolderBelonging{},
	}
	want = append(want, folder1, folder2)
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestFoldersService_GetFolder(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/folder/457", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"id": "457",
				"name": "Updated Folder Name",
				"orderindex": 0,
				"override_statuses": false,
				"hidden": false,
				"space": {
				  "id": "789",
				  "name": "Space Name",
				  "access": true
				},
				"task_count": "0",
				"lists": []
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Folders.GetFolder(ctx, "457")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := &Folder{
		ID:               "457",
		Name:             "Updated Folder Name",
		Orderindex:       "0",
		OverrideStatuses: false,
		Hidden:           false,
		Space:            SpaceOfFolderBelonging{ID: "789", Name: "Space Name", Access: true},
		TaskCount:        "0",
		Lists:            []ListOfFolderBelonging{},
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
