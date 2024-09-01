// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"database/sql"
	"time"
)

type Category struct {
	ID           string
	CategoryName string
	CreatedAt    sql.NullTime
}

type Event struct {
	ID             string
	OrganizationID string
	Title          string
	Img            sql.NullString
	Description    sql.NullString
	Registrant     sql.NullInt32
	MaxRegistrant  sql.NullInt32
	Date           time.Time
	StartTime      sql.NullTime
	EndTime        sql.NullTime
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}

type EventPricing struct {
	ID        int64
	EventID   string
	EventType sql.NullString
	Price     sql.NullInt32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type Organization struct {
	ID          string
	Name        string
	Description string
	HeaderImg   sql.NullString
	ProfileImg  sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

type Speaker struct {
	ID          string
	Name        string
	Title       sql.NullString
	Img         sql.NullString
	Description sql.NullString
	EventID     sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

type User struct {
	ID          string
	Email       string
	Password    sql.NullString
	OauthID     sql.NullString
	Name        string
	Role        string
	Phone       sql.NullString
	Img         sql.NullString
	IsPremium   sql.NullBool
	Lvl         sql.NullInt32
	Dob         sql.NullTime
	Institution sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	Username    sql.NullString
	Address     sql.NullString
}

type UserCategory struct {
	ID         int32
	CategoryID string
	UserID     string
	CreatedAt  sql.NullTime
}
