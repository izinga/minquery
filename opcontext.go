package minquery

import (
	"context"
	"time"
)

// OpTimeout bounds the find command and cursor iteration this package issues
// through the official MongoDB driver.
//
// Previously these used context.Background() and context.TODO(), so a
// paginated fetch had no deadline and could not be cancelled. Those are the
// queries still running when a replica set is already degraded.
//
// Set to zero to restore the previous unbounded behaviour.
var OpTimeout = 30 * time.Second

// opContext returns a context bounded by OpTimeout. The caller must always
// call the returned cancel func, and must keep the context alive for the whole
// cursor iteration - cancel only once the cursor is drained.
func opContext() (context.Context, context.CancelFunc) {
	if OpTimeout <= 0 {
		return context.Background(), func() {}
	}
	return context.WithTimeout(context.Background(), OpTimeout)
}
