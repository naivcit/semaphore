package maestro

import (
	"context"

	"github.com/jexia/maestro/pkg/core/api"
	"github.com/jexia/maestro/pkg/core/instance"
	"github.com/jexia/maestro/pkg/flow"
	"github.com/jexia/maestro/pkg/refs"
)

// WithMiddleware initialises the given middleware and defines all options
func WithMiddleware(middleware api.Middleware) api.Option {
	return func(options *api.Options) {
		options.Middleware = append(options.Middleware, middleware)
	}
}

// AfterConstructor the passed function gets called once all options have been applied
func AfterConstructor(wrapper api.AfterConstructorHandler) api.Option {
	return func(options *api.Options) {
		if options.AfterConstructor == nil {
			options.AfterConstructor = wrapper(func(instance.Context, *api.Collection) error { return nil })
			return
		}

		options.AfterConstructor = wrapper(options.AfterConstructor)
	}
}

// BeforeManagerDo the passed function gets called before a request gets handled by a flow manager
func BeforeManagerDo(wrapper flow.BeforeManagerHandler) api.Option {
	return func(options *api.Options) {
		if options.BeforeManagerDo == nil {
			options.BeforeManagerDo = wrapper(func(ctx context.Context, manager *flow.Manager, store refs.Store) (context.Context, error) {
				return ctx, nil
			})

			return
		}

		options.BeforeManagerDo = wrapper(options.BeforeManagerDo)
	}
}

// BeforeManagerRollback the passed function gets called before a rollback request gets handled by a flow manager
func BeforeManagerRollback(wrapper flow.BeforeManagerHandler) api.Option {
	return func(options *api.Options) {
		if options.BeforeManagerRollback == nil {
			options.BeforeManagerRollback = wrapper(func(ctx context.Context, manager *flow.Manager, store refs.Store) (context.Context, error) {
				return ctx, nil
			})

			return
		}

		options.BeforeManagerRollback = wrapper(options.BeforeManagerRollback)
	}
}

// AfterManagerDo the passed function gets after a flow has been handled by the flow manager
func AfterManagerDo(wrapper flow.AfterManagerHandler) api.Option {
	return func(options *api.Options) {
		if options.AfterManagerDo == nil {
			options.AfterManagerDo = wrapper(func(ctx context.Context, manager *flow.Manager, store refs.Store) (context.Context, error) {
				return ctx, nil
			})

			return
		}

		options.AfterManagerDo = wrapper(options.AfterManagerDo)
	}
}

// AfterManagerRollback the passed function gets after a flow rollback has been handled by the flow manager
func AfterManagerRollback(wrapper flow.AfterManagerHandler) api.Option {
	return func(options *api.Options) {
		if options.AfterManagerRollback == nil {
			options.AfterManagerRollback = wrapper(func(ctx context.Context, manager *flow.Manager, store refs.Store) (context.Context, error) {
				return ctx, nil
			})

			return
		}

		options.AfterManagerRollback = wrapper(options.AfterManagerRollback)
	}
}

// BeforeNodeDo the passed function gets called before a node is executed
func BeforeNodeDo(wrapper flow.BeforeNodeHandler) api.Option {
	return func(options *api.Options) {
		if options.BeforeNodeDo == nil {
			options.BeforeNodeDo = wrapper(func(ctx context.Context, node *flow.Node, tracker *flow.Tracker, processes *flow.Processes, store refs.Store) (context.Context, error) {
				return ctx, nil
			})

			return
		}

		options.BeforeNodeDo = wrapper(options.BeforeNodeDo)
	}
}

// BeforeNodeRollback the passed function gets called before a node rollback is executed
func BeforeNodeRollback(wrapper flow.BeforeNodeHandler) api.Option {
	return func(options *api.Options) {
		if options.BeforeNodeRollback == nil {
			options.BeforeNodeRollback = wrapper(func(ctx context.Context, node *flow.Node, tracker *flow.Tracker, processes *flow.Processes, store refs.Store) (context.Context, error) {
				return ctx, nil
			})

			return
		}

		options.BeforeNodeRollback = wrapper(options.BeforeNodeRollback)
	}
}

// AfterNodeDo the passed function gets called after a node is executed
func AfterNodeDo(wrapper flow.AfterNodeHandler) api.Option {
	return func(options *api.Options) {
		if options.AfterNodeDo == nil {
			options.AfterNodeDo = wrapper(func(ctx context.Context, node *flow.Node, tracker *flow.Tracker, processes *flow.Processes, store refs.Store) (context.Context, error) {
				return ctx, nil
			})
			return
		}

		options.AfterNodeDo = wrapper(options.AfterNodeDo)
	}
}

// AfterNodeRollback the passed function gets called after a node rollback is executed
func AfterNodeRollback(wrapper flow.AfterNodeHandler) api.Option {
	return func(options *api.Options) {
		if options.AfterNodeRollback == nil {
			options.AfterNodeRollback = wrapper(func(ctx context.Context, node *flow.Node, tracker *flow.Tracker, processes *flow.Processes, store refs.Store) (context.Context, error) {
				return ctx, nil
			})
			return
		}

		options.AfterNodeRollback = wrapper(options.AfterNodeRollback)
	}
}
