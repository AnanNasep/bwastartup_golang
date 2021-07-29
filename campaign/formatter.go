package campaign

import (
	"strings"
)

// formatter itu sebagai pengatur apa yang ingin kita tampilkan datanya di JSON API
type CampaignFormatter struct {
	ID               int    `json:"id`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug		     string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""
	//cek klo ada fotonya
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter :=  []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

//formater campaign detail API json
type CampaignDetailFormatter struct{
	ID					int						`json:"id"`
	Name				string					`json:"name"`
	ShortDescription	string					`json:"short_description"`
	Description			string					`json:"description"`
	ImageURL			string					`json:"image_url"`
	GoalAmount			int						`json:"goal_amount"`
	CurrentAmount		int						`json:"current_amopunt"`
	UserID				int						`json:"user_id"`
	Slug				string					`json:"slug"`
	Perks				[]string				`json:"perks"`
	User				CampaignUserFormatter	`json:"user"` 	//ini definisi dari type yang bawah (CampaignUserFormatter)
	Images				[]CampaignImageFormatter`json:"images"` //ini definisi dari type yang bawah (CampaignImageFormatter)
}

type CampaignUserFormatter struct{
	Name		string		`json:"name"`
	ImageURL	string		`json:"image_url"`
}

type CampaignImageFormatter struct{
	ImageURL		string		`json:"image_url"`
	IsPrimary		bool		`json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter{
	campaignDetailFormatter	:= CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.ImageURL = ""
	//cek klo ada fotonya
	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	//ini buat pemecah si perks berdasarkan koma
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ","){
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perks = perks

	//manggil si json user
	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName
	campaignDetailFormatter.User = campaignUserFormatter // return

	//manggil si images
	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1{
			isPrimary = true
		} 
		campaignImageFormatter.IsPrimary = isPrimary
		images = append(images, campaignImageFormatter)
	}
	
	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}