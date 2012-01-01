// Copyright 2011 The Misu Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package component

import "sort"
import "fmt"

type TypeSlice []Type

func (p TypeSlice) Len() int           { return len(p) }
func (p TypeSlice) Less(i, j int) bool { return *p[i] < *p[j] }
func (p TypeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type RegistrationList struct {
	registered_types_list []TypeSlice
}

func MakeRegistrationList(len, cap int) (r RegistrationList) {
	r.registered_types_list = make([]TypeSlice, len, cap)
	return
}

// returns the number of the RegistrationIDCount 
func (t *RegistrationList) RegistrationIDCount() int { return len(t.registered_types_list) }

func (t *RegistrationList) AddNewRegistrationID() (id int) {
	id = t.RegistrationIDCount()
	t.registered_types_list = append(t.registered_types_list, nil)
	return
}

func (t *RegistrationList) Register(id int, collection_types ...Type) {
	if 0 > id || id >= t.RegistrationIDCount() {
		panic(fmt.Sprintf("id (%d) out of bounds [0,%d)", id, t.RegistrationIDCount()))
	}

	// If there's no types, just bail
	if len(collection_types) == 0 {
		t.registered_types_list[id] = nil
		return
	}

	// Sort for duplicate removal and fast access
	sort.Sort(TypeSlice(collection_types))

	// removes all duplicates
	var wi int
	for i := range collection_types {
		if collection_types[i] == collection_types[wi] {
			continue
		}

		wi++
		collection_types[wi] = collection_types[i]
	}
	collection_types = collection_types[:wi+1]

	// Store the slice for safety checks
	t.registered_types_list[id] = TypeSlice(collection_types)
}

func (t *RegistrationList) Check(id int, collection_type Type) bool {
	i := sort.Search(len(t.registered_types_list[id]), func(i int) bool { return *t.registered_types_list[id][i] >= *collection_type })

	return i < len(t.registered_types_list[id]) && t.registered_types_list[id][i] == collection_type
}
