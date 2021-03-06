package campaign

import "gorm.io/gorm"

type Repository interface {
	//get campaign
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	//campaignDetail
	FindByID(ID int)(Campaign, error)
	//create campaign
	Save(campaign Campaign)(Campaign, error)
	//update dampaign
	Update(campaign Campaign)(Campaign, error)
	//upload campaign image
	CreateImage(campaignImage CampaignImage)(CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int)(bool, error)
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

//update campaign
func(r *repository) Update(campaign Campaign) (Campaign, error){
	err := r.db.Save(&campaign).Error
	if err != nil{
		return campaign, err
	}
	return	campaign, nil
}

//Upload campaign image
func (r *repository) CreateImage(campaignImage CampaignImage)(CampaignImage, error){
	err := r.db.Create(&campaignImage).Error
	if err != nil{
		return campaignImage, err
	}
	return	campaignImage, nil
}
func (r *repository) MarkAllImagesAsNonPrimary(campaignID int)(bool, error){
	//update campaign_images SET is_primary = false WHERE campaign_id = 1

	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil{
		return false, err
	}
	return true, nil
}