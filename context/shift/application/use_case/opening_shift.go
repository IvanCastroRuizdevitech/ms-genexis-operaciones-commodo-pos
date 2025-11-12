package usecase

import (
	"encoding/json"
	"log"
	"ms-genexis-pos-operaciones/context/shift/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/shift/domain/ports/repositories"
	"ms-genexis-pos-operaciones/context/shift/domain/value_object/constants"
)

type OpeningShift struct {
	SendOpening irepositories.ISendOpeningShiftRepositoryHttp
}

func (shift *OpeningShift) Execute(shift_info *entities.OpeningShiftRequest, person_id int) (*entities.OpeningShiftHttpResponse, error) {

	body := &entities.OpeningShiftHttp{
		PersonasId:      person_id,
		Surtidores:      shift_info.Surtidores,
		FechaInicio:     shift_info.FechaInicio,
		EquiposId:       shift_info.EquiposId,
		EmpresasId:      shift_info.EmpresasId,
		Atributos:       shift_info.Atributos,
		AjustePeriodico: shift_info.AjustePeriodico,
	}

	header := &map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	bodyByte, err := json.Marshal(body)
	if err != nil {
		log.Println("Error [OpeningShift] JsonMarshall - ", err)
		return nil, err
	}

	response_http, err := shift.SendOpening.Send(constants.API_OPENING_SHIFT, bodyByte, header)

	if err != nil {
		log.Println("SendOpeningShift: ", err)
		return nil, err
	}

	return response_http, nil
}
