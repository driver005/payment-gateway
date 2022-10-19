package driver

import (
	"context"

	"github.com/driver005/gateway/logger"
	"github.com/driver005/gateway/registry"
)

type options struct {
	forcedValues map[string]interface{}
	preload      bool
	validate     bool
}

func newOptions() *options {
	return &options{
		validate: true,
		preload:  true,
	}
}

type OptionsModifier func(*options)

// DisableValidation validating the config.
//
// This does not affect schema validation!
func DisableValidation() OptionsModifier {
	return func(o *options) {
		o.validate = false
	}
}

// DisableValidation validating the config.
//
// This does not affect schema validation!
func DisablePreloading() OptionsModifier {
	return func(o *options) {
		o.preload = false
	}
}

func New(ctx context.Context) registry.Registry {
	l := logger.New("ORY Hydra", "2.1")

	r, err := registry.NewRegistryFromDSN(ctx, l)
	if err != nil {
		l.WithError(err).Fatal("Unable to create service registry.")
	}

	if err = r.Init(ctx); err != nil {
		l.WithError(err).Fatal("Unable to initialize service registry.")
	}

	return r
}
