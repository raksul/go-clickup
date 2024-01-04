package clickup

import (
	"context"
	"fmt"
)

type CustomTaskTypesService service

// See https://clickup.com/api/clickupreference/operation/GetCustomItems/
type GetCustomTaskTypesResponse struct {
	CustomItems []CustomItem `json:"custom_items"` // Array of custom task types.
}

// See https://clickup.com/api/clickupreference/operation/GetCustomItems/
type CustomItem struct {
	Id          int32  `json:"id"`          // Custom task type ID.
	Name        string `json:"name"`        // Custom task type name.
	NamePlural  string `json:"name_plural"` // Custom task type plural name.
	Description string `json:"description"` // Custom task type description.
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
