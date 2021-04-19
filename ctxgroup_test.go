package ctxgroup_test

import (
	"context"
	"errors"
	"testing"

	"github.com/pyr-sh/ctxgroup"
)

func TestContextCancellation(t *testing.T) {
	ctx := context.Background()
	g := ctxgroup.WithContext(ctx)
	g.GoCtx(func(ctx context.Context) error {
		<-ctx.Done()
		return ctx.Err()
	})
	g.GoCtx(func(ctx context.Context) error {
		return errors.New("hello world")
	})
	err := g.Wait()
	if err == nil || err.Error() != "hello world" {
		t.Fatalf(`expected the error to be "hello world", got %v (%T)`, err, err)
	}
}

func TestContextAlreadyCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	g := ctxgroup.WithContext(ctx)
	cancel()
	g.GoCtx(func(ctx context.Context) error {
		<-ctx.Done()
		return ctx.Err()
	})
	err := g.Wait()
	if err == nil || err != context.Canceled {
		t.Fatalf(`expected the error to be "context canceled", got %v (%T)`, err, err)
	}
}

func TestWaitErrorOnFunctionSuccess(t *testing.T) {
	ctx := context.Background()
	g := ctxgroup.WithContext(ctx)
	g.GoCtx(func(ctx context.Context) error {
		return nil
	})
	if err := g.Wait(); err != nil {
		t.Fatalf("expected error to be nil, got %v", err)
	}
}
