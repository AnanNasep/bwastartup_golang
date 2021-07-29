package campaign

import (
	"bwastartup/user"
	"time"
)


type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	BackerCount      string
	Slug             string
	CreatedAt        time.Time
	UpdatedAt		 time.Time
	CampaignImages   []CampaignImage
	//panggil user yang ada di package/foldler user... buat relasi
	User			 user.User
}

type CampaignImage struct {
	ID               int
	CampaignID       int
	FileName         string
	IsPrimary 		 int
	CreatedAt        time.Time
	UpdatedAt		 time.Time
}