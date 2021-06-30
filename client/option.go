package client

type option struct {
	fallback bool
	uidList  []uint64
}

type OpOption func(opt *option)

func WithFallback() OpOption {
	return func(opt *option) {
		opt.fallback = true
	}
}
