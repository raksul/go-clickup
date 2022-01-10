package clickup

import (
	"context"
	"encoding/json"
	"fmt"
)

type CustomFieldsService service

type CustomFieldResponse struct {
	Fields []CustomField `json:"fields"`
}

// TODO: Add type_config
type CustomField struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	DateCreated    string `json:"date_created"`
	HideFromGuests bool   `json:"hide_from_guests"`
}

type CustomFieldOptions struct {
	CustomTaskIDs bool `url:"custom_task_ids,omitempty"`
	TeamID        int  `url:"team_id,omitempty"`
}

func (s *CustomFieldsService) GetAccessibleCustomFields(ctx context.Context, listID string) ([]CustomField, *Response, error) {
	u := fmt.Sprintf("list/%s/field", listID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	cr := new(CustomFieldResponse)
	resp, err := s.client.Do(ctx, req, cr)
	if err != nil {
		return nil, resp, err
	}

	return cr.Fields, resp, nil
}

// The accessible fields can be found on the task object from the get task route. This is where you can retrieve the fieldID.
// If you set tasks, example is as follow.
// 	value := map[string]interface{}{
// 		"value": map[string]interface{}{
//	        "add": []string{"wmq3", "qt15"},
//    	    "rem": []string{"wxm7"},
//    	},
//	}
// Each value setting is placed at ClickUp API docs.
func (s *CustomFieldsService) SetCustomFieldValue(ctx context.Context, taskID string, fieldID string, value map[string]interface{}, opts *CustomFieldOptions) (*Response, error) {
	u := fmt.Sprintf("task/%s/field/%s", taskID, fieldID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	str, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", u, str)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// The accessible fields can be found on the task object from the get task route. This is where you can retrieve the fieldID.
func (s *CustomFieldsService) RemoveCustomFieldValue(ctx context.Context, taskID string, fieldID string, opts *CustomFieldOptions) (*Response, error) {
	u := fmt.Sprintf("task/%s/field/%s", taskID, fieldID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

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
