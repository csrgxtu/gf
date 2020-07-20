// Copyright 2018 gf Author(https://github.com/csrgxtu/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/csrgxtu/gf.

package g

import (
	"github.com/csrgxtu/gf/os/glog"
)

// SetDebug disables/enables debug level for logging component globally.
func SetDebug(debug bool) {
	glog.SetDebug(debug)
}

// SetLogLevel sets the logging level globally.
func SetLogLevel(level int) {
	glog.SetLevel(level)
}

// GetLogLevel returns the global logging level.
func GetLogLevel() int {
	return glog.GetLevel()
}
