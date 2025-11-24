package repositories

import (
	"context"
	"database/sql"
	"time"

	"ms-genexis-pos-operaciones/context/users/domain/entities"
	"ms-genexis-pos-operaciones/context/users/domain/value_object/constants"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type UserRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *UserRepository) GetAll() (*entities_main.Response[[]entities.User], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	rows, err := conn.PgxConn.Query(context.Background(), constants.QueryGetUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]entities.User, 0, 16)
	for rows.Next() {
		var (
			id             int64
			name           string
			identification string
			status         string
			phone          sql.NullString
			address        sql.NullString
			profileID      sql.NullInt64
			tag            sql.NullString
		)
		if err := rows.Scan(
			&id,
			&name,
			&identification,
			&status,
			&phone,
			&address,
			&profileID,
			&tag,
		); err != nil {
			return nil, err
		}
		users = append(users, entities.User{
			ID:             id,
			Name:           name,
			Identification: identification,
			Status:         status,
			Phone:          nullStringOrEmpty(phone),
			Address:        nullStringOrEmpty(address),
			ProfileID: func(v sql.NullInt64) int64 {
				if v.Valid {
					return v.Int64
				}
				return 0
			}(profileID),
			Tag: nullStringOrEmpty(tag),
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	data := users
	success := entities_main.NewSuccessResponse[[]entities.User](
		200,
		"OK",
		time.Now().Format("2006-01-02 15:04:05"),
		&data,
	)

	return &success, nil
}

func nullStringOrEmpty(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
