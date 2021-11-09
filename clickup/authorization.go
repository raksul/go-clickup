package clickup

import (
	"context"
)

type AuthorizationService service

type GetAuthorizedUserResponse struct {
	User User `json:"user"`
}

type GetAuthorizedTeamsResponse struct {
	Team Team `json:"teams"`
}

type User struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Color             string `json:"color"`
	ProfilePicture    string `json:"profilePicture,omitempty"`
	Initials          string `json:"initials"`
	WeekStartDay      int    `json:"week_start_day,omitempty"`
	GlobalFontSupport bool   `json:"global_font_support"`
	Timezone          string `json:"timezone"`
}

// Get the user that belongs to this token.
func (s *AuthorizationService) GetAuthorizedUser(ctx context.Context) (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "user", nil)
	if err != nil {
		return nil, nil, err
	}

	gaur := new(GetAuthorizedUserResponse)
	resp, err := s.client.Do(ctx, req, gaur)
	if err != nil {
		return nil, resp, err
	}

	return &gaur.User, resp, nil
}

// Get the authorized teams for this token.
func (s *AuthorizationService) GetAuthorizedTeams(ctx context.Context) ([]Team, *Response, error) {
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
