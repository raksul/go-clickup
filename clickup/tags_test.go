package clickup

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTagsService_GetTags(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/512/tag", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"tags": [
				    {
				        "name": "Tag name",
				        "tag_fg": "#000000",
				        "tag_bg": "#000000"
				    }
				]
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Tags.GetTags(ctx, "512")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := []Tag{
		{
			Name:  "Tag name",
			TagFg: "#000000",
			TagBg: "#000000",
		},
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}
