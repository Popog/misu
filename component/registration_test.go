// Copyright 2011 The Misu Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package component

import "testing"

func nullcFactory() Component { return nil }
func nulldFactory() Data      { return nil }

func TestRegisterComponentType(t *testing.T) {
	defer resetRegistration(45)
	for i := 0; i < 1000 && !t.Failed(); i++ {
		var testAType Type = RegisterComponentType("testA", nullcFactory, nulldFactory)
		if IsValidType(testAType) {
			t.Fatal("Check returned type was good before FinalizeRegistration was called")
		}
		FinalizeRegistration()
		if !IsValidType(testAType) {
			t.Fatal("Check returned type was bad after FinalizeRegistration was called")
		}
		resetRegistration(45*i)
	}
}

func TestCIT1(t *testing.T) {
	defer resetRegistration(45)
	for i := 0; i < 1000 && !t.Failed(); i++ {
		var testAType Type = RegisterComponentType("testA", nullcFactory, nulldFactory)
		var testBType Type = RegisterComponentType("testB", nullcFactory, nulldFactory)
		var testCType Type = RegisterComponentType("testC", nullcFactory, nulldFactory)
		var testDType Type = RegisterComponentType("testD", nullcFactory, nulldFactory)

		// This Component needs no types, but
		RegisterInitializeTypes(testAType)
		RegisterDestroyTypes(testAType)

		RegisterInitializeTypes(testBType, testAType, testCType)
		RegisterDestroyTypes(testBType)

		RegisterInitializeTypes(testCType, testDType, testCType)
		RegisterDestroyTypes(testCType)

		RegisterInitializeTypes(testDType, testCType, testCType, testCType, testCType)
		RegisterDestroyTypes(testDType)

		FinalizeRegistration()
		
		if CheckInitializeType(testAType, testAType) {
			t.Error("CheckInitializeType(testAType, testAType) gave incorrect result")
		}
		if CheckInitializeType(testAType, testBType) {
			t.Error("CheckInitializeType(testAType, testBType) gave incorrect result")
		}
		if CheckInitializeType(testAType, testCType) {
			t.Error("CheckInitializeType(testAType, testCType) gave incorrect result")
		}
		if CheckInitializeType(testAType, testDType) {
			t.Error("CheckInitializeType(testAType, testDType) gave incorrect result")
		}

		if !CheckInitializeType(testBType, testAType) {
			t.Error("CheckInitializeType(testBType, testAType) gave incorrect result")
		}
		if CheckInitializeType(testBType, testBType) {
			t.Error("CheckInitializeType(testBType, testBType) gave incorrect result")
		}
		if !CheckInitializeType(testBType, testCType) {
			t.Error("CheckInitializeType(testBType, testCType) gave incorrect result")
		}
		if CheckInitializeType(testBType, testDType) {
			t.Error("CheckInitializeType(testBType, testDType) gave incorrect result")
		}

		if CheckInitializeType(testCType, testAType) {
			t.Error("CheckInitializeType(testCType, testAType) gave incorrect result")
		}
		if CheckInitializeType(testCType, testBType) {
			t.Error("CheckInitializeType(testCType, testBType) gave incorrect result")
		}
		if !CheckInitializeType(testCType, testCType) {
			t.Error("CheckInitializeType(testCType, testCType) gave incorrect result")
		}
		if !CheckInitializeType(testCType, testDType) {
			t.Error("CheckInitializeType(testCType, testDType) gave incorrect result")
		}

		if CheckInitializeType(testDType, testAType) {
			t.Error("CheckInitializeType(testDType, testAType) gave incorrect result")
		}
		if CheckInitializeType(testDType, testBType) {
			t.Error("CheckInitializeType(testDType, testBType) gave incorrect result")
		}
		if !CheckInitializeType(testDType, testCType) {
			t.Error("CheckInitializeType(testDType, testCType) gave incorrect result")
		}
		if CheckInitializeType(testDType, testDType) {
			t.Error("CheckInitializeType(testDType, testDType) gave incorrect result")
		}
		
		resetRegistration(45*i)
	}
}

func TestBenchmarkCIT1(t *testing.T) {
	defer resetRegistration(45)
	for i := 0; i < 1000 && !t.Failed(); i++ {
		var testAType Type = RegisterComponentType("testA", nullcFactory, nulldFactory)
		var testBType Type = RegisterComponentType("testB", nullcFactory, nulldFactory)
		var testCType Type = RegisterComponentType("testC", nullcFactory, nulldFactory)
		var testDType Type = RegisterComponentType("testD", nullcFactory, nulldFactory)

		// This Component needs no types, but
		RegisterInitializeTypes(testAType)
		RegisterDestroyTypes(testAType)

		RegisterInitializeTypes(testBType, testAType, testCType)
		RegisterDestroyTypes(testBType)

		RegisterInitializeTypes(testCType, testDType, testCType)
		RegisterDestroyTypes(testCType)

		RegisterInitializeTypes(testDType, testCType, testCType, testCType, testCType)
		RegisterDestroyTypes(testDType)

		FinalizeRegistration()
		
		resetRegistration(45*i)
	}
}
