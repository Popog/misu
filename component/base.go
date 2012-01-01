// Copyright 2011 The Misu Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package component

// base is a useful class to embedded because it takes care of the Initialized and Destroyed for you
type Base struct {
	initialized, destroyed bool
}

// Initialize finializes the initialization of a component.
// If a component does not call Component.Initialize() at the end of their overloaded Initialize()
// that component's Initialize will be called again next initialization loop
func (this *Base) Initialize(parent Collection) {
	if this.initialized {
		panic("Component has already been initialized")
	}
	this.initialized = true
}

// Initialized returns whether or not a component has called Initialize
func (this *Base) Initialized() bool { return this.initialized }

// Destroy finializes the destruction of a component.
// If a component does not call Component.Destroy() at the end of their overloaded Destroy()
// that component's Destroy will be called again next initialization loop
func (this *Base) Destroy(parent Collection) {
	if this.destroyed {
		panic("Component has already been destroyed")
	}
	this.destroyed = true
}

// Destroyed returns whether or not a component has called Initialize
func (this *Base) Destroyed() bool { return this.destroyed }
