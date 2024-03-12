package service

import "PlusOne/models"

type TestService struct {
}

func (s *TestService) Test() []models.Test {
	test := models.Test{}
	return test.GetAll()
}
