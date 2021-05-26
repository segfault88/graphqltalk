# GraphQL talk

Basic implementation of a GraphQL in Go easily showing the N+1 problem. This is because the ground breaking data store implemented here is unfortunately horribly slow.

See loader branch for an upgrade with github.com/vektah/dataloaden that improves the situation drastically.


Generator run with:

```bash
▶ go run github.com/vektah/dataloaden MovieLoader int github.com/segfault88/graphqltalk/store.Movie
```
