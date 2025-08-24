package routers

type EnrollmentsRouters struct{}

func NewEnrollmentsRoutes() *EnrollmentsRouters {
	return &EnrollmentsRouters{}
}

func (er *EnrollmentsRouters) setupEnrollmentsRoutes() string {
	return "Enrollments routes set up"
}

func (er *EnrollmentsRouters) EnrollmentsRoutes() {
	er.setupEnrollmentsRoutes()
}

func SetupEnrollmentsRoutes() {
	router := NewEnrollmentsRoutes()
	router.EnrollmentsRoutes()
}
