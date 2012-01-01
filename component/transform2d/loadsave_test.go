// Copyright 2011 The Misu Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transform2d

import "testing"
import "fmt"
import "encoding/gob"
import "bytes"
import "github.com/Popog/misu/component"

func TestRegisterComponentType(t *testing.T) {
	if !component.RegistrationFinalized() {
		component.FinalizeRegistration()
	}
	
	buf := bytes.NewBuffer(nil)
	
	{
		var d component.Data = Data{Rotation : &([]float64{67}[0])}
		enc := gob.NewEncoder(buf)
		enc.Encode(&d)
	}
	
	fmt.Println(buf)
	
	{
		dec := gob.NewDecoder(buf)
		
		d := component.Data(nil) // component.CreateData(t)
		dec.Decode(&d)
		fmt.Println(d)
		fmt.Println(*d.(Data).Rotation)
	}
}

