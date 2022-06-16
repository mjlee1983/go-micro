package file

import (
	"strings"

	"github.com/mjlee1983/go-micro/go-micro/v2/config/encoder"
)

func format(p string, e encoder.Encoder) string {
	parts := strings.Split(p, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return e.String()
}
