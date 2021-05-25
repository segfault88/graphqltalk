package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/segfault88/graphqltalk/store"
)

func (r *RootResolver) Movie(ctx context.Context, args struct{ MovieID int32 }) *MovieResolver {
	return &MovieResolver{r, r.storage.GetMovies(int(args.MovieID))[0]}
}

func (r *RootResolver) Movies(ctx context.Context) []*MovieResolver {
	movies := r.storage.GetAllMovies()
	resolvers := []*MovieResolver{}
	spew.Dump(movies)
	for _, movie := range movies {
		resolvers = append(resolvers, &MovieResolver{r, movie})
	}
	return resolvers
}

type MovieResolver struct {
	root  *RootResolver
	movie store.Movie
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

func (m *MovieResolver) Director(ctx context.Context) *DirectorResolver {
	return m.root.Director(ctx, struct{ DirectorID int32 }{int32(m.movie.DirectorID)})
}

func (m *MovieResolver) StarIDs(ctx context.Context) []int32 {
	ids := []int32{}
	for _, id := range m.movie.Stars {
		ids = append(ids, int32(id))
	}
	return ids
}

func (m *MovieResolver) Stars(ctx context.Context) []*StarResolver {
	resolvers := []*StarResolver{}
	for _, starID := range m.movie.Stars {
		resolvers = append(resolvers, m.root.Star(ctx, struct{ StarID int32 }{int32(starID)}))
	}
	return resolvers
}
