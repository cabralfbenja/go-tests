package iteration

import "strings"

func Repeat(c string, count int) string {
    var repeated strings.Builder
	for i := 0; i < count; i++ {
        repeated.WriteString(c)
	}
	return repeated.String()
}
