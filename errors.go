package ef

// Error returned when an operation cannot be executed on the given file type.
type ErrInvalidFile struct{}

func (e *ErrInvalidFile) Error() string {
	return "invalid file type"
}
