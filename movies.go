package main

import (
	"context"

	"github.com/segfault88/graphqltalk/store"
)

func (r *RootResolver) Movie(ctx context.Context, args struct{ MovieID int32 }) (*MovieResolver, error) {
	movie, err := r.loaders.Movies.Load(int(args.MovieID))
	if err != nil {
		return nil, err
	}
	return &MovieResolver{r, movie}, nil
}

func (r *RootResolver) Movies(ctx context.Context) []*MovieResolver {
	movies := r.storage.GetAllMovies()
	resolvers := []*MovieResolver{}
	for _, movie := range movies {
		resolvers = append(resolvers, &MovieResolver{r, movie})
	}
	return resolvers
}

type MovieResolver struct {
	root  *RootResolver
	movie store.Movie
}

func newMovieResolversLookup(root *RootResolver, movieIDs []int) ([]*MovieResolver, error) {
	movies, errs := root.loaders.Movies.LoadAll(movieIDs)
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}
	resolvers := []*MovieResolver{}
	for _, movie := range movies {
		resolvers = append(resolvers, &MovieResolver{root, movie})
	}
	return resolvers, nil
}

func (m *MovieResolver) ID() int32 {
	return int32(m.movie.ID)
}

func (m *MovieResolver) Name() string {
	return m.movie.Name
}

func (m *MovieResolver) Runtime() int32 {
	return int32(m.movie.Runtime)
}

func (m *MovieResolver) DirectorID() int32 {
	return int32(m.movie.DirectorID)
}

func (m *MovieResolver) Director(ctx context.Context) (*DirectorResolver, error) {
	return m.root.Director(ctx, struct{ DirectorID int32 }{int32(m.movie.DirectorID)})
}

func (m *MovieResolver) StarIDs(ctx context.Context) []int32 {
	ids := []int32{}
	for _, id := range m.movie.Stars {
		ids = append(ids, int32(id))
	}
	return ids
}

func (m *MovieResolver) Stars(ctx context.Context) ([]*StarResolver, error) {
	return newStarResolversLookup(m.root, m.movie.Stars)
}
