// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateLocalPeeringConnectionDetails The representation of UpdateLocalPeeringConnectionDetails
type UpdateLocalPeeringConnectionDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m UpdateLocalPeeringConnectionDetails) String() string {
	return common.PointerString(m)
}