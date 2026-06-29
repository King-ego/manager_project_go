package repositories

import "gorm.io/gorm"

type DashboardRepository interface {
	GetDashboard() (any, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

func (dr *dashboardRepository) GetDashboard() (any, error) {
	var result any
	err := dr.db.Raw("SELECT * FROM dashboard_data").Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
