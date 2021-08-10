package transaction

import (
	"bwastartup/campaign"
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
	PaymentURL	string
	User		user.User
	Campaign	campaign.Campaign
	CreatedAt	time.Time
	UpdatedAt	time.Time
}