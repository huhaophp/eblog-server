package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(s string) (pass string) {
	w := md5.New()
	io.WriteString(w, s)
	pass = fmt.Sprintf("%x", w.Sum(nil))
	return
}
