package fileutils

import "os"

/*
0644 - Leading 0 specifies that we are using octal notation.

Octal notation is a base-8 number system that uses digits 0-7 to represent numerical values.
It acts as a compact, human-readable shorthand for binary (base-2) code,
with each octal digit representing exactly three binary bits.

Owner has read/write (6),
Group has read-only (4),
Others have read-only (4).
*/

const fileMode = os.FileMode(0644)

func WriteTextFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), fileMode)
}
