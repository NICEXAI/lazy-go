package config

import (
	"github.com/NICEXAI/fstask"
	"github.com/NICEXAI/lazy-go/internal/errorx"
	"github.com/NICEXAI/lazy-go/util"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

// Watch listening file changes and auto convert content to Go Struct
func Watch(from, to, mod, tag string) (*fstask.FsTask, error) {
	var (
		fsTask *fstask.FsTask
		err    error
	)

	originFolderPath := util.GetFolderAbsPath(from)
	fsTask, err = fstask.New(originFolderPath)
	if err != nil {
		return nil, err
	}

	originFileName := util.GetFileFullName(from)
	if err = fsTask.Add(fstask.Task{
		Rule:   ".*" + originFileName,
		Action: []string{"create", "write"},
		Handle: func(event fsnotify.Event) {
			if err = Convert(from, to, mod, tag); err != nil {
				color.Red("%v: %v", errorx.SDKConfigGenerateFailed, err)
				return
			}
			color.Green("file conversion success")
		},
	}); err != nil {
		return nil, err
	}

	return fsTask, nil
}
