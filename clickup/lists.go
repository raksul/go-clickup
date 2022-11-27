package clickup

import (
	"context"
	"encoding/json"
	"fmt"
)

type ListsService service

type GetListsResponse struct {
	Lists []List `json:"Lists"`
}

type ListRequest struct {
	Name        string `json:"name"`
	Content     string `json:"content"`
	DueDate     *Date  `json:"due_date"`
	DueDateTime bool   `json:"due_date_time"`
	Priority    int    `json:"priority"`
	Assignee    int    `json:"assignee"`
	Status      string `json:"status"`
}

type List struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Orderindex json.Number `json:"orderindex"`
	Content    string      `json:"content"`
	Status     struct {
		Status    string `json:"status"`
		Color     string `json:"color"`
		HideLabel bool   `json:"hide_label"`
	} `json:"status"`
	Priority struct {
		Priority string `json:"priority"`
		Color    string `json:"color"`
	} `json:"priority"`
	Assignee      User        `json:"assignee,omitempty"`
	TaskCount     json.Number `json:"task_count"`
	DueDate       string      `json:"due_date"`
	DueDateTime   bool        `json:"due_date_time"`
	StartDate     string      `json:"start_date"`
	StartDateTime bool        `json:"start_date_time"`
	Folder        struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"folder"`
	Space struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"space"`
	Statuses []struct {
		Status     string      `json:"status"`
		Orderindex json.Number `json:"orderindex"`
		Color      string      `json:"color"`
		Type       string      `json:"type"`
	} `json:"statuses"`
	InboundAddress  string `json:"inbound_address"`
	Archived        bool   `json:"archived"`
	PermissionLevel string `json:"permission_level"`
}

// Assignee is a userid of the assignee to be added to this task.
// Priority is an integer mapping as 1 : Urgent, 2 : High, 3 : Normal, 4 : Low.
// The status included in the body of this request refers to the List color rather than the task Statuses available in the List.
func (s *ListsService) CreateList(ctx context.Context, folderID string, listRequest *ListRequest) (List, *Response, error) {
	urlStr := fmt.Sprintf("folder/%v/list", folderID)
	req, err := s.client.NewRequest("POST", urlStr, listRequest)
	var (
		list List
	)
	if err != nil {
		return list, nil, err
	}

	resp, err := s.client.Do(ctx, req, &list)
	if err != nil {
		return list, resp, err
	}

	return list, resp, nil
}

// Assignee is a userid of the assignee to be added to this task.
// Priority is an integer mapping as 1 : Urgent, 2 : High, 3 : Normal, 4 : Low.
// The status included in the body of this request refers to the List color rather than the task Statuses available in the List.
func (s *ListsService) CreateFolderlessList(ctx context.Context, spaceID int, listRequest *ListRequest) (List, *Response, error) {
	urlStr := fmt.Sprintf("space/%v/list", spaceID)
	req, err := s.client.NewRequest("POST", urlStr, listRequest)
	var (
		list List
	)
	if err != nil {
		return list, nil, err
	}

	resp, err := s.client.Do(ctx, req, &list)
	if err != nil {
		return list, resp, err
	}

	return list, resp, nil
}

// Only pass the properties you want to update. It is unnessary to pass the entire list object.
// Assignee is a userid of the assignee to be added to this task.
// Priority is an integer mapping as 1 : Urgent, 2 : High, 3 : Normal, 4 : Low.
// You can set a List color using status as shown in Create List and Create Folderless List,
// or use unset_status as shown in the body of the example request below to clear the List color.
func (s *ListsService) UpdateList(ctx context.Context, listID string, listRequest *ListRequest) (List, *Response, error) {
	urlStr := fmt.Sprintf("list/%v", listID)
	req, err := s.client.NewRequest("PUT", urlStr, listRequest)
	var (
		list List
	)
	if err != nil {
		return list, nil, err
	}

	resp, err := s.client.Do(ctx, req, &list)
	if err != nil {
		return list, resp, err
	}

	return list, resp, nil
}

func (s *ListsService) DeleteList(ctx context.Context, listID string) (*Response, error) {
	urlStr := fmt.Sprintf("list/%v", listID)
	req, err := s.client.NewRequest("DELETE", urlStr, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// The status included in the body of the response refers to the List color rather than the task Statuses available in the List.
func (s *ListsService) GetLists(ctx context.Context, folderID string, archived bool) ([]List, *Response, error) {
	urlStr := fmt.Sprintf("folder/%s/list?archived=%v", folderID, archived)
	req, err := s.client.NewRequest("GET", urlStr, nil)

	var (
		lists []List
		lRes  GetListsResponse
	)

	if err != nil {
		return lists, nil, err
	}

	resp, err := s.client.Do(ctx, req, &lRes)
	if err != nil {
		return lists, resp, err
	}
	lists = lRes.Lists

	return lists, resp, nil
}

// The status included in the body of the response refers to the List color rather than the task Statuses available in the List.
func (s *ListsService) GetFolderlessLists(ctx context.Context, spaceID string, archived bool) ([]List, *Response, error) {
	urlStr := fmt.Sprintf("space/%s/list?archived=%v", spaceID, archived)
	req, err := s.client.NewRequest("GET", urlStr, nil)

	var (
		lists []List
		lRes  GetListsResponse
	)

	if err != nil {
		return lists, nil, err
	}

	resp, err := s.client.Do(ctx, req, &lRes)
	if err != nil {
		return lists, resp, err
	}
	lists = lRes.Lists

	return lists, resp, nil
}

// The status included in the body of the response refers to the List color rather than the task Statuses available in the List.
func (s *ListsService) GetList(ctx context.Context, listID string) (List, *Response, error) {
	urlStr := fmt.Sprintf("list/%s", listID)

	req, err := s.client.NewRequest("GET", urlStr, nil)

	var (
		list List
	)

	if err != nil {
		return list, nil, err
	}

	resp, err := s.client.Do(ctx, req, &list)
	if err != nil {
		return list, resp, err
	}
	return list, resp, nil
}

func (s *ListsService) AddTaskToList(ctx context.Context, listID string, taskID string) (*Response, error) {
	urlStr := fmt.Sprintf("list/%v/task/%v", listID, taskID)
	req, err := s.client.NewRequest("POST", urlStr, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *ListsService) RemoveTaskFromList(ctx context.Context, listID string, taskID string) (*Response, error) {
	urlStr := fmt.Sprintf("list/%v/task/%v", listID, taskID)
	req, err := s.client.NewRequest("DELETE", urlStr, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
