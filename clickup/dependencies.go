package clickup

import (
	"context"
	"fmt"
)

type DependenciesService service

// To create a waiting on dependency, pass the property depends_on in the body.
// To create a blocking dependency, pass the property dependency_of.
// Both can not be passed in the same request.
type AddDependencyRequest struct {
	DependsOn    string `json:"depends_on,omitempty"`
	DependencyOf string `json:"dependency_of,omitempty"`
}

type AddDependencyOptions struct {
	CustomTaskIDs string `url:"custom_task_ids,omitempty"`
	TeamID        int    `url:"team_id,omitempty"`
}

//One and only one of depends_on or dependency_of must be passed in the query params.
type DeleteDependencyOptions struct {
	DependsOn     string `url:"depends_on,omitempty"`
	DependencyOf  string `url:"dependency_of,omitempty"`
	CustomTaskIDs string `url:"custom_task_ids,omitempty"`
	TeamID        int    `url:"team_id,omitempty"`
}

type TaskLinkResponse struct {
	Task Task `json:"task"`
}

type TaskLinkOptions struct {
	CustomTaskIDs string `url:"custom_task_ids,omitempty"`
	TeamID        int    `url:"team_id,omitempty"`
}

// To create a waiting on dependency, pass the property DependsOn in the body.
// To create a blocking dependency, pass the property DependencyOf.
// Both can not be passed in the same request.
func (s *DependenciesService) AddDependency(ctx context.Context, taskID string, adr *AddDependencyRequest, opts *AddDependencyOptions) (*Response, error) {
	u := fmt.Sprintf("task/%v/dependency", taskID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", u, adr)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// One and only one of DependsOn or DependencyOf must be passed in the query params.
func (s *DependenciesService) DeleteDependency(ctx context.Context, taskID string, opts *DeleteDependencyOptions) (*Response, error) {
	u := fmt.Sprintf("task/%v/dependency", taskID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

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

func (s *DependenciesService) AddTaskLink(ctx context.Context, taskID string, linksTo string, opts *TaskLinkOptions) (*Task, *Response, error) {
	u := fmt.Sprintf("task/%v/link/%v", taskID, linksTo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tlr := new(TaskLinkResponse)
	resp, err := s.client.Do(ctx, req, tlr)
	if err != nil {
		return nil, resp, err
	}

	return &tlr.Task, resp, nil
}

func (s *DependenciesService) DeleteTaskLink(ctx context.Context, taskID string, linksTo string, opts *TaskLinkOptions) (*Task, *Response, error) {
	u := fmt.Sprintf("task/%v/link/%v", taskID, linksTo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tlr := new(TaskLinkResponse)
	resp, err := s.client.Do(ctx, req, tlr)
	if err != nil {
		return nil, resp, err
	}

	return &tlr.Task, resp, nil
}
