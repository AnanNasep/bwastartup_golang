package transaction

import "gorm.io/gorm"

//REPOSITORY UNTUK AKES KE DATABASE

type repository struct{
	db *gorm.DB
}

type Repository interface{
	GetByCampaignID(CampaignID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
	//save transaksi midtrans
	Save(transaction Transaction)(Transaction, error)
	//update untuk midtrans
	Update(transaction Transaction)(Transaction, error)
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) GetByCampaignID(CampaignID int) ([]Transaction, error){
	var transactions []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", CampaignID).Order("id desc").Find(&transactions).Error
	if err != nil{
		return transactions, err
	}
	return transactions, nil
}

//buat list transaction
func (r *repository) GetByUserID(userID int) ([]Transaction, error){
	var transactions []Transaction
	//cara menapatkan data yang tidak berelasi secara langsung
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil{
		return transactions, err
	}
	return transactions, nil
}

//save transaksi midtrans
func (s *repository) Save(transaction Transaction)(Transaction, error){
	err := s.db.Create(&transaction).Error

	if err != nil{
		return transaction, err
	}
	return transaction, nil
}

//update untuk midtrans
func (r *repository) Update(transaction Transaction)(Transaction, error){
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}	