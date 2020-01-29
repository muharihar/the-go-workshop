package strings

import (
	_ "unicode/utf8"
	_ "unsafe"
)

type Builder struct { // of receiver, to detect copies by value
	addr *Builder
	buf  []byte
}
