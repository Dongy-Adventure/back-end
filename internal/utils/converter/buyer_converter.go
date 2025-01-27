package converter

import (
	"errors"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/jinzhu/copier"
)

func BuyerModelToDTO(dataModel *model.Buyer) (*dto.Buyer, error) {
	dataDTO := &dto.Buyer{}
	err := copier.Copy(&dataDTO, &dataModel)
	if err != nil {
		return nil, errors.New("error converting buyer model to dto")
	}
	return dataDTO, nil
}

func BuyerDTOToModel(dataDTO *dto.Buyer) (*dto.Buyer, error) {
	dataModel := &dto.Buyer{}
	err := copier.CopyWithOption(&dataModel, &dataDTO, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.New("error converting buyer dto to model")
	}
	return dataModel, nil
}
