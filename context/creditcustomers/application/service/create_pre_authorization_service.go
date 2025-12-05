package service

import (
	"context"
	"log"

	"ms-genexis-pos-operaciones/context/creditcustomers/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/creditcustomers/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CreatePreAuthorizationService struct {
	UseCase iusecase.ICreatePreAuthorization
}

func (s *CreatePreAuthorizationService) Execute(ctx context.Context, req *entities.PreAuthorizationRequest) (*entities_main.Response[bool], error) {
	result, err := s.UseCase.Execute(ctx, req)
	if err != nil {
		log.Println("CreatePreAuthorizationService: ", err)
		return nil, err
	}
	return result, nil
}
