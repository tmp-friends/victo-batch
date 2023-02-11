package struct

type Vtuber struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	BelongsTo       string `json:"belongs_to"`
	ProfileImageURL string `json:"profile_image_url"`
	TwitterUserName string `json:"twitter_user_name"`
	Channel         string `json:"channel"`
}
