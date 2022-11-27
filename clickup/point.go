package clickup

import (
	"encoding/json"
	"strconv"
)

type Point struct {
	Value json.Number

	IntVal   *int64
	FloatVal *float64
}

func (p *Point) String() (val string) {
	if p.IntVal != nil {
		val = strconv.FormatInt(*p.IntVal, 10)
		return
	}

	if p.FloatVal != nil {
		val = strconv.FormatFloat(*p.FloatVal, 'f', -1, 64) //-1 will remove unimportant 0s
		return
	}

	val = string(p.Value)
	return
}

func (p *Point) Int() (val int64) {
	if p.IntVal != nil {
		val = (*(p.IntVal))
	} else if p.FloatVal != nil {
		val = int64(*p.FloatVal)
	}

	return
}

func (p *Point) Float() (val float64) {
	if p.FloatVal != nil {
		val = (*(p.FloatVal))
	} else if (p.IntVal) != nil {
		val = float64(*p.IntVal)
	}

	return
}

func (p *Point) MarshalJSON() ([]byte, error) {
	if p.IntVal != nil {
		return json.Marshal(p.IntVal)
	}

	if p.FloatVal != nil {
		return json.Marshal(p.FloatVal)
	}

	return json.Marshal(p.Value) //default raw message
}

func (p *Point) UnmarshalJSON(b []byte) (err error) {
	p.IntVal = nil
	p.FloatVal = nil

	var i int64
	if e := json.Unmarshal(b, &i); e == nil {
		p.IntVal = &i
		return
	}
	var f float64
	if e := json.Unmarshal(b, &f); e == nil {
		p.FloatVal = &f
		return
	}

	p.Value = json.RawMessage(b) //default to raw message
	return
}
