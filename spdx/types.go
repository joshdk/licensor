// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package spdx

import (
	"github.com/dave/jennifer/jen"
)

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

func (l *License) JenValue() *jen.Statement {
	return jen.Line().Values(jen.Dict{
		jen.Id("Name"):             jen.Lit(l.Name),
		jen.Id("Identifier"):       jen.Lit(l.Identifier),
		jen.Id("Text"):             jen.Lit(l.Text),
		jen.Id("URIs"):             jen.Lit(l.URIs),
		jen.Id("StandardTemplate"): jen.Lit(l.StandardTemplate),
		jen.Id("StandardHeader"):   jen.Lit(l.StandardHeader),
		jen.Id("Deprecated"):       jen.Lit(l.Deprecated),
		jen.Id("OSIApproved"):      jen.Lit(l.OSIApproved),
	})
}
