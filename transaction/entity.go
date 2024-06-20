package transaction

import (
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
	CreateAt   time.Time
	UpdatedAt  time.Time
}
