package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func String(value string) *string {
	return &value
}

func Load(s interface{}, path *string) error {
	if s == nil {
		return nil
	}

	if err := hydrateStructFields(s); err != nil {
		return err
	}

	if path != nil {
		root, _ := os.OpenRoot(".")

		defer func() {
			_ = root.Close()
		}()

		b, err := root.ReadFile(*path)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(b, s); err != nil {
			return err
		}
	}

	return nil
}

func hydrateStructFields(s interface{}) error {
	val := reflect.ValueOf(s).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldKind := field.Kind()

		if fieldKind == reflect.Struct {
			if err := hydrateStructFields(field.Addr().Interface()); err != nil {
				return err
			}
		} else if fieldKind == reflect.Slice {
			fmt.Printf("Slice slicey %v\n", field.Len())
			envTag := fieldType.Tag.Get("env")
			argTag := fieldType.Tag.Get("arg")

			value := ""
			if envTag != "" {
				value = os.Getenv(envTag)
			}

			if argTag != "" {
				arg := GetArgument(argTag)

				if arg != nil {
					value = *arg
				}
			}

			if value == "" {
				continue
			}

			values := strings.Split(value, ",")

			indexType := field.Type().Elem().Kind()
			switch indexType {
			case reflect.String:
				field.Set(reflect.ValueOf(values))
			case reflect.Bool:
				var elements []bool
				for _, v := range values {
					b, err := strconv.ParseBool(v)
					if err != nil {
						continue
					}
					elements = append(elements, b)
				}

				field.Set(reflect.ValueOf(elements))
			case reflect.Int:
				var elements []int
				for _, v := range values {
					n, err := strconv.ParseInt(v, 10, 32)
					if err != nil {
						continue
					}
					elements = append(elements, int(n))
				}

				field.Set(reflect.ValueOf(elements))
			case reflect.Int64:
				var elements []int64
				for _, v := range values {
					n, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						continue
					}
					elements = append(elements, n)
				}

				field.Set(reflect.ValueOf(elements))
			case reflect.Float64:
				var elements []float64
				for _, v := range values {
					n, err := strconv.ParseFloat(v, 64)
					if err != nil {
						continue
					}
					elements = append(elements, n)
				}

				field.Set(reflect.ValueOf(elements))
			}
		} else {
			envTag := fieldType.Tag.Get("env")
			argTag := fieldType.Tag.Get("arg")

			value := ""
			if envTag != "" {
				value = os.Getenv(envTag)
			}

			if argTag != "" {
				arg := GetArgument(argTag)

				if arg != nil {
					value = *arg
				}
			}

			if fieldKind == reflect.String {
				field.SetString(value)
				continue
			}

			if value == "" {
				continue
			}

			switch fieldKind {
			case reflect.Bool:
				b, err := strconv.ParseBool(value)
				if err != nil {
					continue
				}
				field.SetBool(b)
			case reflect.Int:
				n, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					continue
				}
				field.SetInt(n)
			case reflect.Int64:
				n, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					continue
				}
				field.SetInt(n)
			case reflect.Float64:
				n, err := strconv.ParseFloat(value, 64)
				if err != nil {
					continue
				}
				field.SetFloat(n)
			}
		}
	}

	return nil
}
