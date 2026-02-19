package domain

import "fmt"

// IntegrationErrorKind identifies integration failure categories.
type IntegrationErrorKind string

const (
	ErrKindTimeout        IntegrationErrorKind = "timeout"
	ErrKindUpstreamStatus IntegrationErrorKind = "upstream_status"
	ErrKindDecode         IntegrationErrorKind = "decode"
	ErrKindInvalidPayload IntegrationErrorKind = "invalid_payload"
)

// IntegrationError wraps integration failures in a consistent domain error.
type IntegrationError struct {
	Kind       IntegrationErrorKind
	Operation  string
	StatusCode int
	Cause      error
}

func (e *IntegrationError) Error() string {
	if e == nil {
		return "integration error"
	}
	if e.StatusCode > 0 {
		return fmt.Sprintf("integration error kind=%s operation=%s status=%d", e.Kind, e.Operation, e.StatusCode)
	}
	return fmt.Sprintf("integration error kind=%s operation=%s", e.Kind, e.Operation)
}

func (e *IntegrationError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}
