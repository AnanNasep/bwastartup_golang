package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)


type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	//update campaign
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput)(Campaign, error)
	//create campaign image
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string)(CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

//percabangan campaigns, antara ambil by userID atau ambil semua FindALL
func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

//create campaign
func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID
		
	slugCandidate:= fmt.Sprintf("%s, %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)		//Buat bikin nama campaign unik, namacampaign-ID => nama-campaign-10

	NewCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return NewCampaign, err
	}
	return NewCampaign, nil
}

//update campaign
func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput)(Campaign, error){
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil{
		return campaign, err
	}
	
	//PENGECEKAN USER
	if campaign.UserID != inputData.User.ID{
		//kalo yang updagte campaign bukan yang punya campaign
		return campaign, errors.New("Not an owner of the campaign")
	}
	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil{
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}

//create campaign image
func(s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string)(CampaignImage, error){
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil{
		return CampaignImage{}, err
	}
	//PENGECEKAN USER
	if campaign.UserID != input.User.ID{
		//kalo yang updagte campaign bukan yang punya campaign
		return CampaignImage{}, errors.New("Not an owner of the campaign")
	}
	//ini kondisi kalo is_primary nya true,, yang lain di ubah jadi is_primary nya false
	isPrimary := 0
	if input.IsPrimary{
		isPrimary = 1
		_, err:= s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil{
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil{
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}