package converter

import (
	"errors"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/jinzhu/copier"
)

func ProductModelToDTO(dataModel *model.Product) (*dto.Product, error) {
	dataDTO := &dto.Product{}
	err := copier.Copy(&dataDTO, &dataModel)
	if err != nil {
		return nil, errors.New("error converting product model to dto")
	}
	return dataDTO, nil
}

func ProductDTOToModel(dataDTO *dto.Product) (*dto.Product, error) {
	dataModel := &dto.Product{}
	err := copier.CopyWithOption(&dataModel, &dataDTO, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.New("error converting product dto to model")
	}
	return dataModel, nil
}
