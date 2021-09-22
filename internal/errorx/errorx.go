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

	// SDKConfigSyncFailed config auto-sync failed
	SDKConfigSyncFailed Error = "config auto-sync failed"
	// SDKConfigFolderNotExist config folder not exist
	SDKConfigFolderNotExist Error = "config folder not exist"
	// SDKConfigWatchFailed config watch failed
	SDKConfigWatchFailed Error = "config watch failed"
	// SDKConfigGenerateFailed config generate failed
	SDKConfigGenerateFailed Error = "config generate failed"
)
