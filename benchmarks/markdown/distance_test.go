package distance

import (
	"bytes"
	_ "embed"
	"testing"

	"gitlab.com/golang-commonmark/markdown"
)

//go:embed go_readme.md
var src []byte

func BenchmarkMarkdownRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md := markdown.New(
			markdown.XHTMLOutput(true),
			markdown.Typographer(true),
			markdown.Linkify(true),
			markdown.Tables(true),
		)

		var buf bytes.Buffer
		if err := md.Render(&buf, src); err != nil {
			panic(err)
		}
	}
}
