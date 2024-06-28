package clickup

import (
	"context"
	"fmt"
)

type TimeTrackingsService service

type GetTimeTrackingResponse struct {
	Data TimeTrackingData `json:"data"`
}

type CreateTimeTrackingResponse struct {
	Data TimeTrackingData `json:"data"`
}

// See https://clickup.com/api/clickupreference/operation/Createatimeentry/
type TimeTrackingRequest struct {
	Description string            `json:"description,omitempty"`
	Tags        []TimeTrackingTag `json:"tags,omitempty"`
	Start       int64             `json:"start"`
	End         int64             `json:"end,omitempty"`
	Stop        int64             `json:"stop,omitempty"`
	Billable    bool              `json:"billable,omitempty"`
	Duration    int32             `json:"duration"`
	Assignee    int               `json:"assignee,omitempty"`
	Tid         string            `json:"tid,omitempty"`
}

type TimeTrackingTag struct {
	Name    string `json:"name"`
	TagBg   string `json:"tag_bg"`
	TagFg   string `json:"tag_fg"`
	Creator int    `json:"creator"`
}

type TimeTrackingData struct {
	ID           string            `json:"id"`
	Wid          string            `json:"wid"`
	User         User              `json:"user"`
	Billable     bool              `json:"billable"`
	Start        int               `json:"start"`
	End          string            `json:"end"`
	Duration     int               `json:"duration"`
	Description  string            `json:"description"`
	Source       string            `json:"source"`
	At           int               `json:"at"`
	IsLocked     bool              `json:"is_locked"`
	TaskLocation TaskLocation      `json:"task_location"`
	Task         Task              `json:"task"`
	Tags         []TimeTrackingTag `json:"tags"`
	TaskURL      string            `json:"task_url"`
}

type TaskLocation struct {
	ListID     int    `json:"list_id"`
	FolderID   int    `json:"folder_id"`
	SpaceID    int    `json:"space_id"`
	ListName   string `json:"list_name"`
	FolderName string `json:"folder_name"`
	SpaceName  string `json:"space_name"`
}

type CreateTimeTrackingOptions struct {
	CustomTaskIDs bool `url:"custom_task_ids,omitempty"`
	TeamID        int  `url:"team_id,omitempty"`
}

func (s *TimeTrackingsService) CreateTimeTracking(ctx context.Context, teamID string, opts *CreateTimeTrackingOptions, ttr *TimeTrackingRequest) (*CreateTimeTrackingResponse, *Response, error) {
	u := fmt.Sprintf("team/%s/time_entries", teamID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, ttr)
	if err != nil {
		return nil, nil, err
	}

	timeTracking := new(CreateTimeTrackingResponse)
	resp, err := s.client.Do(ctx, req, timeTracking)
	if err != nil {
		return nil, resp, err
	}

	return timeTracking, resp, nil
}

func (s *TimeTrackingsService) GetTimeTracking(ctx context.Context, teamID string, timerID string, opts *CreateTimeTrackingOptions, ttr *TimeTrackingRequest) (*GetTimeTrackingResponse, *Response, error) {
	u := fmt.Sprintf("team/%s/time_entries/%s", teamID, timerID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, ttr)
	if err != nil {
		return nil, nil, err
	}

	getTimeTrackingResponse := new(GetTimeTrackingResponse)
	resp, err := s.client.Do(ctx, req, getTimeTrackingResponse)
	if err != nil {
		return nil, resp, err
	}

	return getTimeTrackingResponse, resp, nil
}
