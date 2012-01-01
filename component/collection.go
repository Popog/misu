// Copyright 2011 The Misu Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package component

// Collections are basically "Game Objects".
// They are access controlled however, and will error if an invalid component is accessed.
// They may also provide a default "garbage" version of that component.
type Collection interface {
	GetComponent(requested_type Type) interface{}
}
