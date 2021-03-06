// Copyright 2018 gf Author(https://github.com/csrgxtu/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/csrgxtu/gf.

package ghttp_test

import (
	"github.com/csrgxtu/gf/container/garray"
	"github.com/csrgxtu/gf/os/genv"
)

var (
	ports = garray.NewIntArray(true)
)

func init() {
	genv.Set("UNDER_TEST", "1")
	for i := 8000; i <= 9000; i++ {
		ports.Append(i)
	}
}
