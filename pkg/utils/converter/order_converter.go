package converter

import (
	"errors"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/jinzhu/copier"
)

func OrderModelToDTO(dataModel *model.Order) (*dto.Order, error) {
	dataDTO := &dto.Order{}
	err := copier.Copy(&dataDTO, &dataModel)
	if err != nil {
		return nil, errors.New("error converting product model to dto")
	}
	return dataDTO, nil
}

func OrderDTOToModel(dataDTO *dto.Order) (*model.Order, error) {
	dataModel := &model.Order{}
	err := copier.CopyWithOption(&dataModel, &dataDTO, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.New("error converting product dto to model")
	}
	return dataModel, nil
}
