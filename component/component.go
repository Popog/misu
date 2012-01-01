// Copyright 2011 The Misu Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package component

import "sort"
import "fmt"
import "encoding/gob"

type Type *int

// The interface which components have to meet to be components
type Component interface {
	LoadData(data Data)
	Initialize(parent Collection)
	Initialized() bool
	Destroy(parent Collection)
	Destroyed() bool
	GetType() Type
}

// Data types should be POD that can be serialized via GOB
type Data interface {
	GetType() Type
}

type factory_data struct {
	component_factory func() Component
	data Data
}

const reservation_size = 16

var registration_finalized bool

var initial_registration_number int = 45
var max_registration_number int = initial_registration_number

var registered_initialize_types, registered_destroy_types RegistrationList = MakeRegistrationList(0, reservation_size), MakeRegistrationList(0, reservation_size)

var name_to_registered_type_map map[string]Type = make(map[string]Type, reservation_size)
var registered_type_name_list []string = make([]string, 0, reservation_size)

var temp_component_factories map[string]factory_data = make(map[string]factory_data, reservation_size)
var component_factories []factory_data

var registration_callbacks []func() = make([]func(), 0, reservation_size)

func resetRegistration(i_reg_number int) {
	registration_finalized = false

	initial_registration_number = i_reg_number
	max_registration_number = initial_registration_number

	registered_initialize_types, registered_destroy_types = MakeRegistrationList(0, reservation_size), MakeRegistrationList(0, reservation_size)

	name_to_registered_type_map = make(map[string]Type, reservation_size)
	registered_type_name_list = make([]string, 0, reservation_size)

	temp_component_factories = make(map[string]factory_data, reservation_size)
	component_factories = nil

	registration_callbacks = make([]func(), 0, reservation_size)
}

func RegistrationFinalized() bool { return registration_finalized }
func FinalizeRegistration() {
	// Sort the names to guarantee id's regardless of order of calls
	sort.Strings(registered_type_name_list)

	// Reserve the space we need
	component_factories = make([]factory_data, len(registered_type_name_list))
	max_registration_number = initial_registration_number + len(registered_type_name_list)

	// process all the (sorte) names to generate the ids 
	for i, v := range registered_type_name_list {
		// Give the type a real value
		*name_to_registered_type_map[v] = initial_registration_number + i
		// Store the factory_data where it should be kept
		component_factories[i] = temp_component_factories[v]
	}

	for _, f := range registration_callbacks {
		f()
	}

	registration_finalized = true
}

// Registers a new type of component with a factory to make it, it is expected to be called during initialization
// Name must be unique, component_data must be 
func RegisterComponentType(name string, component_factory func() Component, data Data) (registered_type Type) {
	// Not to be called after finalizing registration
	if registration_finalized {
		panic("component.RegisterComponentType called after component.FinalizeRegistration")
	}
	if component_factory == nil {
		panic("component_factory is nil")
	}
	if data == nil {
		panic("data is nil")
	}
	// See if this was already registered
	if _, ok := name_to_registered_type_map[name]; ok {
		panic(name + " was already Registered.")
	}

	// Return a new Type, but we can't give it a real value until FinalizeRegistration is called
	registered_type = Type(new(int))
	*registered_type = -1

	// Store the name, type, and factory_data
	name_to_registered_type_map[name] = registered_type
	registered_type_name_list = append(registered_type_name_list, name)
	temp_component_factories[name] = factory_data{component_factory: component_factory, data: data}

	// Reserve some space
	registered_initialize_types.AddNewRegistrationID()
	registered_destroy_types.AddNewRegistrationID()

	gob.RegisterName(name, data.(interface{}))
	return
}

// init_types enumerates all the types that will be available in Component.Initialize
func RegisterInitializeTypes(registered_type Type, collection_types ...Type) {
	f := func() {
		// Error check registered_type
		if !IsValidType(registered_type) {
			panic(fmt.Sprintf("registered_type (%d) is invalid", *registered_type))
		}

		// Error check collection_types
		for i := range collection_types {
			if !IsValidType(collection_types[i]) {
				panic("collection_types contains invalid type")
			}
		}
		// Store the slice for safety checks
		registered_initialize_types.Register(int(*registered_type)-initial_registration_number, collection_types...)
	}
	registration_callbacks = append(registration_callbacks, f)
}

// init_types enumerates all the types that will be available in Component.Destroy
func RegisterDestroyTypes(registered_type Type, collection_types ...Type) {
	f := func() {
		// Error check registered_type
		if !IsValidType(registered_type) {
			panic(fmt.Sprintf("registered_type (%d) is invalid", *registered_type))
		}
		// Error check collection_types
		for i := range collection_types {
			if !IsValidType(collection_types[i]) {
				panic("collection_types contains invalid type")
			}
		}

		// Store the slice for safety checks
		registered_destroy_types.Register(int(*registered_type)-initial_registration_number, collection_types...)
	}
	registration_callbacks = append(registration_callbacks, f)
}

func CheckInitializeType(registered_type, type_to_check Type) bool {
	// Error check registered_type
	if !IsValidType(registered_type) {
		panic("registered_type is invalid")
	}
	// TODO: error check type_to_check
	return registered_initialize_types.Check(int(*registered_type)-initial_registration_number, type_to_check)
}

func CheckDestroyType(registered_type, type_to_check Type) bool {
	// Error check registered_type
	if !IsValidType(registered_type) {
		panic("registered_type is invalid")
	}
	// Error check type_to_check
	if !IsValidType(type_to_check) {
		panic("type_to_check is invalid")
	}
	return registered_initialize_types.Check(int(*registered_type)-initial_registration_number, type_to_check)
}

func IsValidType(t Type) bool {
	return initial_registration_number <= *t && *t < max_registration_number
}

func CreateData(registered_type Type) Data {
	// Error check registered_type
	if !IsValidType(registered_type) {
		panic("registered_type is invalid")
	}
	
	return component_factories[int(*registered_type)-initial_registration_number].data
}
