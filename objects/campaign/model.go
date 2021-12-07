// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package campaign

import (
	"github.com/RegularITCat/libstix2/objects"
	"github.com/RegularITCat/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
Campaign - This type implements the STIX 2 Campaign SDO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties.
*/
type Campaign struct {
	objects.CommonObjectProperties `bson:",inline"`
	properties.NameProperty        `bson:",inline"`
	properties.DescriptionProperty `bson:",inline"`
	properties.AliasesProperty     `bson:",inline"`
	properties.SeenProperties      `bson:",inline"`
	Objective                      string `json:"objective,omitempty" bson:"objective,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Campaign) GetPropertyList() []string {
	return []string{"name", "description", "aliases", "first_seen", "last_seen", "objective"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Campaign object and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Campaign {
	var obj Campaign
	obj.InitSDO("campaign")
	return &obj
}
