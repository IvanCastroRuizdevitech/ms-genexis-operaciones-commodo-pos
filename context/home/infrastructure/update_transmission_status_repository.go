package repositories

import (
	"context"
	"fmt"

	"ms-genexis-pos-operaciones/context/home/domain/value_object/constants"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type UpdateTransmissionStatusRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (r *UpdateTransmissionStatusRepository) MarkSynchronized(id int) error {
	conn, err := r.Connection.GetDatabaseConnection()
	if err != nil {
		return err
	}
	defer conn.PgxConn.Release()

	cmdTag, err := conn.PgxConn.Exec(
		context.Background(),
		constants.QUERY_MARK_TRANSMISSION_AS_SYNCED,
		id,
	)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("no transmission updated for id %d", id)
	}

	return nil
}
