package ctxgroup

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Group struct {
	*errgroup.Group
	ctx context.Context
}

func WithContext(ctx context.Context) Group {
	eg, ctx := errgroup.WithContext(ctx)
	return Group{
		Group: eg,
		ctx:   ctx,
	}
}

func (g Group) Wait() error {
	if err := g.Group.Wait(); err != nil {
		return err
	}
	return g.ctx.Err()
}

func (g Group) GoCtx(f func(ctx context.Context) error) {
	g.Go(func() error {
		return f(g.ctx)
	})
}
