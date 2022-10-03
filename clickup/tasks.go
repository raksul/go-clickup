package clickup

import (
	"context"
	"fmt"
	"strings"
)

type TasksService service

type GetTasksResponse struct {
	Tasks []Task `json:"tasks"`
}

type GetBulkTasksTimeInStatusResponse map[string]TasksInStatus

type TaskRequest struct {
	Name                      string                     `json:"name,omitempty"`
	Description               string                     `json:"description,omitempty"`
	Assignees                 []int                      `json:"assignees,omitempty"`
	Tags                      []string                   `json:"tags,omitempty"`
	Status                    string                     `json:"status,omitempty"`
	Priority                  int                        `json:"priority,omitempty"`
	DueDate                   int64                      `json:"due_date,omitempty"`
	DueDateTime               bool                       `json:"due_date_time,omitempty"`
	TimeEstimate              int                        `json:"time_estimate,omitempty"`
	StartDate                 int64                      `json:"start_date,omitempty"`
	StartDateTime             bool                       `json:"start_date_time,omitempty"`
	NotifyAll                 bool                       `json:"notify_all,omitempty"`
	Parent                    string                     `json:"parent,omitempty"`
	LinksTo                   string                     `json:"links_to,omitempty"`
	CheckRequiredCustomFields bool                       `json:"check_required_custom_fields,omitempty"`
	CustomFields              []CustomFieldInTaskRequest `json:"custom_fields,omitempty"`
}

type CustomFieldInTaskRequest struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

type Task struct {
	ID              string                 `json:"id"`
	CustomID        string                 `json:"custom_id"`
	Name            string                 `json:"name"`
	TextContent     string                 `json:"text_content"`
	Description     string                 `json:"description"`
	Status          TaskStatus             `json:"status"`
	Orderindex      string                 `json:"orderindex"`
	DateCreated     string                 `json:"date_created"`
	DateUpdated     string                 `json:"date_updated"`
	DateClosed      string                 `json:"date_closed"`
	Archived        bool                   `json:"archived"`
	Creator         User                   `json:"creator"`
	Assignees       []User                 `json:"assignees,omitempty"`
	Watchers        []User                 `json:"watchers,omitempty"`
	Checklists      []Checklist            `json:"checklists,omitempty"`
	Tags            []Tag                  `json:"tags,omitempty"`
	Parent          string                 `json:"parent"`
	Priority        TaskPriority           `json:"priority"`
	DueDate         string                 `json:"due_date,omitempty"`
	StartDate       string                 `json:"start_date,omitempty"`
	Points          int                    `json:"points,omitempty"`
	TimeEstimate    int64                  `json:"time_estimate"`
	CustomFields    []CustomField          `json:"custom_fields"`
	Dependencies    []Dependence           `json:"dependencies"`
	LinkedTasks     []LinkedTask           `json:"linked_tasks"`
	TeamID          string                 `json:"team_id"`
	URL             string                 `json:"url"`
	PermissionLevel string                 `json:"permission_level"`
	List            ListOfTaskBelonging    `json:"list"`
	Project         ProjectOfTaskBelonging `json:"project"`
	Folder          FolderOftaskBelonging  `json:"folder"`
	Space           SpaceOfTaskBelonging   `json:"space"`
}

type Dependence struct {
	TaskID      string `json:"task_id"`
	DependsOn   string `json:"depends_on"`
	Type        int    `json:"type"`
	DateCreated string `json:"date_created"`
	Userid      string `json:"userid"`
}

type LinkedTask struct {
	TaskID      string `json:"task_id"`
	LinkID      string `json:"link_id"`
	DateCreated string `json:"date_created"`
	Userid      string `json:"userid"`
}

type TasksInStatus struct {
	taskID        string              `json:"-"`
	CurrentStatus CurrentTaskStatus   `json:"current_status"`
	StatusHistory []TaskStatusHistory `json:"status_history"`
}

type TaskStatus struct {
	Status     string `json:"status"`
	Color      string `json:"color"`
	Type       string `json:"type"`
	Orderindex int    `json:"orderindex"`
}

type TaskPriority struct {
	Priority string `json:"priority"`
	Color    string `json:"color"`
}

type ListOfTaskBelonging struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Access bool   `json:"access"`
}

type ProjectOfTaskBelonging struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
	Access bool   `json:"access"`
}

type FolderOftaskBelonging struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
	Access bool   `json:"access"`
}

type SpaceOfTaskBelonging struct {
	ID string `json:"id"`
}

type CurrentTaskStatus struct {
	Status    string                     `json:"status"`
	Color     string                     `json:"color"`
	TotalTime CurrentTaskStatusTotalTime `json:"total_time"`
}

type CurrentTaskStatusTotalTime struct {
	ByMinute int    `json:"by_minute"`
	Since    string `json:"since"`
}

type TaskStatusHistory struct {
	Status     string                     `json:"status"`
	Color      string                     `json:"color"`
	Type       string                     `json:"type"`
	TotalTime  CurrentTaskStatusTotalTime `json:"total_time"`
	Orderindex int                        `json:"orderindex"`
}

