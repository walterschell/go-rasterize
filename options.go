package rasterize

import "image/color"

// SVG Rendering Options
type Options struct {
	UniformDrawColor color.Color
	BackgroundColor  color.Color
}

func DefaultOptions() *Options {
	return &Options{
		BackgroundColor: color.Transparent,
	}
}

type Option func(*Options) error

// Uses a solid background color
func WithBackgroundColor(color color.Color) Option {
	return func(o *Options) error {
		o.BackgroundColor = color
		return nil
	}
}

// Forces the draw color to be a single color regarless of what
// is in the svg
func WithUniformColor(color color.Color) Option {
	return func(o *Options) error {
		o.UniformDrawColor = color
		return nil
	}
}

func applyOptions(opts []Option) (*Options, error) {
	options := DefaultOptions()
	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}
	return options, nil
}
