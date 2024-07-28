package domain

import (
	"time"

	"github.com/google/uuid"
)

type ID string
type Email string
type Password string
type DocumentRegistry string
type Enabled bool
type CreatedAt time.Time
type UpdatedAt time.Time
type DeletedAt *time.Time
type Version int64

func NewID() ID {
	return ID(uuid.New().String())
}

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}

func NewVersion() Version {
	return Version(0)
}

func (c CreatedAt) String() string {
	return time.Time(c).Format(time.RFC3339Nano)
}

func (u UpdatedAt) String() string {
	return time.Time(u).Format(time.RFC3339Nano)
}

func (c CreatedAt) Format() string {
	return time.Time(c).Format(time.RFC3339)
}

func (u UpdatedAt) Format() string {
	return time.Time(u).Format(time.RFC3339)
}

func (v Version) Int() int64 {
	return int64(v)
}

func (i ID) String() string {
	return string(i)
}

func (e Email) String() string {
	return string(e)
}

func (p Password) String() string {
	return string(p)
}

func (d DocumentRegistry) String() string {
	return string(d)
}

func (e Enabled) Bool() bool {
	return bool(e)
}
