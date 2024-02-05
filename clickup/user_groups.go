package clickup

import (
	"context"
	"fmt"
)

type UserGroupsService service

type GetUserGroupsResponse struct {
	UserGroups []UserGroup `json:"groups"`
}

type UserGroup struct {
	ID          string          `json:"id"`
	TeamID      string          `json:"team_id"`
	UserID      int             `json:"userid"`
	Name        string          `json:"name"`
	Handle      string          `json:"handle"`
	DateCreated string          `json:"date_created"`
	Initials    string          `json:"initials"`
	Members     []GroupMember   `json:"members"`
	Avatar      UserGroupAvatar `json:"avatar"`
}

type UserGroupAvatar struct {
	AttachmentId *string `json:"attachment_id"`
	Color        *string `json:"color"`
	Source       *string `json:"source"`
	Icon         *string `json:"icon"`
}

type GroupMember struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Color          string `json:"color"`
	Initials       string `json:"initials"`
	ProfilePicture string `json:"profilePicture"`
}

type UserGroupRequest struct {
	Name    string `json:"name"`
	Members []int  `json:"members"`
}

type GetUserGroupsOptions struct {
	TeamID   string   `url:"team_id,omitempty"`
	GroupIDs []string `url:"group_ids,omitempty"`
}

type UpdateUserGroupMember struct {
	Add    []int `json:"add"`
	Remove []int `json:"rem"`
}

type UpdateUserGroupRequest struct {
	Name    string                `json:"name"`
	Handle  string                `json:"handle"`
	Members UpdateUserGroupMember `json:"members"`
}

type CreateUserGroupRequest struct {
	Name    string `json:"name"`
	Members []int  `json:"add"`
}

func (s *UserGroupsService) GetUserGroups(ctx context.Context, opts *GetUserGroupsOptions) ([]UserGroup, *Response, error) {
	u, err := addOptions("group", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	response := new(GetUserGroupsResponse)
	resp, err := s.client.Do(ctx, req, response)
	if err != nil {
		return nil, resp, err
	}

	return response.UserGroups, resp, nil
}

func (s *UserGroupsService) CreateUserGroup(ctx context.Context, teamID string, createUserGroupRequest *CreateUserGroupRequest) (*UserGroup, *Response, error) {
	u := fmt.Sprintf("team/%v/group", teamID)
	req, err := s.client.NewRequest("POST", u, createUserGroupRequest)
	if err != nil {
		return nil, nil, err
	}

	group := new(UserGroup)
	resp, err := s.client.Do(ctx, req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

func (s *UserGroupsService) UpdateUserGroup(ctx context.Context, groupID string, updateUserGroupRequest *UpdateUserGroupRequest) (*UserGroup, *Response, error) {
	u := fmt.Sprintf("group/%v", groupID)
	req, err := s.client.NewRequest("PUT", u, updateUserGroupRequest)
	if err != nil {
		return nil, nil, err
	}

	group := new(UserGroup)
	resp, err := s.client.Do(ctx, req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

func (s *UserGroupsService) DeleteUserGroup(ctx context.Context, groupID string) (*Response, error) {
	u := fmt.Sprintf("group/%v", groupID)
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
