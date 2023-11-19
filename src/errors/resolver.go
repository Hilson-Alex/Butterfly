package errors

type Resolver[T any] func() (T, error)

type Runner func() error

func (resolver Resolver[T]) OrPanicOn(logger *BfErrLogger) T {
	var result, err = resolver()
	if err != nil {
		logger.Panic(err)
	}
	return result
}

func (resolver Resolver[T]) AndHandle(handler func(err error)) T {
	var result, err = resolver()
	if err != nil {
		handler(err)
	}
	return result
}

func (runner Runner) OrPanicOn(logger *BfErrLogger) {
	if err := runner(); err != nil {
		logger.Panic(err)
	}
}

func (runner Runner) AndHandle(handler func(err error)) {
	if err := runner(); err != nil {
		handler(err)
	}
}

func Resolve[T any](resolver func() (T, error)) Resolver[T] {
	return resolver
}

func Run(runner func() error) Runner {
	return runner
}
