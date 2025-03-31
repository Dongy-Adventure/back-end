package converter

import (
	"errors"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/jinzhu/copier"
)

func AdvertisementModelToDTO(dataModel *model.Advertisement) (*dto.Advertisement, error) {
	dataDTO := &dto.Advertisement{}
	err := copier.Copy(&dataDTO, &dataModel)
	if err != nil {
		return nil, errors.New("error converting advertisement model to dto")
	}
	return dataDTO, nil
}

func AdvertisementDTOToModel(dataDTO *dto.Advertisement) (*model.Advertisement, error) {
	dataModel := &model.Advertisement{}
	err := copier.CopyWithOption(&dataModel, &dataDTO, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.New("error converting review dto to model")
	}
	return dataModel, nil
}
