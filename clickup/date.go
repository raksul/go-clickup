package clickup

import (
	"encoding/json"
	"strconv"
	"time"
)

type Date struct {
	unix json.Number
	time time.Time
	null bool
}

func NewDate(t time.Time) *Date {
	return &Date{
		unix: int64ToJsonNumber(t.UnixMilli()),
		time: t,
	}
}

func NewDateWithUnixTime(unix int64) *Date {
	return &Date{
		unix: int64ToJsonNumber(unix),
		time: time.UnixMilli(unix),
	}
}

func NullDate() *Date {
	return &Date{null: true}
}

func (d Date) Time() *time.Time {
	if d.null {
		return nil
	}

	return &d.time
}

func (d Date) String() string {
	if d.null {
		return ""
	}

	return d.time.String()
}

// Equal reports whether x and y are equal.
// This method was added to test with google/go-cmp.
// ref: https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal
func (x Date) Equal(y Date) bool {
	if x.null {
		return x.null == y.null
	}

	return x.time.Equal(y.time)
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err == nil {
		if str == "" {
			d.null = true

			return nil
		}
	}

	var v json.Number
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	n, err := jsonNumberToInt64(v)
	if err != nil {
		return err
	}

	d.unix = v
	d.time = time.UnixMilli(n)
	d.null = false

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.null {
		return json.Marshal(nil)
	}

	return json.Marshal(d.unix)
}

func int64ToJsonNumber(n int64) json.Number {
	b := []byte(strconv.Itoa(int(n)))

	var v json.Number
	if err := json.Unmarshal(b, &v); err != nil {
		panic(err.Error())
	}

	return v
}

func jsonNumberToInt64(num json.Number) (int64, error) {
	n, err := num.Int64()
	if err != nil {
		return 0, err
	}

	return n, nil
}
