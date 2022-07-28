package clickup

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type CustomFieldsService service

type CustomFieldResponse struct {
	Fields []CustomField `json:"fields"`
}

type CustomField struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Type           string      `json:"type"`
	TypeConfig     interface{} `json:"type_config"`
	DateCreated    string      `json:"date_created"`
	HideFromGuests bool        `json:"hide_from_guests"`
	Value          interface{} `json:"value"`
}

type CurrencyValue struct {
	Value      float64
	TypeConfig CurrencyTypeConfig
}

type CurrencyTypeConfig struct {
	Precision    float64 `json:"precision"`
	CurrencyType string  `json:"currency_type"`
	Default      float64 `json:"default"`
}

type EmojiValue struct {
	Value      int
	TypeConfig EmojiTypeConfig
}

type EmojiTypeConfig struct {
	CodePoint string `json:"code_point"`
	Count     int    `json:"count"`
}

type LocationValue struct {
	Latitude         float64
	Longitude        float64
	FormattedAddress string
	PlaceID          string
}

type AutomaticProgressValue struct {
	PercentCompleted float64
	TypeConfig       AutomaticProgressTypeConfig
}

type AutomaticProgressTypeConfig struct {
	SubtaskRollup bool    `json:"subtask_rollup"`
	CompleteOn    float64 `json:"complete_on"`
	Tracking      struct {
		Subtasks         bool `json:"subtasks"`
		AssignedComments bool `json:"assigned_comments"`
		Checklist        bool `json:"checklists"`
	} `json:"tracking"`
}

type ManualProgressValue struct {
	PercentCompleted float64
	Current          int64
	TypeConfig       ManualProgressTypeConfig
}

type ManualProgressTypeConfig struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type TaskValue struct {
	ID         string
	Access     bool
	Color      string
	CustomId   string
	CustomType interface{}
	Deleted    bool
	Name       string
	Status     string
	TeamId     string
	URL        string
}

type TasksValue []TaskValue

type UserValue struct {
	ID             json.Number
	Username       string
	Email          string
	Color          string
	Initials       string
	ProfilePicture string
}

type UsersValue []UserValue

// TODO: Set a concrete type for each field.
type AttachmentValue struct {
	ID                  string
	Version             json.Number
	Date                json.Number
	Title               string
	Extension           string
	ThumbnailSmall      string
	ThumbnailMedium     string
	ThumbnailLarge      string
	URL                 string
	Orientation         interface{}
	Type                json.Number
	Hidden              bool
	Size                json.Number
	ParentId            string
	ParentCommentType   interface{}
	ParentCommentParent interface{}
	EmailData           interface{}
	UrlWHost            string
	UrlWQuery           string
	Source              json.Number
	IsFolder            interface{}
	Mimetype            interface{}
	TotalComments       json.Number
	ResolvedComments    json.Number
	Deleted             bool
	User                UserValue
}

type AttachmentsValue []AttachmentValue

type LabelsValue struct {
	Values     []LabelOption
	TypeConfig LabelsTypeConfig
}

type LabelsTypeConfig struct {
	Options []LabelOption
}

type LabelOption struct {
	ID    string
	Label string
	Color string
}

type DropDownValue struct {
	Value      DropDownOption
	TypeConfig DropDownTypeConfig
}

type DropDownTypeConfig struct {
	Default     float64
	Placeholder string
	Options     []DropDownOption
}

type DropDownOption struct {
	ID         string
	OrderIndex int
	Name       string
	Color      string
}

func (cf CustomField) GetValue() interface{} {
	switch cf.Type {
	case "url", "email", "phone", "text", "short_text":
		str, ok := getStringValue(cf.Value)
		if !ok {
			return nil
		}

		return str
	case "number", "formula":
		num, ok := getFloatValue(cf.Value)
		if !ok {
			return nil
		}

		return num
	case "currency":
		num, ok := getFloatValue(cf.Value)
		if !ok {
			return nil
		}

		tc := CurrencyTypeConfig{}
		if ok := getStructValue(cf.TypeConfig, &tc); !ok {
			return nil
		}

		return CurrencyValue{
			Value:      num,
			TypeConfig: tc,
		}
	case "emoji":
		num, ok := getIntValue(cf.Value)
		if !ok {
			return nil
		}

		tc := EmojiTypeConfig{}
		if ok := getStructValue(cf.TypeConfig, &tc); !ok {
			return nil
		}

		return EmojiValue{
			Value:      int(num),
			TypeConfig: tc,
		}
	case "date":
		date, ok := getDateValue(cf.Value)
		if !ok {
			return nil
		}

		return date
	case "checkbox":
		b, ok := getBoolValue(cf.Value)
		if !ok {
			return nil
		}

		return b
	case "location":
		loc, ok := getLocationValue(cf.Value)
		if !ok {
			return nil
		}

		return loc
	case "automatic_progress":
		prog, ok := getAutomaticProgressValue(cf.Value, cf.TypeConfig)
		if !ok {
			return nil
		}

		return prog
	case "manual_progress":
		prog, ok := getManualProgressValue(cf.Value, cf.TypeConfig)
		if !ok {
			return nil
		}

		return prog
	case "tasks":
		v := TasksValue{}
		if ok := getStructValue(cf.Value, &v); !ok {
			return nil
		}

		return v
	case "users":
		v := UsersValue{}
		if ok := getStructValue(cf.Value, &v); !ok {
			return nil
		}

		return v
	case "attachment":
		v := AttachmentsValue{}
		if ok := getStructValue(cf.Value, &v); !ok {
			return nil
		}

		return v
	case "drop_down":
		v, ok := getDropDownValue(cf.Value, cf.TypeConfig)
		if !ok {
			return nil
		}

		return v
	case "labels":
		v, ok := getLabelsValue(cf.Value, cf.TypeConfig)
		if !ok {
			return nil
		}

		return v
	}

	return cf.Value
}

