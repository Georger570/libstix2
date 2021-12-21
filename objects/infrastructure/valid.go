// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *Infrastructure) Valid() (bool, int, map[string]string) {
	problemsFound := 0
	resultDetails := make(map[string]string)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO()
	problemsFound += pBase
	for key, value := range dBase {
		resultDetails[key] = value
	}

	// Verify object Name property is present
	_, pName, dName := o.NameProperty.VerifyExists()
	problemsFound += pName
	for key, value := range dName {
		resultDetails[key] = value
	}

	if len(o.InfrastructureTypes) == 0 {
		problemsFound++
		resultDetails["infrastructure_types"] = "The infrastructure types property is required but missing"
	}

	// TODO add check to make sure values are from vocabulary. Something like
	// helpers.ValidSlice("InfrastructureTypes", o.InfrastructureTypes, vocabs.InfrastructureType)

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
