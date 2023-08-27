package chijson

func WithSuccessStatus(statusCode int) HandlerOption {
	return func(cfg *HandlerConfig) {
		cfg.SuccessStatusCode = statusCode
	}
}
