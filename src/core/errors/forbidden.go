package core

// ForbiddenError ...
type ForbiddenError struct{}

// Forbidden ...
func (*ForbiddenError) Error() string {
	return "403 - Forbidden Error"
}
