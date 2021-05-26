package main

import (
	"context"

	"github.com/segfault88/graphqltalk/store"
)

func (r *RootResolver) Star(ctx context.Context, args struct{ StarID int32 }) (*StarResolver, error) {
	star, err := r.loaders.Directors.Load(int(args.StarID))
	if err != nil {
		return nil, err
	}
	return &StarResolver{r, store.Star(star)}, nil
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

func newStarResolversLookup(root *RootResolver, starIDs []int) ([]*StarResolver, error) {
	stars, errs := root.loaders.Stars.LoadAll(starIDs)
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}
	resolvers := []*StarResolver{}
	for _, star := range stars {
		resolvers = append(resolvers, &StarResolver{root, star})
	}
	return resolvers, nil
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

func (s *StarResolver) Movies(ctx context.Context) ([]*MovieResolver, error) {
	return newMovieResolversLookup(s.root, s.star.Movies)
}
