package clickup

import (
	"context"
	"fmt"
)

type CustomTaskTypesService service

// See https://clickup.com/api/clickupreference/operation/GetCustomItems/
type GetCustomTaskTypesResponse struct {
	CustomItems []CustomItem `json:"custom_items,omitempty"` // Array of custom task types.
}

// See https://clickup.com/api/clickupreference/operation/GetCustomItems/
type CustomItem struct {
	Id          int32  `json:"id,omitempty"`          // Custom task type ID.
	Name        string `json:"name,omitempty"`        // Custom task type name.
	NamePlural  string `json:"name_plural,omitempty"` // Custom task type plural name.
	Description string `json:"description,omitempty"` // Custom task type description.

	// Not documented in API explorer
	Avatar CustomItemAvatar `json:"avatar,omitempty"` // Custom task icon data.
}

// Not documented in API explorer. Comments are observations.
type CustomItemAvatar struct {
	Source string `json:"source,omitempty"` // null (ClickUp Milestone Glyph), fas (Font Awesome Solid), fab (Font Awesome Brands).
	Value  string `json:"value,omitempty"`  // null is for ClickUp Glyphs, e.g., Task and Milestone.
}

// See https://clickup.com/api/clickupreference/operation/GetCustomItems/
func (s *CustomTaskTypesService) GetCustomTaskTypes(ctx context.Context, teamId string) ([]CustomItem, *Response, error) {
	u := fmt.Sprintf("team/%s/custom_item", teamId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gcttr := new(GetCustomTaskTypesResponse)
	resp, err := s.client.Do(ctx, req, gcttr)
	if err != nil {
		return nil, resp, err
	}

	return gcttr.CustomItems, resp, nil
}
