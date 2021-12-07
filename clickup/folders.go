package clickup

import (
	"context"
	"encoding/json"
	"fmt"
)

type FoldersService service

type GetFoldersResponse struct {
	Folders []Folder `json:"Folders"`
}

type GetFolderResponse struct {
	Folder Folder `json:"Folder"`
}

type FolderRequest struct {
	Name string `json:"name"`
}

type Folder struct {
	ID               string                  `json:"id"`
	Name             string                  `json:"name"`
	Orderindex       int                     `json:"orderindex"`
	OverrideStatuses bool                    `json:"override_statuses"`
	Hidden           bool                    `json:"hidden"`
	Space            SpaceOfFolderBelonging  `json:"space"`
	TaskCount        json.Number             `json:"task_count"`
	Archived         bool                    `json:"archived"`
	Statuses         []interface{}           `json:"statuses"`
	Lists            []ListOfFolderBelonging `json:"lists"`
	PermissionLevel  string                  `json:"permission_level"`
}

type SpaceOfFolderBelonging struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Access bool   `json:"access,omitempty"`
}

type ListOfFolderBelonging struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Orderindex int         `json:"orderindex"`
	Status     interface{} `json:"status"`
	Priority   interface{} `json:"priority"`
	Assignee   interface{} `json:"assignee"`
	TaskCount  json.Number `json:"task_count"`
	DueDate    interface{} `json:"due_date"`
	StartDate  interface{} `json:"start_date"`
	Space      struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"space"`
	Archived         bool        `json:"archived"`
	OverrideStatuses interface{} `json:"override_statuses"`
	Statuses         []struct {
		ID         string `json:"id"`
		Status     string `json:"status"`
		Orderindex int    `json:"orderindex"`
		Color      string `json:"color"`
		Type       string `json:"type"`
	} `json:"statuses"`
	PermissionLevel string `json:"permission_level,omitempty"`
}

func (s *FoldersService) CreateFolder(ctx context.Context, spaceID int, folderRequest *FolderRequest) (*Folder, *Response, error) {
	u := fmt.Sprintf("space/%v/folder", spaceID)
	req, err := s.client.NewRequest("POST", u, folderRequest)
	if err != nil {
		return nil, nil, err
	}

	folder := new(Folder)
	resp, err := s.client.Do(ctx, req, folder)
	if err != nil {
		return nil, resp, err
	}

	return folder, resp, nil
}

func (s *FoldersService) UpdateFolder(ctx context.Context, folderID int, folderRequest *FolderRequest) (*Folder, *Response, error) {
	u := fmt.Sprintf("folder/%v", folderID)
	req, err := s.client.NewRequest("PUT", u, folderRequest)
	if err != nil {
		return nil, nil, err
	}

	folder := new(Folder)
	resp, err := s.client.Do(ctx, req, folder)
	if err != nil {
		return nil, resp, err
	}

	return folder, resp, nil
}

func (s *FoldersService) DeleteFolder(ctx context.Context, folderID int) (*Response, error) {
	u := fmt.Sprintf("folder/%v", folderID)
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

func (s *FoldersService) GetFolders(ctx context.Context, spaceID string, archived bool) ([]Folder, *Response, error) {
	u := fmt.Sprintf("space/%s/folder?archived=%v", spaceID, archived)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gfr := new(GetFoldersResponse)
	resp, err := s.client.Do(ctx, req, gfr)
	if err != nil {
		return nil, resp, err
	}

	return gfr.Folders, resp, nil
}

func (s *FoldersService) GetFolder(ctx context.Context, folderID string) (*Folder, *Response, error) {
	u := fmt.Sprintf("folder/%s", folderID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	folder := new(Folder)
	resp, err := s.client.Do(ctx, req, folder)
	if err != nil {
		return nil, resp, err
	}

	return folder, resp, nil
}
