package transaction

import (
	"kami-peduli/campaign"
	"kami-peduli/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreateAt   time.Time
	UpdatedAt  time.Time
}
