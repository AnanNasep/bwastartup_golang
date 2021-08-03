package transaction

import "bwastartup/user"

// MENANGKAP INPUT DARI USER

type GetCampaignTransactionsInput struct {
	ID int `uri:"id" binding:"required"`
	User user.User
}
