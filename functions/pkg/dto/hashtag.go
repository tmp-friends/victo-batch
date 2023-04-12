package dto

type Hashtag struct {
	Name       string `json:"name"`
	VtuberName string `json:"vtuber_name"`
	VtuberId   int
}
