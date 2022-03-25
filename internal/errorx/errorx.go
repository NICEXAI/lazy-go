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
	// SDKProjectNotExist project not exist
	SDKProjectNotExist Error = "project not exist"

	// SDKPkgNotExist pkg folder not exist
	SDKPkgNotExist Error = "pkg folder not exist"

	// SDKConfigOriginFileFormatNotSupport origin file format is not supported
	SDKConfigOriginFileFormatNotSupport Error = "origin file format is not supported"
	// SDKConfigTargetFileFormatNotSupport target file format is not supported
	SDKConfigTargetFileFormatNotSupport Error = "target file format is not supported"
	// SDKConfigOriginFileNotExist origin file is not exist
	SDKConfigOriginFileNotExist Error = "origin file is not exist"
	// SDKConfigGenerateFailed config generate failed
	SDKConfigGenerateFailed Error = "config generate failed"
)
