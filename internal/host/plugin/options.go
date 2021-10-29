package plugin

import "google.golang.org/protobuf/types/known/structpb"

// getOpts - iterate the inbound Options and return a struct
func getOpts(opt ...Option) options {
	opts := options{}
	for _, o := range opt {
		o(&opts)
	}
	return opts
}

// Option - how Options are passed as arguments.
type Option func(*options)

// options = how options are represented
type options struct {
	withPluginId           string
	withName               string
	withDescription        string
	withAttributes         *structpb.Struct
	withSecrets            *structpb.Struct
	withPreferredEndpoints []string
	withIpAddresses        []string
	withDnsNames           []string
	withLimit              int
}

// WithPluginId provides an optional plugin id.
func withPluginId(with string) Option {
	return func(o *options) {
		o.withPluginId = with
	}
}

// WithDescription provides an optional description.
func WithDescription(desc string) Option {
	return func(o *options) {
		o.withDescription = desc
	}
}

// WithName provides an optional name.
func WithName(name string) Option {
	return func(o *options) {
		o.withName = name
	}
}

// WithAttributes provides an optional attributes field.
func WithAttributes(attrs *structpb.Struct) Option {
	return func(o *options) {
		o.withAttributes = attrs
	}
}

// WithSecrets provides an optional secrets field.
func WithSecrets(secrets *structpb.Struct) Option {
	return func(o *options) {
		o.withSecrets = secrets
	}
}

// WithPreferredEndpoints provides an optional preferred endpoints field.
func WithPreferredEndpoints(with []string) Option {
	return func(o *options) {
		o.withPreferredEndpoints = with
	}
}

// withIpAddresses provides an optional list of ip addresses.
func withIpAddresses(with []string) Option {
	return func(o *options) {
		o.withIpAddresses = with
	}
}

// withDnsNames provides an optional list of dns names.
func withDnsNames(with []string) Option {
	return func(o *options) {
		o.withDnsNames = with
	}
}

// WithLimit provides an option to provide a limit. Intentionally allowing
// negative integers. If WithLimit < 0, then unlimited results are
// returned. If WithLimit == 0, then default limits are used for results.
func WithLimit(l int) Option {
	return func(o *options) {
		o.withLimit = l
	}
}
