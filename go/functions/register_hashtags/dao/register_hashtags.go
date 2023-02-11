package dao

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-batch/functions/config"
	"github.com/tmp-friends/victo-batch/functions/dto"
	"github.com/tmp-friends/victo-batch/functions/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type RegisterHashtagsDao struct {
	DB *sql.DB
}

func NewRegisterHashtagsDao() *RegisterHashtagsDao {
	db := config.NewMySQLConnector()

	return &RegisterHashtagsDao{
		DB: db.Conn,
	}
}

func (rhd *RegisterHashtagsDao) GetVtubersByNames(names []string) models.VtuberSlice {
	vtuberSlice, err := models.Vtubers(
		qm.Select("id", "name"),
		models.VtuberWhere.Name.IN(names),
	).All(context.Background(), rhd.DB)

	if err != nil {
		panic(err)
	}

	return vtuberSlice
}

func (rhd *RegisterHashtagsDao) IsExistsHashtags() bool {
	is_exists, err := models.Hashtags().Exists(context.Background(), rhd.DB)
	if err != nil {
		panic(err)
	}

	return is_exists
}

func (rhd *RegisterHashtagsDao) RegisterHashtags(hashtags []*dto.Hashtag, is_exists bool) {
	for i, v := range hashtags {
		hashtag := models.Hashtag{
			// DBのレコードは1から
			ID:       i + 1,
			Name:     v.Name,
			VtuberID: v.VtuberId,
		}

		err := hashtag.Upsert(context.Background(), rhd.DB, boil.Infer(), boil.Infer())
		if err != nil {
			panic(err)
		}
	}
}
