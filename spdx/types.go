// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package spdx

type License struct {
	Name             string   `json:"name"`
	Identifier       string   `json:"licenseId"`
	Text             string   `json:"licenseText"`
	URIs             []string `json:"seeAlso"`
	StandardTemplate string   `json:"standardLicenseTemplate"`
	StandardHeader   string   `json:"standardLicenseHeader"`
	Deprecated       bool     `json:"isDeprecatedLicenseId"`
	OSIApproved      bool     `json:"isOsiApproved"`
}
