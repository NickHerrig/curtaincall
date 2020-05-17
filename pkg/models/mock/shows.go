package mock

import (
	"curtaincall.tech/pkg/models"
)

var mockShow = &models.Show{
	ID:      1,
	Name:    "Hamilton",
	Company: "Company A",
}

type ShowModel struct{}

func (m *ShowModel) Get(id int) (*models.Show, error) {
	switch id {
	case 1:
		return mockShow, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *ShowModel) Latest(id int) ([]*models.Show, error) {
	return []*models.Show{mockShow}, nil
}
