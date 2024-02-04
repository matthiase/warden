package postgres

import (
	"database/sql"

	"github.com/birdbox/authnz/models"
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/ksuid"
)

type UserStore struct {
	sqlx.Ext
}

func (db *UserStore) Find(id string) (*models.User, error) {
	user := models.User{}
	err := sqlx.Get(db, &user, "SELECT * FROM users WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *UserStore) Create(firstName string, lastName string, email string) (*models.User, error) {
	sql := `
		INSERT INTO users (id, first_name, last_name, email)
		VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, email
	`

	uid := ksuid.New()
	result, err := db.Queryx(sql, uid.String(), firstName, lastName, email)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	result.Next()

	user := models.User{}
	if err := result.StructScan(&user); err != nil {
		return nil, err
	}

	return &user, nil

	//now := time.Now()

	//account := &models.User{
	//	FirstName: firstName,
	//	LastName:  lastName,
	//	Email:     email,
	//}

	//result, err := sqlx.NamedQuery(db,
	//	`INSERT INTO accounts (
	//		username,
	//		password,
	//		locked,
	//		require_new_password,
	//		password_changed_at,
	//		created_at,
	//		updated_at
	//	)
	//	VALUES (:username, :password, :locked, :require_new_password, :password_changed_at, :created_at, :updated_at)
	//	RETURNING id`,
	//	account,
	//)
	//if err != nil {
	//	return nil, err
	//}
	//defer result.Close()
	//result.Next()
	//var id int64
	//err = result.Scan(&id)
	//if err != nil {
	//	return nil, err
	//}
	//account.ID = int(id)

	//return account, nil
}
