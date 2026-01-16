# go-config Configuration Library

go-config allows for using Struct based configuration, with fields automatically populated via a number of mechanisms
(listed below in order of precedence):

- Configuration File (JSON)
- Command Line Arguments
- Environment Variables

## Development

This package is under constant development by the team at Flomation (https://www.flomation.co) and we welcome collaboration
with other developers, please feel free to submit a Pull Request with changes.

## Installation

Within your Golang project

`get get github.com/flomation-co/go-config`

## Usage

Usage is via Field Tags within Structs

```golang
type Configuration struct {
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
```

The `env` tag will pull the value from an environment variable with the same name, likewise the `arg` tag will result in the library
looking for a Command Line argument of the same name.

Command Line arguments are supported in the following formats

```bash
-name 1234
--name 1234
-name=1234
--name=1234
-name
--name 
```

When a field is a slice type, it will be assumed that the value of the provided Environment or Command Line Argument is comma-delimited

To use, simply define your structure and call `Load()`

```golang
if err := Load(&configuration, String("path-to-configuration-file.json")); err != nil {
	panic(err)
}
```