package usecases

import "manager_project/repositories"

type DashboardUseCase struct {
	repository repositories.DashboardRepository
}

func NewDashboardUseCase(repository repositories.DashboardRepository) *DashboardUseCase {
	return &DashboardUseCase{
		repository: repository,
	}
}

func (d *DashboardUseCase) GetDashboard() (any, error) {
	dashboard, err := d.repository.GetDashboard()
	if err != nil {
		return nil, err
	}
	return dashboard, nil
}
