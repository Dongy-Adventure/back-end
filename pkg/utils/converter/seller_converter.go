package converter

import (
	"errors"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/jinzhu/copier"
)

func SellerModelToDTO(dataModel *model.Seller) (*dto.Seller, error) {
	dataDTO := &dto.Seller{}
	err := copier.Copy(&dataDTO, &dataModel)
	if err != nil {
		return nil, errors.New("error converting seller model to dto")
	}
	return dataDTO, nil
}

func SellerDTOToModel(dataDTO *dto.Seller) (*model.Seller, error) {
	dataModel := &model.Seller{}
	err := copier.CopyWithOption(&dataModel, &dataDTO, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.New("error converting seller dto to model")
	}
	return dataModel, nil
}
