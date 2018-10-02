package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJparser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jparser Suite")
}
