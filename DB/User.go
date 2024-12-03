package DB

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          int
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivateAt   sql.NullTime
	CreatedAt    time.Time
	UpdateAt     time.Time
	ignored      string
}
