package errors

type ErrNotFound struct {
	error string
}

func (e *ErrNotFound) Error() string {
	return e.error
}

func NewErrNotFound(err string) error {
	return &ErrNotFound{error: err}
}

func WrapErrNotFound(err error) error {
	return &ErrNotFound{error: err.Error()}
}

type ErrNotSingular struct {
	error string
}

func (e *ErrNotSingular) Error() string {
	return e.error
}

func NewErrNotSingular(err string) error {
	return &ErrNotSingular{error: err}
}

func WrapErrNotSingular(err error) error {
	return &ErrNotSingular{error: err.Error()}
}
