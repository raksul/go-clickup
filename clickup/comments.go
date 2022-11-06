package clickup

import (
	"context"
	"fmt"
)

type CommentsService service

type CommentRequest struct {
	CommentText string `json:"comment_text,omitempty"`
	Assignee    int    `json:"assignee,omitempty"`
	NotifyAll   bool   `json:"notify_all,omitempty"`
}

type UpdateCommentRequest struct {
	CommentText string `json:"comment_text,omitempty"`
	Assignee    int    `json:"assignee,omitempty"`
	Resolved    bool   `json:"resolved,omitempty"`
}

type CreateCommentResponse struct {
	ID     string `json:"id"`
	HistId string `json:"hist_id"`
	Date   *Date  `json:"date"`
}

type GetCommentsResponse struct {
	Comments []Comment `json:"comments"`
}

type TaskCommentOptions struct {
	CustomTaskIDs string `url:"custom_task_ids,omitempty"`
	TeamID        int    `url:"team_id,omitempty"`
}

type Comment struct {
	ID          string             `json:"id"`
	Comment     []CommentInComment `json:"comment"`
	CommentText string             `json:"comment_text"`
	User        User               `json:"user"`
	Resolved    bool               `json:"resolved"`
	Assignee    User               `json:"assignee,omitempty"`
	AssignedBy  User               `json:"assigned_by,omitempty"`
	Reactions   []Reaction         `json:"reactions,omitempty"`
	Date        string             `json:"date"`
}

type CommentInComment struct {
	Text string `json:"text"`
}

type Reaction struct {
	Reaction string `json:"reaction"`
	Date     string `json:"date"`
	User     User   `json:"user"`
}

// If NotifyAll is true, creation notifications will be sent to everyone including the creator of the comment.
func (s *CommentsService) CreateTaskComment(ctx context.Context, taskID string, opts *TaskCommentOptions, comment *CommentRequest) (*CreateCommentResponse, *Response, error) {
	u := fmt.Sprintf("task/%v/comment", taskID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, comment)
	if err != nil {
		return nil, nil, err
	}

	ccr := new(CreateCommentResponse)
	resp, err := s.client.Do(ctx, req, ccr)
	if err != nil {
		return nil, resp, err
	}

	return ccr, resp, nil
}

// If NotifyAll is true, creation notifications will be sent to everyone including the creator of the comment.
func (s *CommentsService) CreateChatViewComment(ctx context.Context, viewID string, comment *CommentRequest) (*CreateCommentResponse, *Response, error) {
	u := fmt.Sprintf("view/%v/comment", viewID)
	req, err := s.client.NewRequest("POST", u, comment)
	if err != nil {
		return nil, nil, err
	}

	ccr := new(CreateCommentResponse)
	resp, err := s.client.Do(ctx, req, ccr)
	if err != nil {
		return nil, resp, err
	}

	return ccr, resp, nil
}

// If NotifyAll is true, creation notifications will be sent to everyone including the creator of the comment.
func (s *CommentsService) CreateListComment(ctx context.Context, listID int, comment *CommentRequest) (*CreateCommentResponse, *Response, error) {
	u := fmt.Sprintf("list/%v/comment", listID)
	req, err := s.client.NewRequest("POST", u, comment)
	if err != nil {
		return nil, nil, err
	}

	ccr := new(CreateCommentResponse)
	resp, err := s.client.Do(ctx, req, ccr)
	if err != nil {
		return nil, resp, err
	}

	return ccr, resp, nil
}

func (s *CommentsService) GetTaskComments(ctx context.Context, taskID string, opts *TaskCommentOptions) ([]Comment, *Response, error) {
	u := fmt.Sprintf("task/%v/comment", taskID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gcr := new(GetCommentsResponse)
	resp, err := s.client.Do(ctx, req, gcr)
	if err != nil {
		return nil, resp, err
	}

	return gcr.Comments, resp, nil
}

func (s *CommentsService) GetChatViewComments(ctx context.Context, viewID string) ([]Comment, *Response, error) {
	u := fmt.Sprintf("view/%v/comment", viewID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gcr := new(GetCommentsResponse)
	resp, err := s.client.Do(ctx, req, gcr)
	if err != nil {
		return nil, resp, err
	}

	return gcr.Comments, resp, nil
}

func (s *CommentsService) GetListComments(ctx context.Context, listID int) ([]Comment, *Response, error) {
	u := fmt.Sprintf("list/%v/comment", listID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gcr := new(GetCommentsResponse)
	resp, err := s.client.Do(ctx, req, &gcr)
	if err != nil {
		return nil, resp, err
	}

	return gcr.Comments, resp, nil
}

func (s *CommentsService) UpdateComment(ctx context.Context, commentID int, comment *UpdateCommentRequest) (*Response, error) {
	u := fmt.Sprintf("comment/%v", commentID)
	req, err := s.client.NewRequest("PUT", u, comment)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *CommentsService) DeleteComment(ctx context.Context, commentID int) (*Response, error) {
	u := fmt.Sprintf("comment/%v", commentID)

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
