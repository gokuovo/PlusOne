package models

type Test struct {
	ID   int64
	Name string
}

func (m *Test) TableName() string {
	return "test"
}

func (m *Test) GetAll() []Test {
	db := DB
	var test []Test
	db.Table(m.TableName()).Find(&test)
	return test
}
