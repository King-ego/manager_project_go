package usecases

type ScoreUseCase struct {
}

func NewScoreUseCase() *ScoreUseCase {
	return &ScoreUseCase{}
}

func (s *ScoreUseCase) CreateScore() (string, error) {
	return "", nil
}
