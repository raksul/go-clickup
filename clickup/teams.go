package clickup

import (
	"context"
)

type TeamsService service

type GetTeamsResponse struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Color   string       `json:"color"`
	Avatar  interface{}  `json:"avatar"`
	Members []TeamMember `json:"members"`
}

type TeamMember struct {
	User      TeamUser  `json:"user"`
	InvitedBy InvitedBy `json:"invited_by,omitempty"`
}

type TeamUser struct {
	ID             int         `json:"id"`
	Username       string      `json:"username"`
	Email          string      `json:"email"`
	Color          string      `json:"color"`
	ProfilePicture string      `json:"profilePicture"`
	Initials       string      `json:"initials"`
	Role           int         `json:"role"`
	CustomRole     interface{} `json:"custom_role"`
	LastActive     string      `json:"last_active"`
	DateJoined     string      `json:"date_joined"`
	DateInvited    string      `json:"date_invited"`
}

type InvitedBy struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Color          string `json:"color"`
	Email          string `json:"email"`
	Initials       string `json:"initials"`
	ProfilePicture string `json:"profilePicture"`
}

// Teams is the legacy term for what are now called Workspaces in ClickUp.
// For compatablitly, the term team is still used in this API.
// This is NOT the new "Teams" feature which represents a group of users.
func (s *TeamsService) GetTeams(ctx context.Context) ([]Team, *Response, error) {
	req, err := s.client.NewRequest("GET", "team", nil)
	if err != nil {
		return nil, nil, err
	}

	gtr := new(GetTeamsResponse)
	resp, err := s.client.Do(ctx, req, gtr)
	if err != nil {
		return nil, resp, err
	}

	return gtr.Teams, resp, nil
}
