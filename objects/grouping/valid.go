// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package grouping

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *Grouping) Valid() (bool, int, map[string]string) {
	problemsFound := 0
	resultDetails := make(map[string]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO()
	problemsFound += pBase
	for key, value := range dBase {
		resultDetails[key] = value
	}

	if o.Context == "" {
		problemsFound++
		resultDetails["context"] = "The context property is required but missing"
	}

	// Verify object refs property is present
	_, pObjectRefs, dObjectRefs := o.ObjectRefsProperty.VerifyExists()
	problemsFound += pObjectRefs
	for key, value := range dObjectRefs {
		resultDetails[key] = value
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
