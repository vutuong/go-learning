package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

// Định nghĩa struct để xử lý convert html Template gán nó vào ResponseWriter
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8888", "The addr of the application.")
	flag.Parse() // parse the flags
	key := "256010564559-faafn08cd6jfoshma1l7aqfib8c5j43v.apps.googleusercontent.com"
	secret := "Qy8hilOt8x-GeKsOWkrNDUX9"
	// setup gomniauth
	gomniauth.SetSecurityKey("PUT YOUR AUTH KEY HERE")
	gomniauth.WithProviders(
		facebook.New(key, secret,
			"http://localhost:8080/auth/callback/facebook"),
		github.New(key, secret,
			"http://localhost:8080/auth/callback/github"),
		google.New(key, secret,
			"https://localhost:9090/auth/callback/google"),
	)
	r := newRoom()
	// r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	// run room trong các go-routine để các chatting được hoạt động ở background trong khi
	// main routine dùng dể chạy the web server
	go r.run()
	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
