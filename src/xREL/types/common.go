package types

const TYPE_MOVIE	= "movie"
const TYPE_TV		= "tv"
const TYPE_GAME		= "game"
const TYPE_CONSOLE	= "console"
const TYPE_SOFTWARE	= "software"
const TYPE_XXX		= "xxx"

type Pagination struct {
	CurrentPage int       `json:"current_page"`
	PerPage     int       `json:"per_page"`
	TotalPages  int       `json:"total_pages"`
}