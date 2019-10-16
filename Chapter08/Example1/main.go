package strings
import (
	"unicode/utf8"
	"unsafe"
)
type Builder struct {
	addr *Builder // of receiver, to detect copies by value
	buf  []byte
}
// https://golang.org/src/strings/compare.go
package strings
func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
// https://golang.org/src/strings/replace.go
package strings
import (
	"io"
	"sync"
)
type Replacer struct {
	once   sync.Once // guards buildOnce method
	r      replacer
	oldnew []string
}
