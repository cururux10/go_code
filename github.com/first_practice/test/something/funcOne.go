package something

import "strings"

func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
