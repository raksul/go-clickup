package clickup

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCustomFieldsService_GetAccessibleCustomFields(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/list/123/field", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"fields": [
					{
					  "id": "03efda77-c7a0-42d3-8afd-fd546353c2f5",
					  "name": "Text Field",
					  "type": "text",
					  "type_config": {},
					  "date_created": "1566400407303",
					  "hide_from_guests": false
					},
					{
					  "id": "0a52c486-5f05-403b-b4fd-c512ff05131c",
					  "name": "Number Field",
					  "type": "number",
					  "type_config": {},
					  "date_created": "1565993269460",
					  "hide_from_guests": false
					}
				]
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.CustomFields.GetAccessibleCustomFields(ctx, "123")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	var want []CustomField
	cf1 := CustomField{
		ID:             "03efda77-c7a0-42d3-8afd-fd546353c2f5",
		Name:           "Text Field",
		Type:           "text",
		DateCreated:    "1566400407303",
		HideFromGuests: false,
		TypeConfig:     map[string]interface{}{},
	}
	cf2 := CustomField{
		ID:             "0a52c486-5f05-403b-b4fd-c512ff05131c",
		Name:           "Number Field",
		Type:           "number",
		DateCreated:    "1565993269460",
		HideFromGuests: false,
		TypeConfig:     map[string]interface{}{},
	}
	want = append(want, cf1, cf2)
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestCustomFieldsService_RemoveCustomFieldValue(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/task/9hz/field/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusOK)
	})

	ctx := context.Background()
	_, err := client.CustomFields.RemoveCustomFieldValue(ctx, "9hz", "123", nil)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}
}
