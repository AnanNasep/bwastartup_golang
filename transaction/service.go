package transaction

import (
	"bwastartup/campaign"
	"errors"
)

//SERVICE UNTUK LOGIC

type service struct {
	repository Repository
	campaignRepository campaign.Repository
}

type Service interface{
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int)([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service{
	return &service{repository, campaignRepository}
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
