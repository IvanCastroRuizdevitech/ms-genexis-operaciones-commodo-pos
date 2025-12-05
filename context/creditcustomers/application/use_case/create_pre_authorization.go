package use_case

import (
	"context"
	"log"
	"time"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CreatePreAuthorization struct {
	Repository irepositories.IPreAuthorizationRepository
}

func (uc *CreatePreAuthorization) Execute(ctx context.Context, req *entities.PreAuthorizationRequest) (*entities_main.Response[bool], error) {
	result, err := uc.Repository.Create(ctx, req)
	if err != nil {
		log.Println("CreatePreAuthorization: ", err)
		return nil, err
	}

	success := entities_main.NewSuccessResponse(
		200,
		"Pre autorización creada",
		time.Now().Format("2006-01-02 15:04:05"),
		&result,
	)

	return &success, nil
}
