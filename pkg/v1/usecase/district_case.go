package usecase

import (
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"gorm.io/gorm"
	"strconv"
)

type DistrictCase struct {
	repo interfaces.DistrictRepoInterface
}

func (districtCase *DistrictCase) Create(district models.District) (models.District, error) {
	if _, err := districtCase.repo.GetByCode(strconv.FormatInt(district.Code, 10)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.District{}, errors.New("the code has already been taken")
	}

	return districtCase.repo.Create(district)
}

func (districtCase *DistrictCase) Get(id int64) (models.District, error) {
	var district models.District
	var err error

	district, err = districtCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.District{}, errors.New("no such district with the id supplied")
		}

		return models.District{}, err
	}
	return district, nil
}

func (districtCase *DistrictCase) Update(updateDistrict models.District) (models.District, error) {
	var district models.District
	var err error

	district, err = districtCase.repo.Get(int64(updateDistrict.ID))
	if err != nil {
		return models.District{}, err
	}

	if district.Code != updateDistrict.Code {
		return models.District{}, errors.New("code cannot be changed")
	}

	district, err = districtCase.repo.Update(updateDistrict)
	if err != nil {
		return models.District{}, err
	}
	return district, err
}

func (districtCase *DistrictCase) Delete(id int64) error {
	var err error

	_, err = districtCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = districtCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
