package types

import (
	"time"
	"strconv"
)

type Comments struct {
	TotalCount string     `json:"total_count"`
	Pagination Pagination `json:"pagination"`
	List       []Comment  `json:"list"`
}

type Comment struct {
	Id       string        `json:"id"`
	TimeUnix string        `json:"time"`
	Author   Author        `json:"author"`
	Text     string        `json:"text"`
	LinkHref string        `json:"link_href"`
	Rating   Rating        `json:"rating"`
	Votes    Votes         `json:"votes"`
	Edits    Edits         `json:"edits"`
}

type Author struct {
	Id   string            `json:"id"`
	Name string            `json:"name"`
}

type Rating struct {
	Video string           `json:"video"`
	Audio string           `json:"audio"`
}

type Votes struct {
	Positive int           `json:"positive"`
	Negative int           `json:"negative"`
}

type Edits struct {
	Count    int           `json:"count"`
	LastUnix string        `json:"last"`
}

func (comment *Comment) GetTime() (time.Time, error) {
	var timeResult time.Time
	commentTime, err := strconv.ParseInt(comment.TimeUnix, 10, 64)
	if err == nil && commentTime != 0 {
		timeResult = time.Unix(commentTime, 0)
	}
	return timeResult, err
}

func (edits *Edits) GetLast() (time.Time, error) {
	var timeResult time.Time
	lastEdit, err := strconv.ParseInt(edits.LastUnix, 10, 64)
	if err == nil && lastEdit != 0 {
		timeResult = time.Unix(lastEdit, 0)
	}
	return timeResult, err
}