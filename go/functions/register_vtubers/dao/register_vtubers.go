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
	for _, v := range vtubers {
		vtuber := models.Vtuber{
			Name:            v.Name,
			BelongsTo:       null.StringFrom(v.BelongsTo),
			ProfileImageURL: null.StringFrom(v.ProfileImageURL),
			TwitterUserName: null.StringFrom(v.TwitterUserName),
			Channel:         null.StringFrom(v.Channel),
		}

		vtuber.Insert(context.Background(), rvd.DB, boil.Infer())
	}
}
