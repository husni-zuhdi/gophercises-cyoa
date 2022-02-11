package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cyoa "github.com/hazunanafaru/gophercises-cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON story file you will use")
	port := flag.Int("port", 3000, "The port to start CYOA web app")
	flag.Parse()
	fmt.Printf("Using the story in %s and port %d.\n", *filename, *port)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	// Default Handler
	h := cyoa.NewHandler(story)

	// For testing the Functional Options
	// story_tpl := template.Must(template.New("").Parse(storyTpl))
	// h := cyoa.NewHandler(story,
	// 	cyoa.WithTemplate(story_tpl),
	// 	cyoa.WithPathFn(pathFn),
	// )

	// Add Mux to handle StatusPage(?)
	mux := http.NewServeMux()
	// mux.Handle("/story/", h)
	mux.Handle("/", h)

	fmt.Printf("Starting the server on localhost:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

// TEST THE FUNCTIONAL OPTIONS
// func pathFn(r *http.Request) string {
// 	// Check if the path is empty. If empty, set it to intro
// 	path := strings.TrimSpace(r.URL.Path)
// 	if path == "/story" || path == "/story/" {
// 		path = "/story/intro"
// 	}
// 	// Trim the "/..." from path
// 	return path[len("/story/"):]
// }

// var storyTpl = `
// <!DOCTYPE html>
// <html>
//     <head>
//         <meta charset="utf-8">
//         <title>Choose Your Own Adventure</title>
//     </head>
//     <body>
// 		<section class="page">
// 			<h1>{{.Title}}</h1>
// 			{{range .Paragraph}}
// 				<p>{{.}}</p>
// 			{{end}}
// 			<ul>
// 			{{range .Options}}
// 				<li><a href="/story/{{.Chap}}">{{.Text}}</a></li>
// 			{{end}}
// 			</ul>
// 		</section>
// 		<style>
// 		body {
// 			font-family: helvetica, arial;
// 		}
// 		h1 {
// 			text-align:center;
// 			position:relative;
// 		}
// 		.page {
// 			width: 80%;
// 			max-width: 500px;
// 			margin: auto;
// 			margin-top: 40px;
// 			margin-bottom: 40px;
// 			padding: 80px;
// 			background: #3273A8;
// 			border: 1px solid #eee;
// 			box-shadow: 0 10px 6px -6px #777;
// 		}
// 		ul {
// 			border-top: 1px dotted #ccc;
// 			padding: 10px 0 0 0;
// 			-webkit-padding-start: 0;
// 		}
// 		li {
// 			padding-top: 10px;
// 		}
// 		a,
// 		a:visited {
// 			text-decoration: none;
// 			color: #6295b5;
// 		}
// 		a:active,
// 		a:hover {
// 			color: #7792a2;
// 		}
// 		p {
// 			text-indent: 1em;
// 		}
// 		</style>
//     </body>
// </html>
// `
