package core

// This file declares Functions and Interfaces to used inside
// the use cases instead of using functions like `time.Now` directly.
// This way we can be more predicable in our domain tests.

import (
	"time"

	"github.com/google/uuid"
)

type Now func() time.Time

type UUID func() uuid.UUID
