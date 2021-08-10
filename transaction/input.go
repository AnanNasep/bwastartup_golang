package transaction

import "bwastartup/user"

// MENANGKAP INPUT DARI USER

type GetCampaignTransactionsInput struct {
	ID int `uri:"id" binding:"required"`
	User user.User
}

//save transaksi midtrans
type CreateTransactionInput struct{
	Amount				int		`json:"amount" binding:"required"`
	CampaignID			int		`json:"campaign_id" binding:"required"`
	User				user.User
}