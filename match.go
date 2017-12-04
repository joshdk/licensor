// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package license

import (
	"bytes"
	"regexp"

	"github.com/joshdk/license/spdx"
)

var (
	reWords = regexp.MustCompile(`[\w']+`)
)

func Best(unknown []byte) Match {

	var (
		unknownWords = wordSet(unknown)
		bestMatch    *Match
	)

	for _, license := range spdx.All() {

		licenseWords := wordSet([]byte(license.Text))

		confidence := dice(licenseWords, unknownWords)

		if bestMatch == nil || confidence > bestMatch.Confidence {
			bestMatch = &Match{
				Confidence: confidence,
				License:    license,
			}
		}
	}

	return *bestMatch
}

func dice(reference map[string]struct{}, target map[string]struct{}) float64 {
	var common int

	for w := range target {
		if _, ok := reference[w]; ok {
			common++
		}
	}

	return 2 * float64(common) / float64(len(target)+len(reference))
}

func wordSet(data []byte) map[string]struct{} {

	words := map[string]struct{}{}
	data = bytes.ToLower(data)

	matches := reWords.FindAll(data, -1)

	for _, match := range matches {
		words[string(match)] = struct{}{}
	}

	return words
}
