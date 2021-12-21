// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
ValidSDO - This method will verify and test all of the properties on a STIX
Domain Object to make sure they are valid per the specification. It will return
a boolean, an integer that tracks the number of problems found, and a slice of
strings that contain the detailed results, whether good or bad.
*/
func (o *CommonObjectProperties) ValidSDO() (bool, int, map[string]string) {
	valid := true
	problemsFound := 0
	resultDetails := make(map[string]string)

	// Verify object Type property is present
	if vType, pType, dType := o.TypeProperty.VerifyExists(); vType {
		problemsFound += pType
		for key, value := range dType {
			resultDetails[key] = value
		}
	}

	// Verify Spec Version property is present
	if vSpecVersion, pSpecVersion, dSpecVersion := o.SpecVersionProperty.VerifyExists(); vSpecVersion {
		problemsFound += pSpecVersion
		for key, value := range dSpecVersion {
			resultDetails[key] = value
		}
	}

	// Verify object ID property is present
	if vID, pID, dID := o.IDProperty.VerifyExists(); vID {
		problemsFound += pID
		for key, value := range dID {
			resultDetails[key] = value
		}
	}

	// Verify object Created property is present
	if vCreated, pCreated, dCreated := o.CreatedProperty.VerifyExists(); vCreated {
		problemsFound += pCreated
		for key, value := range dCreated {
			resultDetails[key] = value
		}
	}
	// Verify object Created property is present
	if vModified, pModified, dModified := o.ModifiedProperty.VerifyExists(); vModified {
		problemsFound += pModified
		for key, value := range dModified {
			resultDetails[key] = value
		}
	}

	if problemsFound > 0 {
		valid = false
	}

	return valid, problemsFound, resultDetails
}
