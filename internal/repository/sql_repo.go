package repository

import (
	"gin-boilerplate/infra/logger"
)

func Save(model interface{}) interface{} {
	err := DB.Create(model).Error
	if err != nil {
		logger.Errorf("error, not save data %v", err)
	}
	return err
}

func Get(model interface{}) interface{} {
	err := DB.Find(model).Error
	return err
}

func GetOne(model interface{}) interface{} {
	err := DB.Last(model).Error
	return err
}

func Update(model interface{}) interface{} {
	err := DB.Find(model).Error
	return err
}
