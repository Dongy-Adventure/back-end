package converter

import (
	"errors"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/jinzhu/copier"
)

func ReviewModelToDTO(dataModel *model.Review) (*dto.Review, error) {
	dataDTO := &dto.Review{}
	err := copier.Copy(&dataDTO, &dataModel)
	if err != nil {
		return nil, errors.New("error converting review model to dto")
	}
	return dataDTO, nil
}

func ReviewDTOToModel(dataDTO *dto.Review) (*dto.Review, error) {
	dataModel := &dto.Review{}
	err := copier.CopyWithOption(&dataModel, &dataDTO, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.New("error converting review dto to model")
	}
	return dataModel, nil
}
