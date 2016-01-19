package types

const (
	TYPE_MOVIE    = "movie"
	TYPE_TV       = "tv"
	TYPE_GAME     = "game"
	TYPE_CONSOLE  = "console"
	TYPE_SOFTWARE = "software"
	TYPE_XXX      = "xxx"
)

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalPages  int `json:"total_pages"`
}
