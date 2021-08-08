package transaction

import "time"

// UNTUK MELAKUKAN FORMATING JSON

type CampaignTransactionFormatter struct{
	ID			int			`json:"id"`
	Name		string		`json:"name"`
	Amount		int			`json:"amount"`
	CreatedAt	time.Time	`json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter{
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}
//kalo banyak campaign
func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter{
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter
	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter
}

//buat list transaction
type UserTransactionFormatter struct{
	ID					int					`JSON:"id"`
	Amount				int					`JSON:"amount"`
	Status				string				`JSON:"status"`
	CreatedAt			time.Time			`JSON:"created_at"`
	Campaign			CampaignFormatter	`JSON:"campaign"`
}
type CampaignFormatter struct{
	Name			string			`JSON:"name"`
	ImageURL		string			`JSON:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter{
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	//buat image url
	campaignFormatter.ImageURL = ""
	if len(transaction.Campaign.CampaignImages) > 0{
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}
	formatter.Campaign = campaignFormatter

	return formatter
}
//kalo banyak transaction
func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter{
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter
	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter
}