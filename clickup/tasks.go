package clickup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
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
	DueDate                   *Date                      `json:"due_date,omitempty"`
	DueDateTime               bool                       `json:"due_date_time,omitempty"`
	TimeEstimate              int                        `json:"time_estimate,omitempty"`
	StartDate                 *Date                      `json:"start_date,omitempty"`
	StartDateTime             bool                       `json:"start_date_time,omitempty"`
	NotifyAll                 bool                       `json:"notify_all,omitempty"`
	Parent                    string                     `json:"parent,omitempty"`
	LinksTo                   string                     `json:"links_to,omitempty"`
	CheckRequiredCustomFields bool                       `json:"check_required_custom_fields,omitempty"`
	CustomFields              []CustomFieldInTaskRequest `json:"custom_fields,omitempty"`
}

type CustomFieldInTaskRequest struct {
	ID    string      `json:"id"`
	Value interface{} `json:"value"`
}

type Task struct {
	ID              string                 `json:"id"`
	CustomID        string                 `json:"custom_id"`
	Name            string                 `json:"name"`
	TextContent     string                 `json:"text_content"`
	Description     string                 `json:"description"`
	Status          TaskStatus             `json:"status"`
	Orderindex      json.Number            `json:"orderindex"`
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
	DueDate         *Date                  `json:"due_date,omitempty"`
	StartDate       string                 `json:"start_date,omitempty"`
	Points          Point                  `json:"points,omitempty"`
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
	Attachments     []TaskAttachment       `json:"attachments"`
}

type TaskAttachment struct {
	ID               string `json:"id"`
	Date             string `json:"date"`
	Title            string `json:"title"`
	Type             int    `json:"type"`
	Source           int    `json:"source"`
	Version          int    `json:"version"`
	Extension        string `json:"extension"`
	ThumbnailSmall   string `json:"thumbnail_small"`
	ThumbnailMedium  string `json:"thumbnail_medium"`
	ThumbnailLarge   string `json:"thumbnail_large"`
	IsFolder         bool   `json:"is_folder"`
	Mimetype         string `json:"mimetype"`
	Hidden           bool   `json:"hidden"`
	ParentId         string `json:"parent_id"`
	Size             int    `json:"size"`
	TotalComments    int    `json:"total_comments"`
	ResolvedComments int    `json:"resolved_comments"`
	User             User   `json:"user"`
	Deleted          bool   `json:"deleted"`
	Orientation      string `json:"orientation"`
	Url              string `json:"url"`
	EmailData        string `json:"email_data"`
	UrlWQuery        string `json:"url_w_query"`
	UrlWHost         string `json:"url_w_host"`
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
	Status     string      `json:"status"`
	Color      string      `json:"color"`
	Type       string      `json:"type"`
	Orderindex json.Number `json:"orderindex"`
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
	Orderindex json.Number                `json:"orderindex"`
}

type GetTasksOptions struct {
	Archived      bool                          `url:"archived,omitempty"`
	Page          int                           `url:"page,omitempty"`
	OrderBy       string                        `url:"order_by,omitempty"`
	Reverse       bool                          `url:"reverse,omitempty"`
	Subtasks      bool                          `url:"subtasks,omitempty"`
	Statuses      []string                      `url:"statuses[],omitempty"`
	IncludeClosed bool                          `url:"include_closed,omitempty"`
	Assignees     []string                      `url:"assignees[],omitempty"`
	Tags          []string                      `url:"tags[],omitempty"`
	DueDateGt     *Date                         `url:"due_date_gt,omitempty"`
	DueDateLt     *Date                         `url:"due_date_lt,omitempty"`
	DateCreatedGt *Date                         `url:"date_created_gt,omitempty"`
	DateCreatedLt *Date                         `url:"date_created_lt,omitempty"`
	DateUpdatedGt *Date                         `url:"date_updated_gt,omitempty"`
	DateUpdatedLt *Date                         `url:"date_updated_lt,omitempty"`
	CustomFields  CustomFieldsInGetTasksRequest `url:"custom_fields,omitempty"`
}

// CustomFieldsInGetTasksRequest is used to filter tasks using Custom Fields for GetTasks
type CustomFieldsInGetTasksRequest []CustomFieldInGetTasksRequest

type CustomFieldInGetTasksRequest struct {
	FieldId  string                               `url:"field_id"`
	Operator CustomFieldInGetTasksRequestOperator `url:"operator"`
	Value    []string                             `url:"value,omitempty,comma"`
}

// CustomFieldInGetTasksRequestOperator Values are found here https://clickup.com/api/developer-portal/filtertasks/
type CustomFieldInGetTasksRequestOperator int

const (
	Equals               CustomFieldInGetTasksRequestOperator = iota //=
	LessThan                                                         // <
	LessThanOrEqualTo                                                // <=
	GreaterThan                                                      // > (greater than)
	GreaterThanOrEqualTo                                             // > = (greater than or equal to)
	NotEqualTo                                                       // != (not equal to)
	IsNull                                                           // IS NULL (is not set)
	ISNotNull                                                        // IS NOT NULL (is set)
	Range                                                            // (is between)
	Any                                                              // (contains any matching criteria)
	All                                                              // (contains all matching criteria)
	NotAny                                                           // (does not contain any mathching criteria)
	NotAll                                                           // (does not contain all of the matching criteria)
)

func (c CustomFieldInGetTasksRequestOperator) String() string {
	return [...]string{"=", "<", "<=", ">", ">=", "!=", "IS NULL", "IS NOT NULL", "RANGE", "ANY", "ALL", "NOT ANY", "NOT ALL"}[c]
}

func (cfs CustomFieldsInGetTasksRequest) EncodeValues(key string, v *url.Values) error {
	sb := strings.Builder{}
	sb.WriteString("[")
	sep := ""
	for _, c := range cfs {
		sb.WriteString(sep + "{\"field_id\":\"" + c.FieldId + "\",\"operator\":\"" + c.Operator.String() + "\"")
		if len(c.Value) == 1 {
			sb.WriteString(",\"value\":\"" + c.Value[0] + "\"")
		} else if len(c.Value) > 1 {
			sb.WriteString(",\"value\":[\"" + strings.Join(c.Value, "\",\"") + "\"]")
		}
		sb.WriteString("}")
		sep = ","
	}

	sb.WriteString("]")
	v.Set(key, sb.String())
	return nil
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
