package hgconfig_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGoHgconfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "go-hgconfig Suite")
}
