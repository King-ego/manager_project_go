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

func (dr *dashboardRepository) GetStats() (any, error) {
	var result any
	err := dr.db.Raw("SELECT * FROM dashboard_stats").Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dr *dashboardRepository) GetSummary() (any, error) {
	var result any
	err := dr.db.Raw("SELECT * FROM dashboard_summary").Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dr *dashboardRepository) GetMetrics() (any, error) {
	var result any
	err := dr.db.Raw("SELECT * FROM dashboard_metrics").Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
