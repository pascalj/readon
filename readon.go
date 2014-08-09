package readon

import (
	"code.google.com/p/go-html-transform/h5"
	"code.google.com/p/go-html-transform/html/transform"
	"io"
)

func Readon(reader io.Reader) (io.Reader, error) {
	tree, _ := h5.New(reader)
	t := transform.New(tree)
	removeScripts(t)
	return reader, nil
}

func removeScripts(t *transform.Transformer) {
	t.Apply(transform.Replace(), "script")
}
