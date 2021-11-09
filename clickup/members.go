package clickup

import (
	"context"
	"fmt"
)

type MembersService service

type GetMembersResponse struct {
	Members []Member `json:"Members"`
}

type Member struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Color          string `json:"color"`
	ProfilePicture string `json:"profilePicture"`
	Initials       string `json:"initials"`
	Role           int    `json:"role"`
}

func (s *MembersService) GetTaskMembers(ctx context.Context, taskID string) ([]Member, *Response, error) {
	u := fmt.Sprintf("task/%s/member", taskID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gmr := new(GetMembersResponse)
	resp, err := s.client.Do(ctx, req, gmr)
	if err != nil {
		return nil, resp, err
	}

	return gmr.Members, resp, nil
}

func (s *MembersService) GetListMembers(ctx context.Context, listID string) ([]Member, *Response, error) {
	u := fmt.Sprintf("list/%s/member", listID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gmr := new(GetMembersResponse)
	resp, err := s.client.Do(ctx, req, gmr)
	if err != nil {
		return nil, resp, err
	}

	return gmr.Members, resp, nil
}
