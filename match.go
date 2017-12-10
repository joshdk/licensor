// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package licensor

import (
	"bytes"
	"regexp"
	"sync"

	"github.com/joshdk/licensor/spdx"
)

var (
	reWords = regexp.MustCompile(`[\w']+`)
)

// Best returns the license that is the closes match to the given text.
func Best(unknown []byte) Match {

	var (
		unknownWords = wordSet(unknown)
		matches      = make(chan Match)
		wg           sync.WaitGroup
	)

	for _, license := range spdx.All() {

		wg.Add(1)
		go func(license spdx.License) {
			defer wg.Done()

			licenseWords := wordSet([]byte(license.Text))

			confidence := dice(licenseWords, unknownWords)

			matches <- Match{
				Confidence: confidence,
				License:    license,
			}

		}(license)
	}

	go func() {
		wg.Wait()
		close(matches)
	}()

	best := <-matches

	for match := range matches {
		if match.Confidence > best.Confidence {
			best = match
		}
	}

	return best
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
