package app

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

func DoSomething(ctx context.Context, lastRunTime, thisRunTime time.Time) error {
	activity.GetLogger(ctx).Info("Cron job running.", "lastRunTime_exclude", lastRunTime, "thisRunTime_include", thisRunTime)

	// Query database, call external API, or do any non-deterministic behaviour.

	return nil
}
