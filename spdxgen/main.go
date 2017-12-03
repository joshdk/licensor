// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/joshdk/license/spdx"
)

func find(root string) chan string {
	jsonDir := path.Join(root, "json")

	detailsDir := path.Join(jsonDir, "details")

	out := make(chan string)

	go func() {
		defer close(out)

		err := filepath.Walk(detailsDir, func(fullPath string, info os.FileInfo, err error) error {
			if path.Ext(fullPath) == ".json" {
				out <- fullPath
			}

			return nil
		})

		if err != nil {
			panic(err)
		}
	}()

	return out
}

func read(in chan string) chan []byte {
	out := make(chan []byte)

	go func() {
		defer close(out)

		for filename := range in {
			contents, err := ioutil.ReadFile(filename)
			if err != nil {
				panic(err)
			}

			out <- contents
		}
	}()

	return out
}

func parse(in chan []byte) chan spdx.License {
	out := make(chan spdx.License)

	go func() {
		defer close(out)

		for body := range in {
			var license spdx.License

			if err := json.Unmarshal(body, &license); err != nil {
				panic(err)

			}

			var uris []string

			for _, uri := range license.URIs {
				uri = strings.TrimSpace(uri)
				if uri != "" {
					uris = append(uris, uri)
				}
			}

			license.URIs = uris

			out <- license
		}
	}()

	return out
}

func pkgDir(pkg string) (string, error) {
	if pkg == "." {
		return os.Getwd()
	}

	gopath, ok := os.LookupEnv("GOPATH")
	if !ok {
		return "", errors.New("no GOPATH set in working environment")
	}

	return filepath.Join(gopath, "src", pkg), nil
}

func generate(spdxDataPkg string) error {

	// Resolve SPDX license list data package to the correct directory
	spdxDataPkgDir, err := pkgDir(spdxDataPkg)
	if err != nil {
		return err
	}

	licenses := parse(read(find(spdxDataPkgDir)))

	for license := range licenses {
		fmt.Printf("Found %s (%s)\n", license.Name, license.Identifier)
	}

	return nil
}

func main() {

	err := func() error {
		if len(os.Args) < 2 {
			return errors.New("insufficient arguments")
		}

		var (
			spdxDataPkg = os.Args[1]
		)

		return generate(spdxDataPkg)
	}()

	if err != nil {
		fmt.Fprintf(os.Stderr, "spdxgen: %s", err)
		os.Exit(1)
	}
}
