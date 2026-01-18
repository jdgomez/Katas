package blogrenderer_test

import (
	"bytes"
	"testing"

	"github.com/approvals/go-approval-tests"
	"github.com/jdgomez/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buff := bytes.Buffer{}
		err := blogrenderer.Render(&buff, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buff.String())

		//		got := buff.String()
		//		want := `<h1>hello world</h1>
		//<p>This is a description</p>
		//Tags: <ul><li>go</li><li>tdd</li></ul>`
		//
		//		if got != want {
		//			t.Errorf("got '%s', want '%s'", got, want)
		//		}
	})
}
