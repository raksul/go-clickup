package clickup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWebhooksService_GetWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/team/123/webhook", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w,
			`{
				"webhooks": [
				  {
					"id": "4b67ac88-e506-4a29-9d42-26e504e3435e",
					"userid": 183,
					"team_id": 108,
					"endpoint": "https://yourdomain.com/webhook",
					"client_id": "QVOQP06ZXC6CMGVFKB0ZT7J9Y7APOYGO",
					"events": [
						"taskCreated",
						"keyResultDeleted"
					],
					"task_id": null,
					"list_id": null,
					"folder_id": null,
					"space_id": null,
					"health": {
						"status": "failing",
						"fail_count": 5
					},
					"secret": "O94IM25S7PXBPYTMNXLLET230SRP0S89COR7B1YOJ2ZIE8WQNK5UUKEF26W0Z5GA"
				  },
				  {
					"id": "4b67ac88-e506-4a29-9d42-26e504e3435f",
					"userid": 200,
					"team_id": 108,
					"endpoint": "https://yourdomain.com/webhook",
					"client_id": "QVOQP06ZXC6CMGVFKB0ZT7J9Y7APOYGO",
					"events": [
						"taskTagUpdated",
						"taskMoved"
					],
					"task_id": null,
					"list_id": null,
					"folder_id": null,
					"space_id": null,
					"health": {
						"status": "active",
						"fail_count": 0
					},
					"secret": "O94IM25S7PXBPYTMNXLLET230SRP0S89COR7B1YOJ2ZIE8WQNK5UUKEF26W0Z5GA"
				  }
				]
			  }`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Webhooks.GetWebhook(ctx, 123)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := []Webhook{
		{
			ID:       "4b67ac88-e506-4a29-9d42-26e504e3435e",
			Userid:   183,
			TeamID:   108,
			Endpoint: "https://yourdomain.com/webhook",
			ClientID: "QVOQP06ZXC6CMGVFKB0ZT7J9Y7APOYGO",
			Events:   []string{"taskCreated", "keyResultDeleted"},
			Health: webhookHealth{
				Status:    "failing",
				FailCount: 5,
			},
			Secret: "O94IM25S7PXBPYTMNXLLET230SRP0S89COR7B1YOJ2ZIE8WQNK5UUKEF26W0Z5GA",
		},
		{
			ID:       "4b67ac88-e506-4a29-9d42-26e504e3435f",
			Userid:   200,
			TeamID:   108,
			Endpoint: "https://yourdomain.com/webhook",
			ClientID: "QVOQP06ZXC6CMGVFKB0ZT7J9Y7APOYGO",
			Events:   []string{"taskTagUpdated", "taskMoved"},
			Health: webhookHealth{
				Status:    "active",
				FailCount: 0,
			},
			Secret: "O94IM25S7PXBPYTMNXLLET230SRP0S89COR7B1YOJ2ZIE8WQNK5UUKEF26W0Z5GA",
		},
	}
	if !cmp.Equal(artifacts, want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestWebhooksService_CreateWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &WebhookRequest{
		Endpoint: "https://yourdomain.com/webhook",
		Events:   []string{"keyResultUpdated", "keyResultDeleted"},
	}

	mux.HandleFunc("/team/123/webhook", func(w http.ResponseWriter, r *http.Request) {
		v := new(WebhookRequest)
		json.NewDecoder(r.Body).Decode(v)

		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		testMethod(t, r, "POST")
		fmt.Fprint(w,
			`{
				"id": "4b67ac88-e506-4a29-9d42-26e504e3435e",
				"webhook": {
				  "id": "4b67ac88-e506-4a29-9d42-26e504e3435e",
				  "userid": 183,
				  "team_id": 108,
				  "endpoint": "https://yourdomain.com/webhook",
				  "client_id": "QVOQP06ZXC6CMGVFKB0ZT7J9Y7APOYGO",
				  "events": [
					"keyResultUpdated",
					"keyResultDeleted"
				  ],
				  "task_id": null,
				  "list_id": null,
				  "folder_id": null,
				  "space_id": null,
				  "health": {
					"status": "active",
					"fail_count": 0
				  },
				  "secret": "O94IM25S7PXBPYTMNXLLET230SRP0S89COR7B1YOJ2ZIE8WQNK5UUKEF26W0Z5GA"
				}
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Webhooks.CreateWebhook(ctx, 123, input)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := WebhookResponse{
		ID: "4b67ac88-e506-4a29-9d42-26e504e3435e",
		Webhook: Webhook{
			ID:       "4b67ac88-e506-4a29-9d42-26e504e3435e",
			Userid:   183,
			TeamID:   108,
			Endpoint: "https://yourdomain.com/webhook",
			ClientID: "QVOQP06ZXC6CMGVFKB0ZT7J9Y7APOYGO",
			Events:   []string{"keyResultUpdated", "keyResultDeleted"},
			Health: webhookHealth{
				Status:    "active",
				FailCount: 0,
			},
			Secret: "O94IM25S7PXBPYTMNXLLET230SRP0S89COR7B1YOJ2ZIE8WQNK5UUKEF26W0Z5GA",
		},
	}

	if !cmp.Equal(artifacts, &want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestWebhooksService_UpdateWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &WebhookRequest{
		Endpoint: "https://yourdomain.com/webhook",
		Events:   []string{"*"},
		Status:   "active",
	}

	mux.HandleFunc("/webhook/4b67ac88-e506-4a29-9d42-26e504e3435e", func(w http.ResponseWriter, r *http.Request) {
		v := new(WebhookRequest)
		json.NewDecoder(r.Body).Decode(v)

		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		testMethod(t, r, "PUT")
		fmt.Fprint(w,
			`{
				"id": "4b67ac88-e506-4a29-9d42-26e504e3435e",
				"webhook": {
				  "id": "4b67ac88-e506-4a29-9d42-26e504e3435e",
				  "userid": 183,
				  "team_id": 108,
				  "endpoint": "https://yourdomain.com/webhook",
				  "client_id": "QVOQP06ZXC6CMGVFKB0ZT7J9Y7APOYGO",
				  "events": [
					"taskCreated",
      				"taskUpdated",
      				"taskDeleted",
      				"taskPriorityUpdated",
      				"taskStatusUpdated",
      				"taskAssigneeUpdated",
      				"taskDueDateUpdated",
      				"taskTagUpdated",
      				"taskMoved",
      				"taskCommentPosted",
      				"taskCommentUpdated",
      				"taskTimeEstimateUpdated",
      				"taskTimeTrackedUpdated",
      				"listCreated",
      				"listUpdated",
      				"listDeleted",
      				"folderCreated",
      				"folderUpdated",
      				"folderDeleted",
      				"spaceCreated",
      				"spaceUpdated",
      				"spaceDeleted",
      				"goalCreated",
      				"goalUpdated",
      				"goalDeleted",
      				"keyResultCreated",
      				"keyResultUpdated",
      				"keyResultDeleted"
				  ],
				  "task_id": null,
				  "list_id": null,
				  "folder_id": null,
				  "space_id": null,
				  "health": {
					"status": "active",
					"fail_count": 0
				  },
				  "secret": "O94IM25S7PXBPYTMNXLLET230SRP0S89COR7B1YOJ2ZIE8WQNK5UUKEF26W0Z5GA"
				}
			}`,
		)
	})

	ctx := context.Background()
	artifacts, _, err := client.Webhooks.UpdateWebhook(ctx, "4b67ac88-e506-4a29-9d42-26e504e3435e", input)
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	want := WebhookResponse{
		ID: "4b67ac88-e506-4a29-9d42-26e504e3435e",
		Webhook: Webhook{
			ID:       "4b67ac88-e506-4a29-9d42-26e504e3435e",
			Userid:   183,
			TeamID:   108,
			Endpoint: "https://yourdomain.com/webhook",
			ClientID: "QVOQP06ZXC6CMGVFKB0ZT7J9Y7APOYGO",
			Events: []string{
				"taskCreated",
				"taskUpdated",
				"taskDeleted",
				"taskPriorityUpdated",
				"taskStatusUpdated",
				"taskAssigneeUpdated",
				"taskDueDateUpdated",
				"taskTagUpdated",
				"taskMoved",
				"taskCommentPosted",
				"taskCommentUpdated",
				"taskTimeEstimateUpdated",
				"taskTimeTrackedUpdated",
				"listCreated",
				"listUpdated",
				"listDeleted",
				"folderCreated",
				"folderUpdated",
				"folderDeleted",
				"spaceCreated",
				"spaceUpdated",
				"spaceDeleted",
				"goalCreated",
				"goalUpdated",
				"goalDeleted",
				"keyResultCreated",
				"keyResultUpdated",
				"keyResultDeleted",
			},
			Health: webhookHealth{
				Status:    "active",
				FailCount: 0,
			},
			Secret: "O94IM25S7PXBPYTMNXLLET230SRP0S89COR7B1YOJ2ZIE8WQNK5UUKEF26W0Z5GA",
		},
	}

	if !cmp.Equal(artifacts, &want) {
		t.Errorf("Actions.ListArtifacts returned %+v, want %+v", artifacts, want)
	}
}

func TestWebhooksService_DeleteWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/webhook/4b67ac88-e506-4a29-9d42-26e504e3435e", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, `{}`)
	})

	ctx := context.Background()
	res, err := client.Webhooks.DeleteWebhook(ctx, "4b67ac88-e506-4a29-9d42-26e504e3435e")
	if err != nil {
		t.Errorf("Actions.ListArtifacts returned error: %v", err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)

}
