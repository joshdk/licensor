// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package license

import (
	"github.com/joshdk/license/spdx"
)

type Match struct {
	Confidence float64
	License    spdx.License
}
