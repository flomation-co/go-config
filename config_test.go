package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/gomega"
)

type testStruct struct {
	Name    string `json:"name" env:"CONFIG_NAME" arg:"config-name"`
	Details struct {
		SomeString  string  `json:"some-string" env:"SOME_STRING" arg:"some-string"`
		SomeNumber  int64   `json:"some-number" env:"SOME_NUMBER" arg:"some-number"`
		SomeBoolean bool    `json:"some-boolean" env:"SOME_BOOLEAN" arg:"some-boolean"`
		SomeFloat   float64 `env:"SOME_FLOAT" arg:"some-float"`
	}
	EnvironmentSetValue string  `env:"ENVIRON_VALUE" arg:"environ-value"`
	ArgumentSetValue    string  `env:"ARGUMENT_VALUE" arg:"arg-value"`
	Array               []int64 `json:"array" env:"CONFIG_ARRAY" arg:"config-array"`
	Enabled             bool    `json:"enabled"`
	Amount              float64 `json:"amount"`
}

func TestNil(t *testing.T) {
	RegisterTestingT(t)

	err := Load(nil, nil)
	Expect(err).To(BeNil())
}

func TestLoadFile(t *testing.T) {
	RegisterTestingT(t)

	var cfg testStruct
	err := Load(&cfg, String("test/good.json"))

	Expect(err).To(BeNil())
	Expect(cfg.Name).To(Equal("Good Test Config"))
	Expect(cfg.Details).To(Not(BeNil()))
	Expect(cfg.Details.SomeString).To(Equal("This is a test string"))
	Expect(cfg.Details.SomeNumber).To(Equal(int64(1234)))
	Expect(cfg.Details.SomeBoolean).To(Equal(true))
	Expect(cfg.Enabled).To(BeTrue())
	Expect(cfg.Amount).To(Equal(32.6))
	Expect(len(cfg.Array)).To(Equal(5))
}

func TestLoadBadFile(t *testing.T) {
	RegisterTestingT(t)

	var cfg testStruct
	err := Load(&cfg, String("test/bad.json"))

	Expect(err).To(Not(BeNil()))
}

func TestLoadMissingFile(t *testing.T) {
	RegisterTestingT(t)

	var cfg testStruct
	err := Load(&cfg, String("test/does-not-exist.json"))

	Expect(err).To(Not(BeNil()))
}

func TestLoadFileAndEnvironmentValue(t *testing.T) {
	RegisterTestingT(t)

	randomisedValue := uuid.NewString()
	err := os.Setenv("ENVIRON_VALUE", randomisedValue)
	Expect(err).To(BeNil())

	err = os.Setenv("SOME_FLOAT", "123.45")
	Expect(err).To(BeNil())

	var cfg testStruct
	err = Load(&cfg, String("test/good.json"))

	Expect(err).To(BeNil())
	Expect(cfg.EnvironmentSetValue).To(Equal(randomisedValue))
	Expect(cfg.Details.SomeFloat).To(Equal(123.45))
}

func TestEnvironmentOnly(t *testing.T) {
	RegisterTestingT(t)

	randomisedValue := uuid.NewString()
	err := os.Setenv("ENVIRON_VALUE", randomisedValue)
	Expect(err).To(BeNil())

	var cfg testStruct
	err = Load(&cfg, nil)

	Expect(err).To(BeNil())
	Expect(cfg.EnvironmentSetValue).To(Equal(randomisedValue))
}

func TestArgumentOnly(t *testing.T) {
	RegisterTestingT(t)

	os.Args = append(os.Args, fmt.Sprintf("-arg-value=%v", "an argument value"))

	var cfg testStruct
	err := Load(&cfg, nil)

	Expect(err).To(BeNil())
	Expect(cfg.ArgumentSetValue).To(Equal("an argument value"))
}

func TestSetBoolean(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "true")
	Expect(err).To(BeNil())

	cfg := struct {
		Value bool `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(cfg.Value).To(BeTrue())
}

func TestSetBadBoolean(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "not-a-boolean")
	Expect(err).To(BeNil())

	cfg := struct {
		Value bool `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(cfg.Value).To(BeFalse())
}

func TestSetInt64(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "123")
	Expect(err).To(BeNil())

	cfg := struct {
		Value int64 `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(cfg.Value).To(Equal(int64(123)))
}

func TestSetBadInt64(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "not-a-number64")
	Expect(err).To(BeNil())

	cfg := struct {
		Value int64 `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(cfg.Value).To(Equal(int64(0)))
}

func TestSetInt(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "123")
	Expect(err).To(BeNil())

	cfg := struct {
		Value int `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(cfg.Value).To(Equal(123))
}

func TestSetBadInt(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "not-a-number")
	Expect(err).To(BeNil())

	cfg := struct {
		Value int `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(cfg.Value).To(Equal(0))
}

func TestSetFloat(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "123.45")
	Expect(err).To(BeNil())

	cfg := struct {
		Value float64 `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(cfg.Value).To(Equal(123.45))
}

func TestSetBadFloat(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "not-a-float")
	Expect(err).To(BeNil())

	cfg := struct {
		Value float64 `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(cfg.Value).To(Equal(float64(0)))
}

func TestSetSlice(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "1,2,3,4")
	Expect(err).To(BeNil())

	cfg := struct {
		Value []string `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(len(cfg.Value)).To(Equal(4))
	Expect(cfg.Value[0]).To(Equal("1"))
}

func TestSetSliceInt64(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "1,2,3,4")
	Expect(err).To(BeNil())

	cfg := struct {
		Value []int64 `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(len(cfg.Value)).To(Equal(4))
	Expect(cfg.Value[0]).To(Equal(int64(1)))
}

func TestSetSliceInt(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "1,2,3,4")
	Expect(err).To(BeNil())

	cfg := struct {
		Value []int `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(len(cfg.Value)).To(Equal(4))
	Expect(cfg.Value[0]).To(Equal(1))
}

func TestSetSliceFloat(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "1,2,3,4")
	Expect(err).To(BeNil())

	cfg := struct {
		Value []float64 `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(len(cfg.Value)).To(Equal(4))
	Expect(cfg.Value[0]).To(Equal(float64(1)))
}

func TestSetSliceBool(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "true,false,true,false")
	Expect(err).To(BeNil())

	cfg := struct {
		Value []bool `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(len(cfg.Value)).To(Equal(4))
	Expect(cfg.Value[0]).To(BeTrue())
	Expect(cfg.Value[1]).To(BeFalse())
	Expect(cfg.Value[2]).To(BeTrue())
	Expect(cfg.Value[3]).To(BeFalse())
}

func TestSetSliceBadInt64(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "x,y,z,w")
	Expect(err).To(BeNil())

	cfg := struct {
		Value []int64 `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(len(cfg.Value)).To(Equal(0))
}

func TestSetSliceBadInt(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "x,y,z,w")
	Expect(err).To(BeNil())

	cfg := struct {
		Value []int `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(len(cfg.Value)).To(Equal(0))
}

func TestSetSliceBadFloat(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "x,y,z,w")
	Expect(err).To(BeNil())

	cfg := struct {
		Value []float64 `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(len(cfg.Value)).To(Equal(0))
}

func TestSetSliceBadBool(t *testing.T) {
	RegisterTestingT(t)

	err := os.Setenv("ENVIRON_VALUE", "x,true,z,w")
	Expect(err).To(BeNil())

	cfg := struct {
		Value []bool `env:"ENVIRON_VALUE"`
	}{}

	err = Load(&cfg, nil)
	Expect(err).To(BeNil())
	Expect(len(cfg.Value)).To(Equal(1))
}
