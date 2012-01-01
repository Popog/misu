// Copyright 2011 The Misu Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transform2d

import "github.com/Popog/misu/component"
import math3d "github.com/Popog/math3d/math3d64"

var Type component.Type = component.RegisterComponentType("transform2d", cFactory, Data{})

func init() {
	// This Component needs no types, but
	component.RegisterInitializeTypes(Type)
	component.RegisterDestroyTypes(Type)
}

type Component struct {
	*component.Base
	Translation math3d.Vector2
	Rotation    math3d.Vector2 // cos, sin
	Scale       math3d.Vector2
}

type Data struct {
	Translation *math3d.Vector2
	Rotation    *float64
	Scale       *math3d.Vector2
}

func cFactory() component.Component { return &Component{Base: new(component.Base)} }

func (Component) GetType() component.Type { return Type }

func (this *Component) LoadData(data component.Data) {
	d := data.(*Data)
	if d.Translation != nil {
		this.Translation = *d.Translation
	}
	if d.Rotation != nil {
		this.Rotation = math3d.MakeVector2(math3d.Cosf(*d.Rotation), math3d.Sinf(*d.Rotation))
	}
	if d.Scale != nil {
		this.Scale = *d.Scale
	}
}

func (this *Component) Initialize(parent component.Collection) {
	this.Base.Initialize(parent)
}

func (this *Component) Destroy(parent component.Collection) {
	this.Base.Destroy(parent)
}

func (Data) GetType() component.Type { return Type }