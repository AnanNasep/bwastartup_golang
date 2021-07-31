package campaign

import "gorm.io/gorm"

type Repository interface {
	//campaign
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	//campaignDetail
	FindByID(ID int)(Campaign, error)
	Save(campaign Campaign)(Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}
// campaign find all
func (r *repository)FindAll() ([]Campaign, error){
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
//ambil campaign find by user id
func (r *repository)FindByUserID(userID int) ([]Campaign, error){
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

//ambil campaign detail
func (r *repository) FindByID(ID int) (Campaign, error){
	var campaign Campaign

	//ambil user yang bikin campaign beserta foto campaign nya
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error

	if err != nil{
		return	campaign, err
	}
	return campaign, nil

}

//create campaign
func (r *repository) Save(campaign Campaign)(Campaign, error){
	err := r.db.Create(&campaign).Error
	if err != nil{
		return campaign, err
	}
	return	campaign, nil
}