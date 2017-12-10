// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package licensor

import (
	"github.com/joshdk/licensor/spdx"
)

// A Match represents the result of a statistical comparison against a specific
// license.
type Match struct {

	// Confidence is the statistical certainty of the match. Value lies within
	// the interval [0, 1] to where 0 is less confident and 1 is more
	// confident.
	Confidence float64

	// License is the matched SPDX license.
	License spdx.License
}
