package main
import (
	"log"
	"net/http"
	"sync"
	"text/template"
	"path/filepath"
)
// templ은 하나의 템플릿을 나타냄
type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}
// ServeHTTP가 HTTP 요청을 처리한다
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}
func main() {
/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<html>
			<head>
				<title>Chat</title>
			</head>
			<body>
				Let's Chat!
			</body>
		</html>
		`))
	})
*/
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// 방을 가져옴
	go r.run()
	// 웹 서버 시작
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}