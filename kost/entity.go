package kost

import (
	"backendEkost/user"
	"time"
)

type Kost struct {
	ID                int
	UserID            int
	Name              string
	ShortDescription  string
	Description       string
	Perks             string
	LiverCount        int
	SpaceCount        int
	CurrentSpaceCount int
	Slug              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	KostImages        []KostImage
	User              user.User
}

type KostImage struct {
	ID        int
	KostID    int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
