// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package licensor

import (
	"github.com/joshdk/licensor/spdx"
)

type Match struct {
	Confidence float64
	License    spdx.License
}
