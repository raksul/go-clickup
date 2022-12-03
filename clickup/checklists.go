package clickup

import (
	"context"
	"encoding/json"
	"fmt"
)

type ChecklistsService service

type ChecklistRequest struct {
	Name     string `json:"name"`
	Position int    `json:"position,omitempty"`
}

type ChecklistResponse struct {
	Checklist Checklist `json:"checklist"`
}

// TODO: Add parent.
type ChecklistItemRequest struct {
	Name     string `json:"name"`
	Assignee int    `json:"assignee,omitempty"`
	Resolved bool   `json:"resolved,omitempty"`
}

type Checklist struct {
	ID         string      `json:"id"`
	TaskID     string      `json:"task_id"`
	Name       string      `json:"name"`
	Orderindex json.Number `json:"orderindex"`
	Resolved   int         `json:"resolved"`
	Unresolved int         `json:"unresolved"`
	Items      []Item      `json:"items,omitempty"`
}

type Item struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Orderindex  json.Number   `json:"orderindex"`
	Assignee    User          `json:"assignee"`
	Resolved    bool          `json:"resolved"`
	Parent      interface{}   `json:"parent"`
	DateCreated string        `json:"date_created"`
	Children    []interface{} `json:"children"`
}

type ChecklistOptions struct {
	CustomTaskIDs string `url:"custom_task_ids,omitempty"`
	TeamID        int    `url:"team_id,omitempty"`
}

func (s *ChecklistsService) CreateChecklist(ctx context.Context, taskID string, opts *ChecklistOptions, checklist *ChecklistRequest) (*Checklist, *Response, error) {
	u := fmt.Sprintf("task/%v/checklist/", taskID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, checklist)
	if err != nil {
		return nil, nil, err
	}

	cr := new(ChecklistResponse)
	resp, err := s.client.Do(ctx, req, cr)
	if err != nil {
		return nil, resp, err
	}

	return &cr.Checklist, resp, nil
}

// Position is the zero-based index of the order you want the checklist to exist on the task. If you want the checklist to be in the first position, pass { "position": 0 }
func (s *ChecklistsService) EditChecklist(ctx context.Context, checklistID string, checklist *ChecklistRequest) (*Checklist, *Response, error) {
	u := fmt.Sprintf("checklist/%v", checklistID)
	req, err := s.client.NewRequest("PUT", u, checklist)
	if err != nil {
		return nil, nil, err
	}

	cr := new(ChecklistResponse)
	resp, err := s.client.Do(ctx, req, cr)
	if err != nil {
		return nil, resp, err
	}

	return &cr.Checklist, resp, nil
}

func (s *ChecklistsService) DeleteChecklist(ctx context.Context, checklistID string) (*Response, error) {
	u := fmt.Sprintf("checklist/%v", checklistID)
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

func (s *ChecklistsService) CreateChecklistItem(ctx context.Context, checklistID string, item *ChecklistItemRequest) (*Checklist, *Response, error) {
	u := fmt.Sprintf("checklist/%v/checklist_item", checklistID)
	req, err := s.client.NewRequest("POST", u, item)
	if err != nil {
		return nil, nil, err
	}

	cr := new(ChecklistResponse)
	resp, err := s.client.Do(ctx, req, cr)
	if err != nil {
		return nil, resp, err
	}

	return &cr.Checklist, resp, nil
}

// Parent is another checklist item that you want to nest the target checklist item underneath.
func (s *ChecklistsService) EditChecklistItem(ctx context.Context, checklistID string, checklistItemID string, checklist *ChecklistItemRequest) (*Checklist, *Response, error) {
	u := fmt.Sprintf("checklist/%v/checklist_item/%v", checklistID, checklistItemID)
	req, err := s.client.NewRequest("PUT", u, checklist)
	if err != nil {
		return nil, nil, err
	}

	cr := new(ChecklistResponse)
	resp, err := s.client.Do(ctx, req, cr)
	if err != nil {
		return nil, resp, err
	}

	return &cr.Checklist, resp, nil
}

func (s *ChecklistsService) DeleteChecklistItem(ctx context.Context, checklistID string, checklistItemID string) (*Response, error) {
	u := fmt.Sprintf("checklist/%v/checklist_item/%v", checklistID, checklistItemID)

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
