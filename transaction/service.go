package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"errors"
	"strconv"
)

//SERVICE UNTUK LOGIC

type service struct {
	repository Repository
	campaignRepository campaign.Repository
	//midtrans
	paymentService	payment.Service
}

type Service interface{
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int)([]Transaction, error)
	//save transaksi midtrans
	CreateTransaction(input CreateTransactionInput)(Transaction, error)
	//ambil data notifikasi dari midtrans
	ProcessPayment(input TransactionNotificationInput) error
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service{
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error){
	// CARA AGAR YANG MELIHAT DATA TRANSAKSI ADALAH USER YANG LOGIN AJA
	// TAMPAH campaignRepository
	//get campaign
	//cek campaign.userid != user_id_yang_melakukan_request
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil{
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID{
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil{
		return transactions, err
	}
	return transactions, nil
}

//get list transaction
func (s *service) GetTransactionsByUserID(userID int)([]Transaction, error){
	transactions, err := s.repository.GetByUserID(userID)
	if err !=nil{
		return transactions, err
	}
	return transactions, nil
}

//save transaksi midtrans
func(s *service) CreateTransaction(input CreateTransactionInput)(Transaction, error){
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"
	//transaction.Code = ""

	newTransaction, err := s.repository.Save(transaction)
	if err !=nil{
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID: newTransaction.ID,
		Amount: newTransaction.Amount,
	}
	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err !=nil{
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL
	newTransaction, err = s.repository.Update(transaction)
	if err !=nil{
		return newTransaction, err
	}

	return newTransaction, nil
}

//ambil data notifikasi dari midtrans
func (s *service) ProcessPayment(input TransactionNotificationInput) error{
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByID(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement"{
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expired" || input.TransactionStatus == "cancel"{
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err:= s.campaignRepository.FindByID(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err = s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}
