package errorx

import "reflect"

// Error error
type Error string

func (e Error) Error() string {
	return reflect.ValueOf(e).String()
}

const (
	// SDKProjectNameGetFailed project name get failed
	SDKProjectNameGetFailed Error = "project name get failed"
	// SDKProjectTypeGetFailed project type get failed
	SDKProjectTypeGetFailed Error = "project type get failed"
	// SDKProjectInitFailed project init failed
	SDKProjectInitFailed Error = "project init failed"
	// SDKProjectAlreadyExist project already exists
	SDKProjectAlreadyExist Error = "project already exists"
)
