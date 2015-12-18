package hgconfig_test

import (
	"github.com/nishanths/go-hgconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"path/filepath"
)

func useConfig(s string) func() {
	tmp, err := ioutil.TempDir("", "go-hgconfig-test")
	if err != nil {
		panic(err)
	}

	hgrc := filepath.Join(tmp, ".hgrc")
	err = ioutil.WriteFile(hgrc, []byte(s), 0700)
	if err != nil {
		panic(err)
	}

	prevHome := os.Getenv("HOME")
	os.Setenv("HOME", tmp)

	return func() {
		os.Setenv("HOME", prevHome)
	}
}

var _ = Describe("GoHgconfig", func() {
	var undo func()

	BeforeEach(func() {
		undo = useConfig(`
[ui]
username = Alice Wonderland <alice@mit.edu>

[extensions]
color =
# progress =

[color]
mode = auto`)
	})

	Describe("Username", func() {
		Context("Username exists in config", func() {
			It("should return the username with nil error", func() {
				username, err := hgconfig.Username()
				Expect(err).NotTo(HaveOccurred())
				Expect(username).To(Equal("Alice Wonderland <alice@mit.edu>"))
			})
		})
	})

	Describe("Get", func() {
		Context("Getting existing name", func() {
			It("should return the value with nil error", func() {
				value, err := hgconfig.Get("color.mode")
				Expect(err).NotTo(HaveOccurred())
				Expect(value).To(Equal("auto"))
			})
		})

		Context("Getting existing name but with empty value", func() {
			It("should return the value with nil error", func() {
				value, err := hgconfig.Get("extensions.color")
				Expect(err).NotTo(HaveOccurred())
				Expect(value).To(Equal(""))
			})
		})

		Context("Name does not exist", func() {
			It("should appropriate error", func() {
				_, err := hgconfig.Get("extensions.progress")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("extensions.progress"))
			})
		})
	})

	AfterEach(func() {
		undo()
	})
})
