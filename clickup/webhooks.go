package clickup

import (
	"context"
	"fmt"
)

type WebhooksService service

// You may filter the location of resources that get sent to a webhook
// by passing an optional space_id, folder_id, list_id, or task_id in the body of the request.
// Status is only available for update.
// Below are a list of events that you can subscribe
// to in order to listen to specific changes to your resources.
// You can also pass the wildcard * to listen to all events that are available.
// - taskCreated
// - taskUpdated
// - taskDeleted
// - taskPriorityUpdated
// - taskStatusUpdated
// - taskAssigneeUpdated
// - taskDueDateUpdated
// - taskTagUpdated
// - taskMoved
// - taskCommentPosted
// - taskCommentUpdated
// - taskTimeEstimateUpdated
// - taskTimeTrackedUpdated
// - listCreated
// - listUpdated
// - listDeleted
// - folderCreated
// - folderUpdated
// - folderDeleted
// - spaceCreated
// - spaceUpdated
// - spaceDeleted
// - goalCreated
// - goalUpdated
// - goalDeleted
// - keyResultCreated
// - keyResultUpdated
// - keyResultDeleted
type WebhookRequest struct {
	Endpoint string   `json:"endpoint"`
	Events   []string `json:"events"`
	Status   string   `json:"status,omitempty"`
	TaskID   string   `json:"task_id,omitempty"`
	ListID   int      `json:"list_id,omitempty"`
	FolderID string   `json:"folder_id,string"`
	SpaceID  string   `json:"space_id,omitempty"`
}

type WebhookResponse struct {
	ID      string  `json:"id"`
	Webhook Webhook `json:"webhook"`
}

type GetWebhooksResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

type Webhook struct {
	ID       string        `json:"id"`
	Userid   int           `json:"userid"`
	TeamID   int           `json:"team_id"`
	Endpoint string        `json:"endpoint"`
	ClientID string        `json:"client_id"`
	Events   []string      `json:"events"`
	TaskID   string        `json:"task_id,omitempty"`
	ListID   int           `json:"list_id,omitempty"`
	FolderID string        `json:"folder_id,string"`
	SpaceID  string        `json:"space_id,omitempty"`
	Health   webhookHealth `json:"health"`
	Secret   string        `json:"secret"`
}

type webhookHealth struct {
	Status    string `json:"status"`
	FailCount int    `json:"fail_count"`
}

func (s *WebhooksService) GetWebhook(ctx context.Context, teamID int) ([]Webhook, *Response, error) {
	u := fmt.Sprintf("team/%v/webhook", teamID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gwr := new(GetWebhooksResponse)
	resp, err := s.client.Do(ctx, req, gwr)
	if err != nil {
		return nil, resp, err
	}

	return gwr.Webhooks, resp, nil
}

func (s *WebhooksService) CreateWebhook(ctx context.Context, teamID int, webhookReq *WebhookRequest) (*WebhookResponse, *Response, error) {
	u := fmt.Sprintf("team/%v/webhook", teamID)
	req, err := s.client.NewRequest("POST", u, webhookReq)
	if err != nil {
		return nil, nil, err
	}

	wr := new(WebhookResponse)
	resp, err := s.client.Do(ctx, req, wr)
	if err != nil {
		return nil, resp, err
	}

	return wr, resp, nil
}

func (s *WebhooksService) UpdateWebhook(ctx context.Context, webhookID string, webhookReq *WebhookRequest) (*WebhookResponse, *Response, error) {
	u := fmt.Sprintf("webhook/%v", webhookID)
	req, err := s.client.NewRequest("PUT", u, webhookReq)
	if err != nil {
		return nil, nil, err
	}

	wr := new(WebhookResponse)
	resp, err := s.client.Do(ctx, req, wr)
	if err != nil {
		return nil, resp, err
	}

	return wr, resp, nil
}

func (s *WebhooksService) DeleteWebhook(ctx context.Context, webhookID string) (*Response, error) {
	u := fmt.Sprintf("webhook/%v", webhookID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
