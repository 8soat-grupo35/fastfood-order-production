package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFastfoodOrderProduction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FastfoodOrderProduction Suite")
}
