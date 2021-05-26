package store

import (
	"time"

	"github.com/rs/zerolog/log"
)

const delay = 200 * time.Millisecond

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (*Service) GetMovies(movieIDs ...int) []Movie {
	log.Info().Msgf("looking up movie by ids: %v", movieIDs)
	time.Sleep(delay)
	found := []Movie{}
	for _, movieID := range movieIDs {
	inner:
		for _, m := range movies {
			if m.ID == movieID {
				found = append(found, m)
				break inner
			}
		}
	}
	return found
}

func (*Service) GetAllMovies() []Movie {
	log.Info().Msg("looking up all movies")
	time.Sleep(delay)
	return movies
}

func (*Service) GetDirectors(directorIDs ...int) []Director {
	log.Info().Msgf("looking up director by ids: %v", directorIDs)
	time.Sleep(delay)
	found := []Director{}
	for _, directorID := range directorIDs {
	inner:
		for _, d := range directors {
			if d.ID == directorID {
				found = append(found, d)
				break inner
			}
		}
	}
	return found
}

func (*Service) GetAllDirectors() []Director {
	log.Info().Msg("looking up all directors")
	time.Sleep(delay)
	return directors
}

func (*Service) GetStars(starIDs ...int) []Star {
	log.Info().Msgf("looking up star by ids: %v", starIDs)
	time.Sleep(delay)
	found := []Star{}
	for _, starID := range starIDs {
	inner:
		for _, s := range stars {
			if s.ID == starID {
				found = append(found, s)
				break inner
			}
		}
	}
	return found
}

func (*Service) GetAllStars() []Star {
	log.Info().Msg("looking up all stars")
	time.Sleep(delay)
	return stars
}
