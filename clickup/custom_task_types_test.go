package clickup

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCustomTaskTypesService_GetCustomTaskTypes(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(("/team/1234/custom_item"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"custom_items": [{"id": 1300, "name": "Bug", "name_plural": "Bugs", "description": "Custom item type for bugs."}]}`)
	})

	ctx := context.Background()
	customTaskTypes, _, err := client.CustomTaskTypes.GetCustomTaskTypes(ctx, "1234")
	if err != nil {
		t.Errorf("CustomTaskTypes.GetCustomTaskTypes returned error: %v", err)
	}

	want := []CustomItem{
		{Id: 1300, Name: "Bug", NamePlural: "Bugs", Description: "Custom item type for bugs."},
	}
	if !cmp.Equal(customTaskTypes, want) {
		t.Errorf("CustomTaskTypes.GetCustomTaskTypes returned %+v, want %+v", customTaskTypes, want)
	}
}
