package transaction

import "time"

// UNTUK MEWAKILI TABEL YG ADA DARI DB

type transaction struct {
	ID			int
	CampaignID	int
	userID		int
	Amount		int
	Status		string
	Code		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}