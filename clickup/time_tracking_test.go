package clickup

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func TestTimeTrackingService_CreateTimeTracking(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	teamID := "teamID"
	taskID := "taskID"
	assigneeID := 1234
	var dur int32 = 120000

	start := time.Now().Add(time.Millisecond * time.Duration(-dur))

	_, _, err := client.TimeTrackings.CreateTimeTracking(context.Background(), teamID,
		&CreateTimeTrackingOptions{},
		&TimeTrackingRequest{
			Description: "description",
			Start:       start.Unix() * 1000,
			Duration:    dur,
			Assignee:    assigneeID,
			Tid:         taskID,
			Billable:    true,
		})

	OK(t, err)
}

func OK(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}
