package clickup

import (
	"context"
	"fmt"
)

type ViewsService service

type View struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Parent struct {
		ID   string `json:"id"`
		Type int    `json:"type"`
	} `json:"parent"`
	Grouping struct {
		Field     string   `json:"field"`
		Dir       int      `json:"dir"`
		Collapsed []string `json:"collapsed"`
		Ignore    bool     `json:"ignore"`
	} `json:"grouping"`
	Filters struct {
		Op     string `json:"op"`
		Fields []struct {
			Field string `json:"field"`
			Op    string `json:"op"`
			Idx   int    `json:"idx"`
		} `json:"fields"`
		Search             string `json:"search"`
		SearchCustomFields bool   `json:"search_custom_fields"`
		SearchDescription  bool   `json:"search_description"`
		SearchName         bool   `json:"search_name"`
		ShowClosed         bool   `json:"show_closed"`
	} `json:"filters"`
	Columns struct {
		Fields []struct {
			Field  string `json:"field"`
			Idx    int    `json:"idx"`
			Width  int    `json:"width"`
			Hidden bool   `json:"hidden"`
			Name   string `json:"name"`
		} `json:"fields"`
	} `json:"columns"`
	TeamSidebar struct {
		AssignedComments bool `json:"assigned_comments"`
		UnassignedTasks  bool `json:"unassigned_tasks"`
	} `json:"team_sidebar"`
	Settings struct {
		ShowTaskLocations      bool `json:"show_task_locations"`
		ShowSubtasks           int  `json:"show_subtasks"`
		ShowSubtaskParentNames bool `json:"show_subtask_parent_names"`
		ShowClosedSubtasks     bool `json:"show_closed_subtasks"`
		ShowAssignees          bool `json:"show_assignees"`
		ShowImages             bool `json:"show_images"`
		ShowTimer              bool `json:"show_timer"`
		MeComments             bool `json:"me_comments"`
		MeSubtasks             bool `json:"me_subtasks"`
		MeChecklists           bool `json:"me_checklists"`
		ShowEmptyStatuses      bool `json:"show_empty_statuses"`
		AutoWrap               bool `json:"auto_wrap"`
		TimeInStatusView       int  `json:"time_in_status_view"`
	} `json:"settings"`
	DateCreated string `json:"date_created"`
	Creator     int    `json:"creator"`
	Visibility  string `json:"visibility"`
	Protected   bool   `json:"protected"`
	Orderindex  int    `json:"orderindex"`
}

type ViewResponse struct {
	View View `json:"view"`
}

type GetViewsResponse struct {
	Views []View `json:"views"`
}

type GetViewTasksResponse struct {
	Tasks    []Task `json:"tasks"`
	LastPage bool   `json:"last_page"`
}

type ViewType int

const (
	TeamView ViewType = iota
	SpaceView
	FolderView
	ListView
)

func (v ViewType) String() string {
	switch v {
	case TeamView:
		return "team"
	case SpaceView:
		return "space"
	case FolderView:
		return "folder"
	case ListView:
		return "list"
	default:
		return "INVALID_TYPE"
	}
}

func (s *ViewsService) CreateViewOf(ctx context.Context, viewType ViewType, id string, view map[string]interface{}) (*View, *Response, error) {
	t := viewType.String()
	if t == "INVALID_TYPE" {
		return nil, nil, fmt.Errorf("invalid view type")
	}

	u := fmt.Sprintf("%v/%v/view", t, id)

	req, err := s.client.NewRequest("POST", u, view)
	if err != nil {
		return nil, nil, err
	}

	vr := new(ViewResponse)
	resp, err := s.client.Do(ctx, req, vr)
	if err != nil {
		return nil, resp, err
	}

	return &vr.View, resp, nil
}

func (s *ViewsService) GetViewsOf(ctx context.Context, viewType ViewType, id string) ([]View, *Response, error) {
	t := viewType.SelectedType()
	if t == "INVALID_TYPE" {
		return nil, nil, fmt.Errorf("invalid view type")
	}

	u := fmt.Sprintf("%v/%v/view", t, id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gvr := new(GetViewsResponse)
	resp, err := s.client.Do(ctx, req, gvr)
	if err != nil {
		return nil, resp, err
	}

	return gvr.Views, resp, nil
}

func (s *ViewsService) GetView(ctx context.Context, viewID string) (*View, *Response, error) {
	u := fmt.Sprintf("view/%v", viewID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	vr := new(ViewResponse)
	resp, err := s.client.Do(ctx, req, vr)
	if err != nil {
		return nil, resp, err
	}

	return &vr.View, resp, nil
}

// This request will always return paged responses.
// Each page includes 30 tasks.
func (s *ViewsService) GetViewTasks(ctx context.Context, viewID string, page int) ([]Task, bool, *Response, error) {
	u := fmt.Sprintf("view/%v/task?page=%v", viewID, page)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, false, nil, err
	}

	gvtr := new(GetViewTasksResponse)
	resp, err := s.client.Do(ctx, req, gvtr)
	if err != nil {
		return nil, false, resp, err
	}

	return gvtr.Tasks, gvtr.LastPage, resp, nil
}

func (s *ViewsService) UpdateView(ctx context.Context, viewID string, value map[string]interface{}) (*View, *Response, error) {
	u := fmt.Sprintf("view/%v", viewID)

	req, err := s.client.NewRequest("PUT", u, value)
	if err != nil {
		return nil, nil, err
	}

	vr := new(ViewResponse)
	resp, err := s.client.Do(ctx, req, vr)
	if err != nil {
		return nil, resp, err
	}

	return &vr.View, resp, nil
}

func (s *ViewsService) DeleteView(ctx context.Context, viewID string) (*Response, error) {
	u := fmt.Sprintf("view/%v", viewID)

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
