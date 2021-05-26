package main

import (
	"context"

	"github.com/segfault88/graphqltalk/store"
)

func (r *RootResolver) Director(ctx context.Context, args struct{ DirectorID int32 }) (*DirectorResolver, error) {
	director, err := r.loaders.Directors.Load(int(args.DirectorID))
	if err != nil {
		return nil, err
	}
	return &DirectorResolver{r, director}, nil
}

func (r *RootResolver) Directors(ctx context.Context) []*DirectorResolver {
	directors := r.storage.GetAllDirectors()
	resolvers := []*DirectorResolver{}
	for _, director := range directors {
		resolvers = append(resolvers, &DirectorResolver{r, director})
	}
	return resolvers
}

type DirectorResolver struct {
	root     *RootResolver
	director store.Director
}

func (d *DirectorResolver) ID() int32 {
	return int32(d.director.ID)
}

func (d *DirectorResolver) Name() string {
	return d.director.Name
}

func (d *DirectorResolver) MovieIDs() []int32 {
	ids := []int32{}
	for _, movieID := range d.director.Movies {
		ids = append(ids, int32(movieID))
	}
	return ids
}

func (d *DirectorResolver) Movies(ctx context.Context) ([]*MovieResolver, error) {
	return newMovieResolversLookup(d.root, d.director.Movies)
}
