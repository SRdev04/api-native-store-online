package execption

type NotFoundError struct {
	Error string
}

func NewNotFound(err string) NotFoundError {
	return NotFoundError{Error: err}

}
