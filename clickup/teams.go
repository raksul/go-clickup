package clickup

import (
	"context"
	"fmt"
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

type SeatsResponse struct {
	Seats
}

type Seats struct {
	Members Members `json:"members"`
	Guests  Guests  `json:"guests"`
}

type Members struct {
	FilledMembersSeats int `json:"filled_members_seats"`
	TotalMemberSeats   int `json:"total_member_seats"`
	EmptyMemberSeats   int `json:"empty_member_seats"`
}

type Guests struct {
	FilledGuestSeats int `json:"filled_guest_seats"`
	TotalGuestSeats  int `json:"total_guest_seats"`
	EmptyGuestSeats  int `json:"empty_guest_seats"`
}

type PlanResponse struct {
	Plan
}

type Plan struct {
	Id   int    `json:"plan_id"`
	Name string `json:"plan_name"`
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

// Teams is the legacy term for what are now called Workspaces in ClickUp.
// For compatablitly, the term team is still used in this API.
// This is NOT the new "Teams" feature which represents a group of users.
func (s *TeamsService) GetSeats(ctx context.Context, teamId string) (Seats, *Response, error) {
	u := fmt.Sprintf("team/%s/seats", teamId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return Seats{}, nil, err
	}

	sr := new(SeatsResponse)
	resp, err := s.client.Do(ctx, req, sr)
	if err != nil {
		return Seats{}, resp, err
	}

	return sr.Seats, resp, nil
}

// Teams is the legacy term for what are now called Workspaces in ClickUp.
// For compatablitly, the term team is still used in this API.
// This is NOT the new "Teams" feature which represents a group of users.
func (s *TeamsService) GetPlan(ctx context.Context, teamId string) (Plan, *Response, error) {
	u := fmt.Sprintf("team/%s/plan", teamId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return Plan{}, nil, err
	}

	pr := new(PlanResponse)
	resp, err := s.client.Do(ctx, req, pr)
	if err != nil {
		return Plan{}, resp, err
	}

	return pr.Plan, resp, nil
}
