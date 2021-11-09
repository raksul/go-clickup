package clickup

import (
	"context"
	"fmt"
)

type SpacesService service

type GetSpacesResponse struct {
	Spaces []Space `json:"spaces"`
}

type SpaceRequest struct {
	Name              string   `json:"name"`
	MultipleAssignees bool     `json:"multiple_assignees"`
	Features          Features `json:"features"`
}

type DueDates struct {
	Enabled            bool `json:"enabled"`
	StartDate          bool `json:"start_date"`
	RemapDueDates      bool `json:"remap_due_dates"`
	RemapClosedDueDate bool `json:"remap_closed_due_date"`
}

type TimeTracking struct {
	Enabled bool `json:"enabled"`
}

type Tags struct {
	Enabled bool `json:"enabled"`
}

type TimeEstimates struct {
	Enabled     bool `json:"enabled"`
	Rollup      bool `json:"rollup"`
	PerAssignee bool `json:"per_assignee"`
}

type Checklists struct {
	Enabled bool `json:"enabled"`
}

type CustomFields struct {
	Enabled bool `json:"enabled"`
}

type RemapDependencies struct {
	Enabled bool `json:"enabled"`
}

type DependencyWarning struct {
	Enabled bool `json:"enabled"`
}

type Portfolios struct {
	Enabled bool `json:"enabled"`
}

type Features struct {
	DueDates          DueDates          `json:"due_dates"`
	TimeTracking      TimeTracking      `json:"time_tracking"`
	Tags              Tags              `json:"tags"`
	TimeEstimates     TimeEstimates     `json:"time_estimates"`
	Checklists        Checklists        `json:"checklists"`
	CustomFields      CustomFields      `json:"custom_fields"`
	RemapDependencies RemapDependencies `json:"remap_dependencies"`
	DependencyWarning DependencyWarning `json:"dependency_warning"`
	Portfolios        Portfolios        `json:"portfolios"`
}

type Space struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Private  bool   `json:"private"`
	Statuses []struct {
		ID         string `json:"id"`
		Status     string `json:"status"`
		Type       string `json:"type"`
		Orderindex int    `json:"orderindex"`
		Color      string `json:"color"`
	} `json:"statuses"`
	MultipleAssignees bool `json:"multiple_assignees"`
	Features          struct {
		DueDates DueDates `json:"due_dates"`
		Sprints  struct {
			Enabled bool `json:"enabled"`
		} `json:"sprints"`
		TimeTracking TimeTracking `json:"time_tracking"`
		Points       struct {
			Enabled bool `json:"enabled"`
		} `json:"points"`
		CustomItems struct {
			Enabled bool `json:"enabled"`
		} `json:"custom_items"`
		Tags            Tags          `json:"tags"`
		TimeEstimates   TimeEstimates `json:"time_estimates"`
		CheckUnresolved struct {
			Enabled    bool        `json:"enabled"`
			Subtasks   bool        `json:"subtasks"`
			Checklists interface{} `json:"checklists"`
			Comments   interface{} `json:"comments"`
		} `json:"check_unresolved"`
		Zoom struct {
			Enabled bool `json:"enabled"`
		} `json:"zoom"`
		Milestones struct {
			Enabled bool `json:"enabled"`
		} `json:"milestones"`
		RemapDependencies RemapDependencies `json:"remap_dependencies"`
		DependencyWarning DependencyWarning `json:"dependency_warning"`
		MultipleAssignees struct {
			Enabled bool `json:"enabled"`
		} `json:"multiple_assignees"`
		Emails struct {
			Enabled bool `json:"enabled"`
		} `json:"emails"`
	} `json:"features"`
	Archived bool `json:"archived"`
}

func (s *SpacesService) CreateSpace(ctx context.Context, teamID int, spaceRequest *SpaceRequest) (*Space, *Response, error) {
	u := fmt.Sprintf("team/%v/space", teamID)
	req, err := s.client.NewRequest("POST", u, spaceRequest)
	if err != nil {
		return nil, nil, err
	}

	space := new(Space)
	resp, err := s.client.Do(ctx, req, space)
	if err != nil {
		return nil, resp, err
	}

	return space, resp, nil
}

func (s *SpacesService) UpdateSpace(ctx context.Context, spaceID int, spaceRequest *SpaceRequest) (*Space, *Response, error) {
	u := fmt.Sprintf("space/%v", spaceID)
	req, err := s.client.NewRequest("PUT", u, spaceRequest)
	if err != nil {
		return nil, nil, err
	}

	space := new(Space)
	resp, err := s.client.Do(ctx, req, space)
	if err != nil {
		return nil, resp, err
	}

	return space, resp, nil
}

func (s *SpacesService) DeleteSpace(ctx context.Context, spaceID int) (*Response, error) {
	u := fmt.Sprintf("space/%v", spaceID)
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

func (s *SpacesService) GetSpaces(ctx context.Context, teamID string) ([]Space, *Response, error) {
	u := fmt.Sprintf("team/%s/space", teamID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gsr := new(GetSpacesResponse)
	resp, err := s.client.Do(ctx, req, gsr)
	if err != nil {
		return nil, resp, err
	}

	return gsr.Spaces, resp, nil
}

func (s *SpacesService) GetSpace(ctx context.Context, spaceID string) (*Space, *Response, error) {
	u := fmt.Sprintf("space/%s", spaceID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	space := new(Space)
	resp, err := s.client.Do(ctx, req, space)
	if err != nil {
		return nil, resp, err
	}

	return space, resp, nil
}
