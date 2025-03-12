package converter

import (
	"errors"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/jinzhu/copier"
)

func AppointmentModelToDTO(dataModel *model.Appointment) (*dto.Appointment, error) {
	dataDTO := &dto.Appointment{}
	err := copier.Copy(&dataDTO, &dataModel)
	if err != nil {
		return nil, errors.New("error converting appointment model to dto")
	}
	return dataDTO, nil
}

func AppointmentDTOToModel(dataDTO *dto.Appointment) (*dto.Appointment, error) {
	dataModel := &dto.Appointment{}
	err := copier.CopyWithOption(&dataModel, &dataDTO, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.New("error converting appointment dto to model")
	}
	return dataModel, nil
}
