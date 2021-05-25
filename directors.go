package main

import (
	"context"

	"github.com/segfault88/graphqltalk/store"
)

func (r *RootResolver) Director(ctx context.Context, args struct{ DirectorID int32 }) *DirectorResolver {
	return &DirectorResolver{r, r.storage.GetDirectors(int(args.DirectorID))[0]}
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

func (d *DirectorResolver) Movies(ctx context.Context) []*MovieResolver {
	resolvers := []*MovieResolver{}
	for _, movieID := range d.director.Movies {
		resolvers = append(resolvers, d.root.Movie(ctx, struct{ MovieID int32 }{int32(movieID)}))
	}
	return resolvers
}
