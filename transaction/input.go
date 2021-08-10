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

//ambil data notifikasi dari midtrans
type TransactionNotificationInput struct{
	TransactionStatus	string	`json:"transaction_status"`
	OrderID				string	`json:"order_id"`
	PaymentType			string	`json:"payment_type"`
	FraudStatus			string	`json:"fraud_status"`
}