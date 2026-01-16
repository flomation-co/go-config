package config

import (
	"os"
	"strings"
)

func GetArgument(name string) *string {
	args := os.Args

	for n, a := range args {
		if a == "-"+name || a == "--"+name {
			if n+1 < len(args) {
				v := args[n+1]
				if strings.HasPrefix(v, "-") {
					v = ""
					return &v
				}

				return &v
			}
		} else if strings.HasPrefix(a, "-"+name+"=") {
			v := strings.TrimPrefix(a, "-"+name+"=")
			return &v
		} else if strings.HasPrefix(a, "--"+name+"=") {
			v := strings.TrimPrefix(a, "--"+name+"=")
			return &v
		}
	}

	return nil
}
