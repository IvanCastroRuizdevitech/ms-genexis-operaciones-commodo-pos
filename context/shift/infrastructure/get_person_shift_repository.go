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
	return g.fetchPerson(infoShift.Usuario, infoShift.Clave, "")
}

func (g *GetPersonShiftRepository) ValidatePerson(info *entities.PersonValidationRequest) (*entities.PersonShift, error) {
	return g.fetchPerson(info.Usuario, info.Clave, info.Tag)
}

func (g *GetPersonShiftRepository) fetchPerson(usuario string, clave string, tag string) (*entities.PersonShift, error) {

	conn, err := g.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	defer conn.PgxConn.Release()
	log.Println("CONSULTANDO: ", constants.QUERY_GET_PERSON_SHIFT)
	log.Println("ARGUMENTO 1 : ", usuario)
	log.Println("ARGUMENTO 2 : ", clave)
	log.Println("ARGUMENTO 3 : ", tag)

	response := &entities.PersonShift{}
	err = conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_GET_PERSON_SHIFT,
		usuario,
		clave,
		tag,
	).Scan(
		&response.Id,
		&response.Identificacion,
		&response.Pin,
		&response.Nombres,
		&response.Apellidos,
		&response.PerfilesId,
		&response.JornadasId,
	)
	if err != nil {
		log.Println("Error [GetPersonShiftRepository] - ", err)
		return nil, errors.New("no se encontro usuario a identificar")
	}

	return response, nil

}
