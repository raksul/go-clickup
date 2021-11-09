package clickup

import (
	"context"
	"fmt"
)

type TagsService service

type GetTagsResponse struct {
	Tags []Tag `json:"tags"`
}

type TagRequest struct {
	Tag Tag `json:"tag"`
}

type Tag struct {
	Name    string `json:"name"`
	TagFg   string `json:"tag_fg"`
	TagBg   string `json:"tag_bg"`
	Creator int    `json:"creator,omitempty"`
}

type TagOptions struct {
	CustomTaskIDs string `url:"custom_task_ids,omitempty"`
	TeamID        int    `url:"team_id,omitempty"`
}

func (s *TagsService) GetTags(ctx context.Context, spaceID string) ([]Tag, *Response, error) {
	u := fmt.Sprintf("space/%s/tag", spaceID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gtr := new(GetTagsResponse)
	resp, err := s.client.Do(ctx, req, gtr)
	if err != nil {
		return nil, resp, err
	}

	return gtr.Tags, resp, nil
}

func (s *TagsService) CreateSpaceTag(ctx context.Context, spaceID string, tagReq *TagRequest) (*Response, error) {
	u := fmt.Sprintf("space/%s/tag", spaceID)
	req, err := s.client.NewRequest("POST", u, tagReq)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *TagsService) EditSpaceTag(ctx context.Context, spaceID string, tagName string, tagReq *TagRequest) (*Response, error) {
	u := fmt.Sprintf("space/%s/tag/%s", spaceID, tagName)
	req, err := s.client.NewRequest("PUT", u, tagReq)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *TagsService) DeleteSpaceTag(ctx context.Context, spaceID string, tagName string) (*Response, error) {
	u := fmt.Sprintf("space/%s/tag/%s", spaceID, tagName)
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

func (s *TagsService) AddTagToTask(ctx context.Context, taskID string, tagName string, opts *TagOptions) (*Response, error) {
	u := fmt.Sprintf("task/%s/tag/%s", taskID, tagName)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *TagsService) RemoveTagToTask(ctx context.Context, taskID string, tagName string, opts *TagOptions) (*Response, error) {
	u := fmt.Sprintf("task/%s/tag/%s", taskID, tagName)
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
