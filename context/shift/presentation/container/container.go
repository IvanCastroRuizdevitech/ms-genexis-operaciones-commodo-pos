package container_shift

import (
	"ms-genexis-pos-operaciones/context/shift/application/service"
	iservice "ms-genexis-pos-operaciones/context/shift/domain/ports/application/service"
)

//CLIENTS

// REPOSITORIES

// USECASE

// SERVICE
var openingShift iservice.IOpeningShift

func initializes() {

	openingShift = &service.OpeningShiftClient{}

}

func ResolveOpeningShiftContainer() iservice.IOpeningShift {

	if openingShift == nil {
		initializes()
	}
	return openingShift
}
