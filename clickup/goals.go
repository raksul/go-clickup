package clickup

import (
	"context"
	"fmt"
)

type GoalsService service

type GetGoalsResponse struct {
	Goals   []Goal       `json:"goals"`
	Folders []GoalFolder `json:"folders"`
}

type GetGoalResponse struct {
	Goal Goal `json:"goal"`
}

type CreateGoalRequest struct {
	Name           string `json:"name"`
	DueDate        *Date  `json:"due_date"`
	Description    string `json:"description"`
	MultipleOwners bool   `json:"multiple_owners"`
	Owners         []int  `json:"owners"`
	Color          string `json:"color"`
}

type UpdateGoalRequest struct {
	Name        string `json:"name"`
	DueDate     *Date  `json:"due_date"`
	Description string `json:"description"`
	RemOwners   []int  `json:"rem_owners"`
	AddOwners   []int  `json:"add_owners"`
	Color       string `json:"color"`
}

type CreateKeyResultRequest struct {
	Name       string   `json:"name"`
	Owners     []int    `json:"owners"`
	Type       string   `json:"type"`
	StepsStart int      `json:"steps_start"`
	StepsEnd   int      `json:"steps_end"`
	Unit       string   `json:"unit"`
	TaskIds    []string `json:"task_ids"`
	ListIds    []string `json:"list_ids"`
}

type EditKeyResultRequest struct {
	StepsCurrent int    `json:"steps_current"`
	Note         string `json:"note"`
}

type KeyResultResponse struct {
	KeyResult KeyResult `json:"key_result"`
}

type GoalFolder struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	TeamID       string       `json:"team_id"`
	Private      bool         `json:"private"`
	DateCreated  string       `json:"date_created"`
	Creator      int          `json:"creator"`
	GoalCount    int          `json:"goal_count"`
	GroupMembers []GoalMember `json:"group_members"`
	Goals        []Goal       `json:"goals"`
}

type Goal struct {
	ID               string       `json:"id"`
	PrettyID         string       `json:"pretty_id"`
	Name             string       `json:"name"`
	TeamID           string       `json:"team_id"`
	Creator          int          `json:"creator"`
	Owner            GoalOwner    `json:"owner"`
	Color            string       `json:"color"`
	DateCreated      string       `json:"date_created"`
	StartDate        string       `json:"start_date"`
	DueDate          string       `json:"due_date"`
	Description      string       `json:"description"`
	Private          bool         `json:"private"`
	Archived         bool         `json:"archived"`
	MultipleOwners   bool         `json:"multiple_owners"`
	EditorToken      string       `json:"editor_token"`
	DateUpdated      string       `json:"date_updated"`
	LastUpdate       string       `json:"last_update"`
	FolderID         string       `json:"folder_id"`
	FolderAccess     bool         `json:"folder_access,omitempty"`
	Pinned           bool         `json:"pinned"`
	Owners           []GoalOwner  `json:"owners"`
	KeyResultCount   int          `json:"key_result_count"`
	Members          []GoalMember `json:"members"`
	GroupMembers     []GoalMember `json:"group_members"`
	PercentCompleted float64      `json:"percent_completed"`
}

type GoalOwner struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Color          string `json:"color"`
	ProfilePicture string `json:"profilePicture"`
	Initials       string `json:"initials"`
}

type GoalMember struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Color           string `json:"color"`
	PermissionLevel string `json:"permission_level"`
	ProfilePicture  string `json:"profilePicture,omitempty"`
	Initials        string `json:"initials"`
	IsCreator       bool   `json:"isCreator"`
}

type KeyResult struct {
	ID               string      `json:"id"`
	GoalID           string      `json:"goal_id"`
	Name             string      `json:"name"`
	Type             string      `json:"type"`
	Unit             string      `json:"unit"`
	Creator          int         `json:"creator"`
	DateCreated      string      `json:"date_created"`
	GoalPrettyID     string      `json:"goal_pretty_id"`
	PercentCompleted float64     `json:"percent_completed,omitempty"`
	Completed        bool        `json:"completed"`
	TaskIds          []string    `json:"task_ids"`
	SubcategoryIds   []string    `json:"subcategory_ids"`
	Owners           []GoalOwner `json:"owners"`
	LastAction       LastAction  `json:"last_action"`
}

