package models

import "time"

type User struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	APIToken string `db:"api_token"`
}

type URLStatus string

const (
	StatusQueued     URLStatus = "queued"
	StatusProcessing URLStatus = "processing"
	StatusDone       URLStatus = "done"
	StatusError      URLStatus = "error"
)

type URLAnalysis struct {
	ID            int       `db:"id"`
	UserID        int       `db:"user_id"`
	URL           string    `db:"url"`
	Status        URLStatus `db:"status"`
	HTMLVersion   string    `db:"html_version"`
	Title         string    `db:"title"`
	H1            int       `db:"h1"`
	H2            int       `db:"h2"`
	H3            int       `db:"h3"`
	H4            int       `db:"h4"`
	H5            int       `db:"h5"`
	H6            int       `db:"h6"`
	InternalLinks int       `db:"internal_links"`
	ExternalLinks int       `db:"external_links"`
	BrokenLinks   int       `db:"broken_links"`
	HasLoginForm  bool      `db:"has_login_form"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
