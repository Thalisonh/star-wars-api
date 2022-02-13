package middlewares

type SwapiServiceSpy struct {
	SwapiService
	TotalFilmsStatusCode int
	TotalFilmsCount      int
	TotalFilmsError      error
}

func (s *SwapiServiceSpy) TotalFilms(planetName string) (int, int, error) {
	return s.TotalFilmsStatusCode, s.TotalFilmsCount, s.TotalFilmsError
}
