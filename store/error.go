package store

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "receipt not found"
}
