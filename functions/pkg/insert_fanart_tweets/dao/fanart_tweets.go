package dao

import (
	"context"
	"database/sql"

	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/tmp-friends/victo-batch/functions/models"
	"github.com/tmp-friends/victo-batch/functions/pkg/config"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type FanartTweetsDao struct {
	DB *sql.DB
}

func NewFanartTweetsDao() *FanartTweetsDao {
	db := config.NewMySQLConnector()

	return &FanartTweetsDao{
		DB: db.Conn,
	}
}

func (ftd *FanartTweetsDao) GetHashtags() models.HashtagSlice {
	hashtags, err := models.Hashtags(qm.Select("id", "name")).All(context.Background(), ftd.DB)
	if err != nil {
		panic(err)
	}

	return hashtags
}

func (ftd *FanartTweetsDao) InsertTweetObject(tweet *twitterscraper.Tweet, hashtagId int) {
	to := models.TweetObject{
		ID:           string(tweet.ID),
		Text:         null.StringFrom(tweet.Text),
		RetweetCount: int(tweet.Retweets),
		LikeCount:    int(tweet.Likes),
		AuthorID:     tweet.UserID,
		URL:          tweet.PermanentURL,
		TweetedAt:    tweet.TimeParsed,
		HashtagID:    hashtagId,
	}

	err := to.Upsert(context.Background(), ftd.DB, boil.Blacklist("created_at"), boil.Infer())
	if err != nil {
		panic(err)
	}
}

func (ftd *FanartTweetsDao) InsertMediaObject(tweet *twitterscraper.Tweet) {
	var mediaType string
	switch true {
	case len(tweet.Videos) > 1:
		mediaType = "video"
	default:
		mediaType = "photo"
	}

	for _, photoUrl := range tweet.Photos {
		mo := models.MediaObject{
			Type:    mediaType,
			URL:     photoUrl,
			TweetID: tweet.ID,
		}

		err := mo.Insert(context.Background(), ftd.DB, boil.Infer())
		if err != nil {
			panic(err)
		}
	}
}

func (ftd *FanartTweetsDao) InsertAuthor(author *twitterscraper.Profile) {
	to := models.Author{
		ID:              author.UserID,
		Name:            author.Name,
		Username:        author.Username,
		ProfileImageURL: null.StringFrom(author.Avatar),
		BannerImageURL:  null.StringFrom(author.Banner),
		Biography:       null.StringFrom(author.Biography),
		Website:         null.StringFrom(author.Website),
		FollowersCount:  author.FollowersCount,
		FollowingCount:  author.FollowingCount,
		TweetsCount:     author.TweetsCount,
		Joined:          *author.Joined,
	}

	err := to.Upsert(context.Background(), ftd.DB, boil.Blacklist("created_at"), boil.Infer())
	if err != nil {
		panic(err)
	}
}
