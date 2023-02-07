package clickup

import (
	"context"
)

type UserGroupsService service

type GetUserGroupsResponse struct {
	UserGroups []UserGroup `json:"groups"`
}

type UserGroup struct {
	ID          string        `json:"id"`
	TeamID      string        `json:"team_id"`
	UserID      int           `json:"userid"`
	Name        string        `json:"name"`
	Handle      string        `json:"handle"`
	DateCreated string        `json:"date_created"`
	Initials    string        `json:"initials"`
	Members     []GroupMember `json:"members"`
	Avatar      interface{}   `json:"avatar"`
}

type GroupMember struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Color          string `json:"color"`
	Initials       string `json:"initials"`
	ProfilePicture string `json:"profilePicture"`
}

type GetUserGroupsOptions struct {
	TeamID   string   `url:"team_id,omitempty"`
	GroupIDs []string `url:"group_ids,omitempty"`
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

	gugr := new(GetUserGroupsResponse)
	resp, err := s.client.Do(ctx, req, gugr)
	if err != nil {
		return nil, resp, err
	}

	return gugr.UserGroups, resp, nil
}