func getStringValue(v interface{}) (string, bool) {
	str, ok := v.(string)
	if !ok {
		return "", false
	}

	return str, true
}

func getFloatValue(v interface{}) (float64, bool) {
	switch v := v.(type) {
	case int:
		return float64(v), true
	case float64:
		return v, true
	case string:
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, false
		}

		return num, true
	}

	return 0, false
}

func getIntValue(v interface{}) (int64, bool) {
	switch v := v.(type) {
	case int:
		return int64(v), true
	case string:
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, false
		}

		return num, true
	}

	return 0, false
}

func getDateValue(v interface{}) (time.Time, bool) {
	num, ok := getIntValue(v)
	if !ok {
		return time.Time{}, false
	}

	return time.UnixMilli(num), true
}

func getBoolValue(v interface{}) (bool, bool) {
	str, ok := v.(string)
	if !ok {
		return false, false
	}

	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, false
	}

	return b, true
}

func getLocationValue(v interface{}) (LocationValue, bool) {
	m, ok := v.(map[string]interface{})
	if !ok {
		return LocationValue{}, false
	}

	loc := LocationValue{}

	for k, v := range m {
		switch k {
		case "location":
			v, ok := v.(map[string]interface{})
			if !ok {
				return LocationValue{}, false
			}

			lat, ok := func() (float64, bool) {
				lat, ok := v["lat"]
				if !ok {
					return 0, false
				}

				num, ok := lat.(float64)
				if !ok {
					return 0, false
				}

				return num, true
			}()
			if !ok {
				return LocationValue{}, false
			}

			loc.Latitude = lat

			lng, ok := func() (float64, bool) {
				lng, ok := v["lng"]
				if !ok {
					return 0, false
				}

				num, ok := lng.(float64)
				if !ok {
					return 0, false
				}

				return num, true
			}()
			if !ok {
				return LocationValue{}, false
			}

			loc.Longitude = lng
		case "formatted_address":
			str, ok := v.(string)
			if !ok {
				return LocationValue{}, false
			}

			loc.FormattedAddress = str
		case "place_id":
			str, ok := v.(string)
			if !ok {
				return LocationValue{}, false
			}

			loc.PlaceID = str
		}
	}

	return loc, true
}

func getAutomaticProgressValue(v interface{}, typeConfig interface{}) (AutomaticProgressValue, bool) {
	m, ok := v.(map[string]interface{})
	if !ok {
		return AutomaticProgressValue{}, false
	}

	percentStr, ok := m["percent_complete"]
	if !ok {
		return AutomaticProgressValue{}, false
	}

	percent, ok := percentStr.(float64)
	if !ok {
		return AutomaticProgressValue{}, false
	}

	tc := AutomaticProgressTypeConfig{}
	if ok := getStructValue(typeConfig, &tc); !ok {
		return AutomaticProgressValue{}, false
	}

	return AutomaticProgressValue{
		PercentCompleted: percent,
		TypeConfig:       tc,
	}, true
}

func getManualProgressValue(v interface{}, typeConfig interface{}) (ManualProgressValue, bool) {
	m, ok := v.(map[string]interface{})
	if !ok {
		return ManualProgressValue{}, false
	}

	percent, ok := func() (float64, bool) {
		percent, ok := m["percent_completed"]
		if !ok {
			return 0, false
		}

		num, ok := percent.(float64)
		if !ok {
			return 0, false
		}

		return num, true
	}()
	if !ok {
		return ManualProgressValue{}, false
	}

	current, ok := func() (int64, bool) {
		current, ok := m["current"]
		if !ok {
			return 0, false
		}

		str, ok := current.(string)
		if !ok {
			return 0, false
		}

		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0, false
		}

		return i, true
	}()
	if !ok {
		return ManualProgressValue{}, false
	}

	tc := ManualProgressTypeConfig{}
	if ok := getStructValue(typeConfig, &tc); !ok {
		return ManualProgressValue{}, false
	}

	return ManualProgressValue{
		PercentCompleted: percent,
		Current:          current,
		TypeConfig:       tc,
	}, true
}

func getDropDownValue(v, typeConfig interface{}) (DropDownValue, bool) {
	ind, ok := v.(float64)
	if !ok {
		return DropDownValue{}, false
	}

	i := int(ind)

	tc := DropDownTypeConfig{}
	if ok := getStructValue(typeConfig, &tc); !ok {
		return DropDownValue{}, false
	}

	ddv := DropDownValue{
		Value:      DropDownOption{},
		TypeConfig: tc,
	}

	for _, option := range tc.Options {
		if option.OrderIndex == i {
			ddv.Value = option
		}
	}

	return ddv, true
}

func getLabelsValue(v, typeConfig interface{}) (LabelsValue, bool) {
	arr, ok := v.([]interface{})
	if !ok {
		return LabelsValue{}, false
	}

	ids := make([]string, len(arr))
	for i, v := range arr {
		ids[i] = v.(string)
	}

	tc := LabelsTypeConfig{}
	if ok := getStructValue(typeConfig, &tc); !ok {
		return LabelsValue{}, false
	}

	lv := LabelsValue{
		Values:     make([]LabelOption, len(ids)),
		TypeConfig: tc,
	}

	for i, id := range ids {
		for _, option := range tc.Options {
			if option.ID == id {
				lv.Values[i] = option
			}
		}
	}

	return lv, true
}

func getStructValue(src, dst interface{}) bool {
	b, err := json.Marshal(src)
	if err != nil {
		return false
	}

	if err := json.Unmarshal(b, &dst); err != nil {
		return false
	}

	return true
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

	req, err := s.client.NewRequest("POST", u, value)
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
