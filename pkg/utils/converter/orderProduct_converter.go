package converter

import (
	"errors"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/jinzhu/copier"
)

func OrderProductModelToDTO(dataModel *model.OrderProduct) (*dto.OrderProduct, error) {
	dataDTO := &dto.OrderProduct{}
	err := copier.Copy(&dataDTO, &dataModel)
	if err != nil {
		return nil, errors.New("error converting order product model to dto")
	}
	return dataDTO, nil
}

func OrderProductDTOToModel(dataDTO *dto.OrderProduct) (*model.OrderProduct, error) {
	dataModel := &model.OrderProduct{}
	err := copier.CopyWithOption(&dataModel, &dataDTO, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.New("error converting order product dto to model")
	}
	return dataModel, nil
}
