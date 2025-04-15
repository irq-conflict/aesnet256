package aesnet256

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
)

// Map to provide for allowed AES blocksizes
// Needs to be initialized in init()
var AllowedSizes map[string]int

//go:embed htmx/decrypt.html
//go:embed htmx/encrypt.html
//go:embed htmx/style.css
//go:embed htmx/lib/htmx.1.19.12.min.js
//go:embed htmx/aki.png
//go:embed htmx/index.html
var embFs embed.FS

func init() {
	// Setup the allowed sizes map
	AllowedSizes = map[string]int{
		"128": Aes_128_length,
		"192": Aes_192_length,
		"256": Aes_256_length,
	}
	return
}

// Request is designed to be json marshaled into
// from an API request.
type Request struct {
	Iv           string
	Key          string
	KeySize      string
	Text         string
	Base64Encode bool
}

// Response data sent back to the UI
type Response struct {
	Text string
}

// encrypt API endpoint HTTP Handler
func encrypt(w http.ResponseWriter, r *http.Request) {
	tmp := &Request{}
	tb, e := io.ReadAll(r.Body)
	if e != nil {
		log.Printf("encrypt io.ReadAll error: %s", e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	e = json.Unmarshal(tb, tmp)
	if e != nil {
		log.Printf("encrypt.JsonUnmarshal error: %s", e)
		log.Println(string(tb))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", e)))
		return
	}
	aes_size, ok := AllowedSizes[tmp.KeySize]
	if !ok {
		log.Printf("AES size is invalid, must be 128, 192 or 256.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("AES size is invalid, must be 128, 192 or 256."))
		return
	}
	encReq := &Aes256{Source: []byte(tmp.Text), IV: []byte(Default_aesnet_iv), EncodeDest: true}
	if tmp.Iv != "" {
		encReq.IV = []byte(tmp.Iv)
	}
	e = encReq.Encrypt(PadKey([]byte(tmp.Key), aes_size))
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", e)))
		return
	}
	tb, e = json.Marshal(Response{Text: fmt.Sprintf("%s", encReq)})
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", e)))
		return

	}
	w.WriteHeader(http.StatusOK)
	w.Write(tb)
	return
}

// decrypt API endpoint API Handler
func decrypt(w http.ResponseWriter, r *http.Request) {
	tmp := &Request{}
	tb, e := io.ReadAll(r.Body)
	if e != nil {
		log.Printf("encrypt io.ReadAll error: %s", e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	e = json.Unmarshal(tb, tmp)
	if e != nil {
		log.Printf("encrypt.JsonUnmarshal error: %s", e)
		log.Println(string(tb))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("encrypt.JsonUnmarshal error: %s", e)))
		return
	}
	aes_size, ok := AllowedSizes[tmp.KeySize]
	if !ok {
		log.Printf("AES size is invalid, must be 128, 192 or 256.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("AES size is invalid, must be 128, 192 or 256."))
		return
	}
	encReq := &Aes256{Source: []byte(tmp.Text), IV: []byte(Default_aesnet_iv)}
	if tmp.Iv != "" {
		encReq.IV = []byte(tmp.Iv)
	}
	e = encReq.Decrypt(PadKey([]byte(tmp.Key), aes_size))
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", e)))
		return
	}
	tb, e = json.Marshal(Response{Text: fmt.Sprintf("%s", encReq)})
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", e)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(tb)
	return
}

// ping API endpoint Handler
func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PONG!"))
	return
}

// StaticHandler returns a function that is compatible with http.HandleFunc to display
// static HTML, CSS and images to facilitate the web UI. Errors are directly written to
// the underlying http.ResponseWriter.
func StaticHandler(fName, mimeType string) (fx func(http.ResponseWriter, *http.Request)) {
	fx = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mimeType)
		tmp, e := fs.ReadFile(embFs, fName)
		if e != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("%s", e)))
			return
		}
		w.Write(tmp)
		return
	}
	return
}

// RedirectHandler returns a function that is compatible with http.HandleFunc to
// provide HTTP redirects.
func RedirectHandler(dest string, statusCode int) (fx func(http.ResponseWriter, *http.Request)) {
	if statusCode < http.StatusMultipleChoices || statusCode > http.StatusPermanentRedirect {
		statusCode = http.StatusTemporaryRedirect
	}
	fx = func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dest, statusCode)
		return
	}
	return
}

// This is the main ListenAndServe block that registers the endpoints, static handler
// and redirectors.
func ListenAndServe(listenOn string) (e error) {
	http.HandleFunc("/api/encrypt", encrypt)
	http.HandleFunc("/api/decrypt", decrypt)
	http.HandleFunc("/api/ping", ping)
	http.HandleFunc("/", RedirectHandler("/index.html", http.StatusPermanentRedirect))
	http.HandleFunc("/index.html", StaticHandler("htmx/index.html", "text/html; charset=utf-8"))
	http.HandleFunc("/encrypt.html", StaticHandler("htmx/encrypt.html", "text/html; charset=utf-8"))
	http.HandleFunc("/decrypt.html", StaticHandler("htmx/decrypt.html", "text/html; charset=utf-8"))
	http.HandleFunc("/htmx.min.js", StaticHandler("htmx/lib/htmx.1.19.12.min.js", "text/javascript; charset=utf-8"))
	http.HandleFunc("/style.css", StaticHandler("htmx/style.css", "text/css; charset=utf-8"))
	http.HandleFunc("/favicon.ico", StaticHandler("htmx/aki.png", "image/png"))
	log.Printf("Server mode active. Read more here https://github.com/irq-conflict/aesnet256.")
	log.Printf("To use the server, go to the url http://localhost%s/encrypt.html -or- http://localhost%s/decrypt.html", listenOn, listenOn)
	log.Printf("Your local firewall will need to allow access to the port.")
	log.Printf("If you encouter issues, log them at https://github.com/irq-conflict/aesnet256/issues")
	e = http.ListenAndServe(listenOn, nil)
	if e != nil {
		log.Fatal("Server error: ", e)
	}
	return
}
