package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(s string) string {
	w := md5.New()
	io.WriteString(w, s)
	return fmt.Sprintf("%x", w.Sum(nil))
}
