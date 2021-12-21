// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package observeddata

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *ObservedData) Valid() (bool, int, map[string]string) {
	problemsFound := 0
	resultDetails := make(map[string]string)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO()
	problemsFound += pBase
	for key, value := range dBase {
		resultDetails[key] = value
	}

	// Verify First Observed property is present
	if o.FirstObserved == "" {
		problemsFound++
		resultDetails["first_observed"] = "The first observed property is required but missing"
	}

	// Verify Last Observed property is present
	if o.LastObserved == "" {
		problemsFound++
		resultDetails["last_observed"] = "The last observed property is required but missing"
	}

	// Verify Number Observed property is present
	if o.NumberObserved == 0 {
		problemsFound++
		resultDetails["number_observed"] = "The number observed property is required and is missing or set to zero"
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
