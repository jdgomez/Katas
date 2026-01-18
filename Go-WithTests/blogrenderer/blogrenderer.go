package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, post Post) error {
	blogTemplate, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := blogTemplate.Execute(w, post); err != nil {
		return err
	}

	return nil
}

//func Render(w io.Writer, p Post) error {
//	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
//	if err != nil {
//		return err
//	}
//
//	_, err = fmt.Fprintf(w, "\n<p>%s</p>", p.Description)
//	if err != nil {
//		return err
//	}
//
//	_, err = fmt.Fprintf(w, "\nTags: <ul>")
//	if err != nil {
//		return err
//	}
//
//	for _, tag := range p.Tags {
//		_, err = fmt.Fprintf(w, "<li>%s</li>", tag)
//		if err != nil {
//			return err
//		}
//	}
//	_, err = fmt.Fprintf(w, "</ul>")
//	if err != nil {
//		return err
//	}
//	return err
//}
