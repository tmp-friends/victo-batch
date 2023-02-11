package dao

import (
	"context"
	"database/sql"

	"github.com/tmp-friends/victo-batch/functions/config"
	"github.com/tmp-friends/victo-batch/functions/dto"
	"github.com/tmp-friends/victo-batch/functions/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type RegisterVtubersDao struct {
	DB *sql.DB
}

func NewRegisterVtubersDao() *RegisterVtubersDao {
	db := config.NewMySQLConnector()

	return &RegisterVtubersDao{
		DB: db.Conn,
	}
}

func (rvd *RegisterVtubersDao) RegisterVtubers(vtubers []*dto.Vtuber) {
	for i, v := range vtubers {
		vtuber := models.Vtuber{
			// DBのレコードは1から
			ID:              i + 1,
			Name:            v.Name,
			BelongsTo:       null.StringFrom(v.BelongsTo),
			ProfileImageURL: null.StringFrom(v.ProfileImageURL),
			TwitterUserName: null.StringFrom(v.TwitterUserName),
			Channel:         null.StringFrom(v.Channel),
		}

		err := vtuber.Upsert(context.Background(), rvd.DB, boil.Infer(), boil.Infer())
		if err != nil {
			panic(err)
		}
	}
}
