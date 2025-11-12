package repositories

import (
	"context"
	"errors"
	"log"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	"ms-genexis-pos-operaciones/context/shift/domain/value_object/constants"
	infrastructure_db_client "ms-genexis-pos-operaciones/infrastructure/db/client"
)

type GetPersonShiftRepository struct {
	Connection infrastructure_db_client.DatabaseConnectionInterface
}

func (g *GetPersonShiftRepository) GetPersonShit(infoShift *entities.OpeningShiftRequest) (*entities.PersonShift, error) {

	conn, err := g.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	defer conn.PgxConn.Release()
	log.Println("CONSULTANDO: ", constants.QUERY_GET_PERSON_SHIFT)
	log.Println("ARGUMENTO 1 : ", string(infoShift.Usuario))
	log.Println("ARGUMENTO 2 : ", string(infoShift.Clave))

	response := &entities.PersonShift{}
	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_PERSON_SHIFT,
		infoShift.Usuario,
		infoShift.Clave,
	).Scan(
		&response.Id,
		&response.Identificacion,
		&response.Pin,
		&response.Nombres,
		&response.Apellidos,
		&response.PerfilesId,
	)
	if err != nil {
		log.Println("Error [GetPersonShiftRepository] - ", err)
		return nil, errors.New("no se encontro usuario a identificar")
	}

	return response, nil

}
