package main

import (
	"context"

	"github.com/segfault88/graphqltalk/store"
)

func (r *RootResolver) Star(ctx context.Context, args struct{ StarID int32 }) *StarResolver {
	return &StarResolver{r, r.storage.GetStars(int(args.StarID))[0]}
}

func (r *RootResolver) Stars(ctx context.Context) []*StarResolver {
	stars := r.storage.GetAllStars()
	resolvers := []*StarResolver{}
	for _, star := range stars {
		resolvers = append(resolvers, &StarResolver{r, star})
	}
	return resolvers
}

type StarResolver struct {
	root *RootResolver
	star store.Star
}

func (s *StarResolver) ID() int32 {
	return int32(s.star.ID)
}

func (s *StarResolver) Name() string {
	return s.star.Name
}

func (s *StarResolver) MovieIDs(ctx context.Context) []int32 {
	ids := []int32{}
	for _, id := range s.star.Movies {
		ids = append(ids, int32(id))
	}
	return ids
}

func (s *StarResolver) Movies(ctx context.Context) []*MovieResolver {
	resolvers := []*MovieResolver{}
	for _, movieID := range s.star.Movies {
		resolvers = append(resolvers, s.root.Movie(ctx, struct{ MovieID int32 }{int32(movieID)}))
	}
	return resolvers
}
