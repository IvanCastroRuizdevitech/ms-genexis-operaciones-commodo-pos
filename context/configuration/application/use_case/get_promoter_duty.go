package usecase

import (
	"log"
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/configuration/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPromoterDuty struct {
	GetPromoterDuty irepositories.IGetPromoterDutyRepository
}

func (c *GetPromoterDuty) Execute() (*entities_main.Response[[]entities.PromoterDuty], error) {
	result, err := c.GetPromoterDuty.GetPromoterDuty()
	if err != nil {
		log.Println("GetPromoterDuty: ", err)
		return nil, err
	}
	return result, nil
}
