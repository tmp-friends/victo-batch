package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/sivchari/gotwtr"
	"github.com/tmp-friends/victo-batch/functions/config"
	"github.com/tmp-friends/victo-batch/functions/models"
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

func (ftd *FanartTweetsDao) InsertTweetObject(tweet *gotwtr.Tweet, url string, hashtagId int) {
	to := models.TweetObject{
		TweetID:      string(tweet.ID),
		Text:         null.StringFrom(tweet.Text),
		RetweetCount: int(tweet.PublicMetrics.RetweetCount),
		LikeCount:    int(tweet.PublicMetrics.LikeCount),
		AuthorID:     tweet.AuthorID,
		URL:          url,
		TweetedAt:    ftd.strToTime(tweet.CreatedAt),
		HashtagID:    hashtagId,
	}

	err := to.Insert(context.Background(), ftd.DB, boil.Infer())
	if err != nil {
		panic(err)
	}
}

func (ftd *FanartTweetsDao) InsertMediaObject(media *gotwtr.Media, tweetId string) {
	mo := models.MediaObject{
		MediaKey: media.MediaKey,
		Type:     media.Type,
		URL:      media.URL,
		TweetID:  tweetId,
	}

	err := mo.Upsert(context.Background(), ftd.DB, boil.Blacklist("created_at"), boil.Infer())
	if err != nil {
		panic(err)
	}
}

func (ftd *FanartTweetsDao) strToTime(str string) time.Time {
	// RFC3339 ex."2006-01-02T15:04:05Z07:00"
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic(err)
	}

	return t
}
