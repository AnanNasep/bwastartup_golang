package transaction

import (
	"bwastartup/user"
	"time"
)

// UNTUK MEWAKILI TABEL YG ADA DARI DB

type Transaction struct {
	ID			int
	CampaignID	int
	UserID		int
	Amount		int
	Status		string
	Code		string
	User		user.User
	CreatedAt	time.Time
	UpdatedAt	time.Time
}