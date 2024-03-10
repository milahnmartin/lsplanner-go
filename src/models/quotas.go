package models

import "time"

type Quota struct {
	ID            int
	current_count int
	max_count     int
	created_at    time.Time
	updated_at    time.Time
}
