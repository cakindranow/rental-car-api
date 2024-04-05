package exception

type ForbiddenError struct {
	err string
}

func (e ForbiddenError) Error() string {
	return e.err
}

func NewForbiddenError(error string) ForbiddenError {
	return ForbiddenError{err: error}
}
