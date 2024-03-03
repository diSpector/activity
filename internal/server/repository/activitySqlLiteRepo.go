package repository

import (
	"database/sql"

	"github.com/diSpector/activity.git/pkg/activity/entities"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type ActivitySqlLiteRepo struct {
	db *sql.DB
}

func NewActivitySqlLiteRepo(dbPath string) (*ActivitySqlLiteRepo, error) {
	db, err := sql.Open(`sqlite3`, dbPath)
	if err != nil {
		return nil, errors.Errorf(`err open db: %s`, err)
	}

	return &ActivitySqlLiteRepo{
		db: db,
	}, nil
}

func (s *ActivitySqlLiteRepo) Close() error {
	return s.db.Close()
}

func (s *ActivitySqlLiteRepo) Insert(activity *entities.Activity) (int64, error) {
	sql := `insert into activities (description, persons) values (:description, :persons)`
	smnt, err := s.db.Prepare(sql)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer smnt.Close()

	res, err := smnt.Exec(activity.Activity, activity.Participants)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	if id == 0 {
		return 0, errors.New(`non-proper insert`)
	}

	return id, nil
}

func (s *ActivitySqlLiteRepo) SelectByName(name string) ([]*entities.Activity, error) {
	sql := `select id, description, persons from activities where description like concat('%', :desc, '%')`
	rows, err := s.db.Query(sql, name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	var acts []*entities.Activity
	for rows.Next() {
		var act entities.Activity
		err := rows.Scan(&act.Id, &act.Activity, &act.Participants)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		acts = append(acts, &act)
	}

	return acts, nil
}

func (s *ActivitySqlLiteRepo) SelectAll() ([]*entities.Activity, error) {
	sql := `select id, description, persons from activities`
	rows, err := s.db.Query(sql)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	var acts []*entities.Activity
	for rows.Next() {
		var act entities.Activity
		err := rows.Scan(&act.Id, &act.Activity, &act.Participants)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		acts = append(acts, &act)
	}

	return acts, nil
}
