package repositories

import (
	"context"
	"encoding/json"
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

func (g *GetPersonShiftRepository) ValidatePerson(info *entities.PersonValidationRequest, requireAdmin bool) (*entities.PersonValidationResult, error) {
	return g.validatePersonAdmin(info.Usuario, info.Clave, info.Tag, requireAdmin)
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

func (g *GetPersonShiftRepository) validatePersonAdmin(usuario string, clave string, tag string, requireAdmin bool) (*entities.PersonValidationResult, error) {
	conn, err := g.Connection.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer conn.PgxConn.Release()

	log.Println("CONSULTANDO: ", constants.QUERY_VALIDATE_PERSON_SHIFT_ADMIN)
	log.Println("ARGUMENTO 1 : ", usuario)
	log.Println("ARGUMENTO 2 : ", clave)
	log.Println("ARGUMENTO 3 : ", tag)
	log.Println("ARGUMENTO 4 : ", requireAdmin)

	rawJSON := "{}"
	if err := conn.PgxConn.QueryRow(
		context.Background(),
		constants.QUERY_VALIDATE_PERSON_SHIFT_ADMIN,
		usuario,
		clave,
		tag,
		requireAdmin,
	).Scan(&rawJSON); err != nil {
		log.Println("Error [GetPersonShiftRepository] - ", err)
		return nil, errors.New("no se encontro usuario a identificar")
	}

	result := &entities.PersonValidationResult{}
	if err := json.Unmarshal([]byte(rawJSON), result); err != nil {
		return nil, err
	}

	return result, nil
}
