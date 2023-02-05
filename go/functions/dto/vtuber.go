package dto

type Vtuber struct {
	Name            string `json:"name"`
	BelongsTo       string `json:"belongs_to"`
	ProfileImageURL string `json:"profile_image_url"`
	TwitterUserName string `json:"twitter_user_name"`
	Channel         string `json:"channel"`
	HashtagName     string `json:"hashtag_name"`
}
