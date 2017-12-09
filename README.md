[![License](https://img.shields.io/github/license/joshdk/licensor.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/joshdk/licensor?status.svg)](https://godoc.org/github.com/joshdk/licensor)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/licensor)](https://goreportcard.com/report/github.com/joshdk/licensor)
[![CircleCI](https://circleci.com/gh/joshdk/licensor.svg?&style=shield)](https://circleci.com/gh/joshdk/licensor/tree/master)

# Licensor

📝 Detect what license a project is distributed under

## Installing

You can fetch this library by running the following

    go get -u github.com/joshdk/licensor

## Usage

```go
import (
	"fmt"
	"github.com/joshdk/licensor"
)

// Example content from https://github.com/golang/go/blob/master/LICENSE
const unknown = `
	Copyright (c) 2009 The Go Authors. All rights reserved.
	Redistribution and use in source and binary forms, with or without
	...
	(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
	OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

// Find license that is the closest match 
match := licensor.Best([]byte(unknown))

fmt.Printf("License name:     %s\n",   match.License.Name)
fmt.Printf("SPDX identifier:  %s\n",   match.License.Identifier)
fmt.Printf("Match confidence: %.2f\n", match.Confidence)
// License name:     BSD 3-clause "New" or "Revised" License
// SPDX identifier:  BSD-3-Clause
// Match confidence: 0.96
```

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.
