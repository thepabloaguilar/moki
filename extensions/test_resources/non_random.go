package test_resources

import (
	"time"

	"github.com/google/uuid"

	"github.com/thepabloaguilar/moki/core"
)

func Now() core.Now {
	timeNow := time.Now().UTC().Truncate(time.Second)

	return func() time.Time { return timeNow }
}

func UUID() core.UUID {
	toReturn := uuid.New()

	return func() uuid.UUID { return toReturn }
}
