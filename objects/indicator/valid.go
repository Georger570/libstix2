// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"time"

	"github.com/RegularITCat/libstix2/timestamp"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *Indicator) Valid() (bool, int, map[string]string) {
	result := true
	problemsFound := 0
	resultDetails := make(map[string]string)

	// Check common base properties first
	if vBase, pBase, dBase := o.CommonObjectProperties.ValidSDO(); vBase {
		problemsFound += pBase
		for key, value := range dBase {
			resultDetails[key] = value
		}
	}

	if len(o.IndicatorTypes) == 0 {
		problemsFound++
		resultDetails["indicator_types"] = "The indicator types property is required but missing"
	}

	if o.Pattern == "" {
		problemsFound++
		resultDetails["pattern"] = "The pattern property is required but missing"
	}
	// TODO, check value to see if it comes from open vocabulary
	if o.PatternType == "" {
		problemsFound++
		resultDetails["pattern_type"] = "The pattern type property is required but missing"
	}

	if o.ValidFrom == "" {
		problemsFound++
		resultDetails["valid_from"] = "The valid from property is required but missing"

	} else if valid := timestamp.Valid(o.ValidFrom); !valid {
		problemsFound++
		resultDetails["valid_from"] = "The valid from property does not contain a valid STIX timestamp"
	}

	if o.ValidUntil != "" && !timestamp.Valid(o.ValidUntil) {
		problemsFound++
		resultDetails["valid_until"] = "The valid until property does not contain a valid STIX timestamp"

	} else {
		validFrom, _ := time.Parse(time.RFC3339, o.ValidFrom)
		validUntil, _ := time.Parse(time.RFC3339, o.ValidUntil)
		if yes := validUntil.After(validFrom); !yes {
			problemsFound++
			resultDetails["valid_until"] = "The valid until timestamp is not later than the valid from timestamp"
		}
	}

	if problemsFound > 0 {
		result = false
	}

	return result, problemsFound, resultDetails
}