// TODO: Implement custom field
type GetTasksOptions struct {
	Archived      bool     `url:"archived,omitempty"`
	Page          int      `url:"page,omitempty"`
	OrderBy       string   `url:"order_by,omitempty"`
	Reverse       bool     `url:"reverse,omitempty"`
	Subtasks      bool     `url:"subtasks,omitempty"`
	Statuses      []string `url:"statuses[],omitempty"`
	IncludeClosed bool     `url:"include_closed,omitempty"`
	Assignees     []string `url:"assignees[],omitempty"`
	DueDateGt     int64    `url:"due_date_gt,omitempty"`
	DueDateLt     int64    `url:"due_date_lt,omitempty"`
	DateCreatedGt int64    `url:"date_created_gt,omitempty"`
	DateCreatedLt int64    `url:"date_created_lt,omitempty"`
	DateUpdatedGt int64    `url:"date_updated_gt,omitempty"`
	DateUpdatedLt int64    `url:"date_updated_lt,omitempty"`
}

type GetTaskOptions struct {
	CustomTaskIDs   string `url:"custom_task_ids,omitempty"`
	TeamID          int    `url:"team_id,omitempty"`
	IncludeSubTasks bool   `url:"include_subtasks,omitempty"`
}

type GetBulkTasksTimeInStatusOptions struct {
	CustomTaskIDs bool `url:"custom_task_ids,omitempty"`
	TeamID        int  `url:"team_id,omitempty"`
}

func (s *TasksService) CreateTask(ctx context.Context, listID string, tr *TaskRequest) (*Task, *Response, error) {
	u := fmt.Sprintf("list/%s/task", listID)
	req, err := s.client.NewRequest("POST", u, tr)
	if err != nil {
		return nil, nil, err
	}

	task := new(Task)
	resp, err := s.client.Do(ctx, req, task)
	if err != nil {
		return nil, resp, err
	}

	return task, resp, nil
}

// FIXME: assignees add/rem
func (s *TasksService) UpdateTask(ctx context.Context, taskID string, opts *GetTaskOptions, tr *TaskRequest) (*Task, *Response, error) {
	u := fmt.Sprintf("task/%v/", taskID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("PUT", u, tr)
	if err != nil {
		return nil, nil, err
	}

	task := new(Task)
	resp, err := s.client.Do(ctx, req, task)
	if err != nil {
		return nil, resp, err
	}

	return task, resp, nil
}

func (s *TasksService) DeleteTask(ctx context.Context, taskID string, opts *GetTaskOptions) (*Response, error) {
	u := fmt.Sprintf("task/%v/", taskID)
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

// The maximum number of tasks returned in this response is 100.
// When you are paging this request, you should check list limit against the length of each response to determine if you are on the last page.
func (s *TasksService) GetTasks(ctx context.Context, listID string, opts *GetTasksOptions) ([]Task, *Response, error) {
	u := fmt.Sprintf("list/%s/task", listID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gtr := new(GetTasksResponse)
	resp, err := s.client.Do(ctx, req, gtr)
	if err != nil {
		return nil, resp, err
	}

	return gtr.Tasks, resp, nil
}

func (s *TasksService) GetTask(ctx context.Context, taskID string, opts *GetTaskOptions) (*Task, *Response, error) {
	u := fmt.Sprintf("task/%v/", taskID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	task := new(Task)
	resp, err := s.client.Do(ctx, req, task)
	if err != nil {
		return nil, resp, err
	}

	return task, resp, nil
}

// This request will always return paged responses.
// If you do not include the page parameter, it will return page 0.
// Each page includes 100 tasks.
func (s *TasksService) GetFilteredTeamTasks(ctx context.Context, teamID string, opts *GetTasksOptions) ([]Task, *Response, error) {
	u := fmt.Sprintf("team/%s/task", teamID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gtr := new(GetTasksResponse)
	resp, err := s.client.Do(ctx, req, gtr)
	if err != nil {
		return nil, resp, err
	}

	return gtr.Tasks, resp, nil
}

func (s *TasksService) GetTasksTimeInStatus(ctx context.Context, taskID string, opts *GetTaskOptions) (*TasksInStatus, *Response, error) {
	u := fmt.Sprintf("task/%v/time_in_status/", taskID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tis := new(TasksInStatus)
	resp, err := s.client.Do(ctx, req, tis)
	if err != nil {
		return nil, resp, err
	}

	return tis, resp, nil
}

// You must include at least 2 task_ids.
func (s *TasksService) GetBulkTasksTimeInStatus(ctx context.Context, taskIDs []string, opts *GetBulkTasksTimeInStatusOptions) ([]TasksInStatus, *Response, error) {

	if len(taskIDs) < 2 {
		return nil, nil, fmt.Errorf("you must include at least 2 task_ids. len: %d", len(taskIDs))
	}

	q := fmt.Sprintf("task_ids=%s", strings.Join(taskIDs, "&task_ids="))
	u, err := addOptions("task/bulk_time_in_status/task_ids/", opts)

	if strings.Contains(u, "?") {
		u += "&" + q
	} else {
		u += "?" + q
	}
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	gbtr := new(GetBulkTasksTimeInStatusResponse)
	resp, err := s.client.Do(ctx, req, gbtr)
	if err != nil {
		return nil, resp, err
	}

	var statuses []TasksInStatus
	for id, status := range *gbtr {
		s := TasksInStatus{
			taskID:        id,
			CurrentStatus: status.CurrentStatus,
			StatusHistory: status.StatusHistory,
		}
		statuses = append(statuses, s)
	}

	return statuses, resp, nil
}
