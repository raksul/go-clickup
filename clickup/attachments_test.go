package clickup

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAttachmentsService_CreateTaskAttachment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Attachment{
		FileName: "test.txt",
		Reader:   bytes.NewBufferString("test test test test"),
	}

	mux.HandleFunc("/task/9hz/attachment", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w,
			`{
				"id": "ac434d4e-8b1c-4571-951b-866b6d9f2ee6.png",
				"version": "0",
				"date": 1569988578766,
				"title": "image.png",
				"extension": "png",
				"thumbnail_small": "https://attachments-public.clickup.com/ac434d4e-8b1c-4571-951b-866b6d9f2ee6/logo_small.png",
				"thumbnail_large": "https://attachments-public.clickup.com/ac434d4e-8b1c-4571-951b-866b6d9f2ee6/logo_small.png",
				"url": "https://attachments-public.clickup.com/ac434d4e-8b1c-4571-951b-866b6d9f2ee6/logo_small.png"
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Attachments.CreateTaskAttachment(ctx, "9hz", nil, input)
	if err != nil {
		t.Errorf("Actions.CreateTaskAttachment returned error: %v", err)
	}

	want := &CreateAttachmentResponse{
		ID:             "ac434d4e-8b1c-4571-951b-866b6d9f2ee6.png",
		Version:        "0",
		Date:           1569988578766,
		Title:          "image.png",
		Extension:      "png",
		ThumbnailSmall: "https://attachments-public.clickup.com/ac434d4e-8b1c-4571-951b-866b6d9f2ee6/logo_small.png",
		ThumbnailLarge: "https://attachments-public.clickup.com/ac434d4e-8b1c-4571-951b-866b6d9f2ee6/logo_small.png",
		URL:            "https://attachments-public.clickup.com/ac434d4e-8b1c-4571-951b-866b6d9f2ee6/logo_small.png",
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.CreateTaskAttachment returned %+v, want %+v", artifacts, want)
	}
}
