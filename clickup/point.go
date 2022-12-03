package clickup

import (
	"encoding/json"
)

type Point struct {
	Value json.Number

	IntVal   *int64
	FloatVal *float64
}

func (p *Point) MarshalJSON() ([]byte, error) {
	if p.IntVal != nil {
		return json.Marshal(p.IntVal)
	}

	if p.FloatVal != nil {
		return json.Marshal(p.FloatVal)
	}

	return json.Marshal(p.Value)
}

func (p *Point) UnmarshalJSON(b []byte) error {
	p.IntVal = nil
	p.FloatVal = nil

	var i int64
	var f float64
	if err := json.Unmarshal(b, &i); err == nil {
		p.IntVal = &i
	} else {
		if err = json.Unmarshal(b, &f); err == nil {
			p.FloatVal = &f
		} else {
			return err
		}
	}

	p.Value = json.Number(b)
	return nil
}
