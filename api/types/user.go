package types

import (
	"time"
	"strconv"
)

type User struct {
	Id				string	`json:"id"`
	Name			string	`json:"name"`
	Secret			string	`json:"secret"`
	Locale			string	`json:"locale"`
	AvatarURL		string	`json:"avatar_url"`
	AvatarThumbURL	string	`json:"avatar_thumb_url"`
}

type RateLimitStatus struct {
	RemainingCalls	int			`json:"remaining_calls"`
	ResetTimeU		int			`json:"reset_time_u"`
}

func (rateLimitStatus *RateLimitStatus) GetResetTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, strconv.Itoa(rateLimitStatus.ResetTimeU))
}