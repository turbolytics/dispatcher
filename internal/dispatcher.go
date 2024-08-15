package internal

import (
	"github.com/google/uuid"
	"net/url"
)

// Service auth required
type Organization struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Service Auth required
type Application struct {
	ID       uuid.UUID `json:"id"`
	OrgID    uuid.UUID `json:"org_id"`
	APIToken string    `json:"api_token"`
}

// App Auth Required
type Subscription struct {
	ID            uuid.UUID `json:"id"`
	AppID         uuid.UUID `json:"app_id"`
	Target        url.URL   `json:"target"`
	SigningSecret string    `json:"signing_secret"`
}
