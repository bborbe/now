// Copyright (c) 2018 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package now_test

import (
	"bytes"
	"testing"

	"github.com/bborbe/now"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Now", func() {
	It("Contains more than zero locations", func() {
		Expect(len(now.DefaultLocations)).To(BeNumerically(">", 0))
	})
	It("Writes without error", func() {
		buf := &bytes.Buffer{}
		err := now.DefaultLocations.Write(buf)
		Expect(err).To(BeNil())
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Now Suite")
}
