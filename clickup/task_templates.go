package clickup

import (
	"context"
	"fmt"
)

type TaskTemplatesService service

type GetTaskTemplateResponse struct {
	Templates []Template `json:"templates"`
}

type CreateTaskFromTemplateRequest struct {
	Name string `json:"name"`
}

type CreateTaskFromTemplateResponse struct {
	ID   string `json:"id"`
	Task Task   `json:"task"`
}

type Template struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// To page task templates, pass the page number you wish to fetch.
func (s *TaskTemplatesService) GetTaskTemplates(ctx context.Context, teamID int, page int) ([]Template, *Response, error) {
	u := fmt.Sprintf("team/%v/taskTemplate?page=%v", teamID, page)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gttr := new(GetTaskTemplateResponse)
	resp, err := s.client.Do(ctx, req, gttr)
	if err != nil {
		return nil, resp, err
	}

	return gttr.Templates, resp, nil
}

func (s *TaskTemplatesService) CreateTaskFromTemplate(ctx context.Context, listID string, templateID string, taskReq CreateTaskFromTemplateRequest) (*Task, *Response, error) {
	u := fmt.Sprintf("list/%v/taskTemplate/%v", listID, templateID)

	req, err := s.client.NewRequest("POST", u, taskReq)
	if err != nil {
		return nil, nil, err
	}

	ctftr := new(CreateTaskFromTemplateResponse)
	resp, err := s.client.Do(ctx, req, ctftr)
	if err != nil {
		return nil, resp, err
	}

	return &ctftr.Task, resp, nil
}
