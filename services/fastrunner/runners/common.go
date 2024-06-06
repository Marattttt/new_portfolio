package runners

import "errors"

// Result of running user-provided code
// If code could not be run, a separate error should be returned
type RunResult struct {
	// Code output
	Text string
	// Error with running code
	Err error
}

func IsClientError(err error) bool {
	return errors.Is(err, CodeError{}) || errors.Is(err, RuntimeError{})
}

type CodeError struct {
	Err error
}

func (ce CodeError) Error() string {
	return ce.Err.Error()
}

func (ce CodeError) Unwrap() error {
	return ce.Err
}

type RuntimeError struct {
	Msg *string
	Err error
}

func (re RuntimeError) Error() string {
	resp := ""
	if re.Msg == nil {
		resp += *re.Msg
	}
	return resp + re.Err.Error()
}

func (re RuntimeError) Unwrap() error {
	return re.Err
}
