package usecase

import (
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"gorm.io/gorm"
	"strconv"
)

type ClientTypeClassifierCase struct {
	repo interfaces.ClientTypeClassifierRepoInterface
}

func (classifierCase *ClientTypeClassifierCase) Create(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error) {
	if _, err := classifierCase.repo.GetByCode(strconv.FormatInt(classifier.Code, 10)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ClientTypeClassifier{}, errors.New("the code has already been taken")
	}

	return classifierCase.repo.Create(classifier)
}

func (classifierCase *ClientTypeClassifierCase) Get(id int64) (models.ClientTypeClassifier, error) {
	var classifier models.ClientTypeClassifier
	var err error

	classifier, err = classifierCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ClientTypeClassifier{}, errors.New("no such classifier with the id supplied")
		}

		return models.ClientTypeClassifier{}, err
	}
	return classifier, nil
}

func (classifierCase *ClientTypeClassifierCase) Update(updateClassifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error) {
	var classifier models.ClientTypeClassifier
	var err error

	classifier, err = classifierCase.repo.Get(int64(updateClassifier.ID))
	if err != nil {
		return models.ClientTypeClassifier{}, err
	}

	if classifier.Code != updateClassifier.Code {
		return models.ClientTypeClassifier{}, errors.New("code cannot be changed")
	}

	classifier, err = classifierCase.repo.Update(updateClassifier)
	if err != nil {
		return models.ClientTypeClassifier{}, err
	}
	return classifier, err
}

func (classifierCase *ClientTypeClassifierCase) Delete(id int64) error {
	var err error

	_, err = classifierCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = classifierCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
