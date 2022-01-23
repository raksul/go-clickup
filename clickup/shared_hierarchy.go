package clickup

import (
	"context"
	"fmt"
)

type SharedHierarchyService service

type SharedHierarchyResponse struct {
	Shared Shared `json:"shared"`
}

type Shared struct {
	Tasks   []Task   `json:"tasks"`
	Lists   []List   `json:"lists"`
	Folders []Folder `json:"folders"`
}

// Returns all resources you have access to where you don't have access to its parent.
// For example, if you have a access to a shared task,
// but don't have access to its parent list, it will come back in this request.
func (s *SharedHierarchyService) SharedHierarchy(ctx context.Context, teamID int) (*Shared, *Response, error) {
	u := fmt.Sprintf("team/%v/shared", teamID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	shr := new(SharedHierarchyResponse)
	resp, err := s.client.Do(ctx, req, shr)
	if err != nil {
		return nil, resp, err
	}

	return &shr.Shared, resp, nil
}
