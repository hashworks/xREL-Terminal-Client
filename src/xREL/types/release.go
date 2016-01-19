package types

import "time"

type Releases struct {
	TotalCount int        `json:"total_count"`
	Pagination Pagination `json:"pagination"`
	List       []Release  `json:"list"`
}

type Release struct {
	Id          string       `json:"id"`
	Dirname     string       `json:"dirname"`
	LinkHref    string       `json:"link_href"`
	TimeUnix    int64        `json:"time"`
	GroupName   string       `json:"group_name"`
	NukeReason  string       `json:"nuke_reason"`
	Size        Size         `json:"size"`
	VideoType   string       `json:"video_type"`
	AudioType   string       `json:"audio_type"`
	NumRatings  int          `json:"num_ratings"`
	VideoRating float32      `json:"video_rating"`
	AudioRating float32      `json:"audio_rating"`
	ExtInfo     ShortExtInfo `json:"ext_info"`
	TVSeason    int          `json:"tv_season"`
	TVEpisode   int          `json:"tv_episode"`
	Comments    int          `json:"comments"`
	Flags       Flags        `json:"flags"`
	ProofUrl    string       `json:"proof_url"`
}

func (release *Release) GetTime() time.Time {
	return time.Unix(release.TimeUnix, 0)
}

type Size struct {
	Number int    `json:"number"`
	Unit   string `json:"unit"`
}

// http://www.xrel.to/comments/wiki/1680.html#cpost311058
type Flags struct {
	HasReadNFO bool `json:"read_nfo"`
	IsFixRLS   bool `json:"fix_rls"`
	IsTopRLS   bool `json:"top_rls"`
	IsEnglish  bool `json:"english"`
}

type Category struct {
	Name      string `json:"name"`
	ParentCat string `json:"parent_cat"`
}

type Filter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
