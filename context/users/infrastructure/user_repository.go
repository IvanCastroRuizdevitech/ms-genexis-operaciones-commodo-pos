package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5"

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

func (r *UserRepository) AssignTag(request *entities.AssignTagRequest) (*entities_main.Response[entities.AssignTagResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	ctx := context.Background()
	tx, err := conn.PgxConn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, constants.QueryClearTag, request.Tag); err != nil {
		return nil, err
	}

	var id int64
	if err := tx.QueryRow(ctx, constants.QueryAssignTag, request.Tag, request.Identification).Scan(&id); err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	result := entities.AssignTagResult{ID: id}
	success := entities_main.NewSuccessResponse(
		200,
		"Tag assigned",
		time.Now().Format("2006-01-02 15:04:05"),
		&result,
	)

	return &success, nil
}

func (r *UserRepository) GenerateAssignTagTransmissions(request *entities.AssignTagTransmissionsRequest) (*entities_main.Response[entities.AssignTagTransmissionsResult], error) {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	var transmissionsRaw string
	if err := conn.PgxConn.QueryRow(context.Background(), constants.QueryGenerateAssignTagTransmissions, request.Identification, request.Tag, request.Medio).Scan(&transmissionsRaw); err != nil {
		return nil, err
	}

	result := entities.AssignTagTransmissionsResult{TransmissionsRaw: transmissionsRaw}
	success := entities_main.NewSuccessResponse(
		200,
		"Tag transmissions generated",
		time.Now().Format("2006-01-02 15:04:05"),
		&result,
	)

	return &success, nil
}
