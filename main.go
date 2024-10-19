package main

import (
	"log"
	"net/http"

	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func homePage() g.Node {
	return c.HTML5(c.HTML5Props{
		Title: "GAS Stack",
		Head: []g.Node{
			h.Script(
				h.Defer(),
				h.Src("/static/js/cdn.min.js"),
				h.Integrity("sha512-YrobzcejccKh0S5g/usPkVOXAVMKY3visIRbpuDIIvTIk99T/3Fv8E88m4mywdwXvd3YmkfLj54wiNT4XpG89w=="),
				h.CrossOrigin("anonymous"),
			),
			/* h.Link(
				h.Rel("stylesheet"),
				h.Integrity(""),
				h.Href(""),
			), */
		},
		Body: []g.Node{
			g.Attr("x-data", "counter"),
			h.Div(
				h.P(g.Text("Hi, press to increase the hi counter: ")),
				h.Button(
					g.Attr("@click", "increment"),
					h.P(g.Text("Press me!")),
				),
			),
			h.P(
				g.Attr("x-text", "count"),
			),
			h.Script(
				h.Src("/static/js/script.js"),
				h.Integrity("sha512-VpmKNfngVLt4SznCeokFobSiw67lNiNJVecQ9IJgjcFH9FllWSTQNoiIvKoVZSRPQkH9+HjPOujxg3vvAK+GFA=="),
				h.CrossOrigin("anonymous"),
			),
		},
	})
}

func main() {
	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; object-src 'none';")
		homePage().Render(w)
	})

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
