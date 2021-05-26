package store

import "time"

type Loaders struct {
	Movies    *MovieLoader
	Stars     *StarLoader
	Directors *DirectorLoader
}

func (s *Service) NewLoaders() Loaders {
	return Loaders{
		Movies: &MovieLoader{
			wait:     2 * time.Millisecond,
			maxBatch: 100,
			fetch: func(keys []int) ([]Movie, []error) {
				return s.GetMovies(keys...), nil
			},
		},
		Stars: &StarLoader{
			wait:     2 * time.Millisecond,
			maxBatch: 100,
			fetch: func(keys []int) ([]Star, []error) {
				return s.GetStars(keys...), nil
			},
		},
		Directors: &DirectorLoader{
			wait:     2 * time.Millisecond,
			maxBatch: 100,
			fetch: func(keys []int) ([]Director, []error) {
				return s.GetDirectors(keys...), nil
			},
		},
	}
}
