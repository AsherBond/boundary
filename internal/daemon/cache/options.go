// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cache

import (
	"github.com/hashicorp/go-dbw"
)

type options struct {
	withIdContains            string
	withNameContains          string
	withDescriptionContains   string
	withAddressContains       string
	withIdStartsWith          string
	withNameStartsWith        string
	withDescriptionStartsWith string
	withAddressStartsWith     string
	withIdEndsWith            string
	withNameEndsWith          string
	withDescriptionEndsWith   string
	withAddressEndsWith       string
	withDebug                 bool
	withUrl                   string
	withDbType                dbw.DbType
}

// Option - how options are passed as args
type Option func(*options) error

func getDefaultOptions() options {
	return options{
		withDbType: dbw.Sqlite,
	}
}

func getOpts(opt ...Option) (options, error) {
	opts := getDefaultOptions()

	for _, o := range opt {
		if err := o(&opts); err != nil {
			return opts, err
		}
	}
	return opts, nil
}

// WithUrls provides optional url
func WithUrl(url string) Option {
	const op = "cache.WithUrl"
	return func(o *options) error {
		o.withUrl = url
		return nil
	}
}

// WithNameContains provides an optional name contains value.
func WithNameContains(value string) Option {
	return func(o *options) error {
		o.withNameContains = value
		return nil
	}
}

// WithNameStartsWith provides an optional name starts with value.
func WithNameStartsWith(value string) Option {
	return func(o *options) error {
		o.withNameStartsWith = value
		return nil
	}
}

// WithNameEndsWith provides an optional name ends with value.
func WithNameEndsWith(value string) Option {
	return func(o *options) error {
		o.withNameEndsWith = value
		return nil
	}
}

// WithDescriptionContains provides an optional description contains value.
func WithDescriptionContains(value string) Option {
	return func(o *options) error {
		o.withDescriptionContains = value
		return nil
	}
}

// WithDescriptionStartsWith provides an optional description starts with value.
func WithDescriptionStartsWith(value string) Option {
	return func(o *options) error {
		o.withDescriptionStartsWith = value
		return nil
	}
}

// WithDescriptionEndsWith provides an optional description ends with value.
func WithDescriptionEndsWith(value string) Option {
	return func(o *options) error {
		o.withDescriptionEndsWith = value
		return nil
	}
}

// WithIdContains provides an optional id contains value.
func WithIdContains(value string) Option {
	return func(o *options) error {
		o.withIdContains = value
		return nil
	}
}

// WithIdStartsWith provides an optional id starts with value.
func WithIdStartsWith(value string) Option {
	return func(o *options) error {
		o.withIdStartsWith = value
		return nil
	}
}

// WithIdEndsWith provides an optional id ends with value.
func WithIdEndsWith(value string) Option {
	return func(o *options) error {
		o.withIdEndsWith = value
		return nil
	}
}

// WithAddressContains provides an optional address contains value.
func WithAddressContains(value string) Option {
	return func(o *options) error {
		o.withAddressContains = value
		return nil
	}
}

// WithAddressStartsWith provides an optional address starts with value.
func WithAddressStartsWith(value string) Option {
	return func(o *options) error {
		o.withAddressStartsWith = value
		return nil
	}
}

// WithAddressEndsWith provides an optional address ends with value.
func WithAddressEndsWith(value string) Option {
	return func(o *options) error {
		o.withAddressEndsWith = value
		return nil
	}
}

// WithDbType provides an optional db type.
func WithDbType(dbType dbw.DbType) Option {
	return func(o *options) error {
		o.withDbType = dbType
		return nil
	}
}

// WithDebug provides an optional debug flag.
func WithDebug(debug bool) Option {
	return func(o *options) error {
		o.withDebug = debug
		return nil
	}
}