package config

import (
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetArgumentSingleDashAndEquals(t *testing.T) {
	RegisterTestingT(t)

	os.Args = append(os.Args, "-single-dash")
	os.Args = append(os.Args, "12")
	os.Args = append(os.Args, "--double-dash")
	os.Args = append(os.Args, "34")
	os.Args = append(os.Args, "-single-dash-equals=56")
	os.Args = append(os.Args, "--double-dash-equals=78")
	os.Args = append(os.Args, "-enable-some-setting")
	os.Args = append(os.Args, "-unused-positional-setting")
	os.Args = append(os.Args, "-another-unused-positional-setting")

	value := GetArgument("single-dash")
	Expect(value).To(Not(BeNil()))
	Expect(*value).To(Equal("12"))

	value = GetArgument("double-dash")
	Expect(value).To(Not(BeNil()))
	Expect(*value).To(Equal("34"))

	value = GetArgument("single-dash-equals")
	Expect(value).To(Not(BeNil()))
	Expect(*value).To(Equal("56"))

	value = GetArgument("double-dash-equals")
	Expect(value).To(Not(BeNil()))
	Expect(*value).To(Equal("78"))

	value = GetArgument("enable-some-setting")
	Expect(value).To(Not(BeNil()))
	Expect(*value).To(Equal(""))

	value = GetArgument("missing-setting")
	Expect(value).To(BeNil())
}
