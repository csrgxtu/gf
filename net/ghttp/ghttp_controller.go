// Copyright 2017 gf Author(https://github.com/csrgxtu/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/csrgxtu/gf.

package ghttp

// Controller is the base struct for controller.
type Controller interface {
	Init(*Request)
	Shut()
}