type LastAction struct {
	ID           string `json:"id"`
	KeyResultID  string `json:"key_result_id"`
	Userid       int    `json:"userid"`
	DateModified string `json:"date_modified"`
	Note         string `json:"note"`
}

func (s *GoalsService) CreateGoal(ctx context.Context, teamID int, createGoalRequest *CreateGoalRequest) (*Goal, *Response, error) {
	u := fmt.Sprintf("team/%v/goal", teamID)
	req, err := s.client.NewRequest("POST", u, createGoalRequest)
	if err != nil {
		return nil, nil, err
	}

	ggr := new(GetGoalResponse)
	resp, err := s.client.Do(ctx, req, ggr)
	if err != nil {
		return nil, resp, err
	}

	return &ggr.Goal, resp, nil
}

func (s *GoalsService) UpdateGoal(ctx context.Context, goalID string, updateGoalRequest *UpdateGoalRequest) (*Goal, *Response, error) {
	u := fmt.Sprintf("goal/%v", goalID)
	req, err := s.client.NewRequest("PUT", u, updateGoalRequest)
	if err != nil {
		return nil, nil, err
	}

	ggr := new(GetGoalResponse)
	resp, err := s.client.Do(ctx, req, ggr)
	if err != nil {
		return nil, resp, err
	}

	return &ggr.Goal, resp, nil
}

func (s *GoalsService) DeleteGoal(ctx context.Context, goalID string) (resp *Response, err error) {
	u := fmt.Sprintf("goal/%v", goalID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.Do(ctx, req, nil)
	return resp, err
}

// Use includeCompleted to include new, in progress, and completed goals in the response.
func (s *GoalsService) GetGoals(ctx context.Context, teamID string, includeCompleted bool) ([]Goal, []GoalFolder, *Response, error) {
	u := fmt.Sprintf("team/%s/goal?include_completed=%T", teamID, includeCompleted)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	ggr := new(GetGoalsResponse)
	resp, err := s.client.Do(ctx, req, ggr)
	if err != nil {
		return nil, nil, resp, err
	}

	return ggr.Goals, ggr.Folders, resp, nil
}

func (s *GoalsService) GetGoal(ctx context.Context, goalID string) (*Goal, *Response, error) {
	u := fmt.Sprintf("goal/%s", goalID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ggr := new(GetGoalResponse)
	resp, err := s.client.Do(ctx, req, ggr)
	if err != nil {
		return nil, resp, err
	}

	return &ggr.Goal, resp, nil
}

// Key result types can be number, currency, boolean, percentage, or automatic.
// The task ID's array and list ID's array can be used to attach resources to the goal.
func (s *GoalsService) CreateKeyResult(ctx context.Context, goalID string, createKeyResultRequest *CreateKeyResultRequest) (*KeyResult, *Response, error) {
	u := fmt.Sprintf("goal/%v/key_result", goalID)
	req, err := s.client.NewRequest("POST", u, createKeyResultRequest)
	if err != nil {
		return nil, nil, err
	}

	krr := new(KeyResultResponse)
	resp, err := s.client.Do(ctx, req, krr)
	if err != nil {
		return nil, resp, err
	}

	return &krr.KeyResult, resp, nil
}

func (s *GoalsService) EditKeyResult(ctx context.Context, keyResultID string, editKeyResultRequest *EditKeyResultRequest) (*KeyResult, *Response, error) {
	u := fmt.Sprintf("key_result/%v", keyResultID)
	req, err := s.client.NewRequest("PUT", u, editKeyResultRequest)
	if err != nil {
		return nil, nil, err
	}

	krr := new(KeyResultResponse)
	resp, err := s.client.Do(ctx, req, krr)
	if err != nil {
		return nil, resp, err
	}

	return &krr.KeyResult, resp, nil
}

func (s *GoalsService) DeleteKeyResult(ctx context.Context, keyResultID string) (resp *Response, err error) {
	u := fmt.Sprintf("key_result/%v", keyResultID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.Do(ctx, req, nil)
	return resp, err
}
