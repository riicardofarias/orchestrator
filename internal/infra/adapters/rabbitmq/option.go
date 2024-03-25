package rabbitmq

type option struct {
	retries int
}

type Option interface {
	apply(opt *option)
}

type optionFunc func(opt *option)

func (f optionFunc) apply(opt *option) { f(opt) }

func WithRetries(retries int) Option {
	return optionFunc(func(opt *option) {
		opt.retries = retries
	})
}
