package main

import (
	"compress/flate"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/itchio/go-brotli/enc"
	"github.com/lucas-clemente/quic-go/http3"
)

type compressResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w compressResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w compressResponseWriter) WriteHeader(code int) {
	w.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(code)
}

func gzipHandler(fn http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	if err != nil {
		fmt.Printf("Error closing gzip: %+v\n", err)
	}
	defer func() {
		err := gz.Close()
		if err != nil {
			fmt.Printf("Error closing gzip: %+v\n", err)
		}
	}()
	gzr := compressResponseWriter{Writer: gz, ResponseWriter: w}
	fn(gzr, r)
}

func deflateHandler(fn http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "deflate")
	df, err := flate.NewWriter(w, flate.BestSpeed)
	if err != nil {
		fmt.Printf("Error closing deflate: %+v\n", err)
	}
	defer func() {
		err := df.Close()
		if err != nil {
			fmt.Printf("Error closing deflate: %+v\n", err)
		}
	}()
	dfr := compressResponseWriter{Writer: df, ResponseWriter: w}
	fn(dfr, r)
}

func brotliHandler(fn http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "br")

	op := &enc.BrotliWriterOptions{
		Quality: 4,
		LGWin:   10,
	}
	br := enc.NewBrotliWriter(w, op)
	defer func() {
		err := br.Close()
		if err != nil {
			fmt.Printf("Error closing brotli: %+v\n", err)
		}
	}()
	brr := compressResponseWriter{Writer: br, ResponseWriter: w}
	fn(brr, r)
}

type handler struct {
	enableCompress bool
}

func (h *handler) makeCompressionHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if h.enableCompress {
			w.Header().Set("Vary", "Assept-Encoding")
			if strings.Contains(r.Header.Get("Accept-Encoding"), "br") {
				brotliHandler(fn, w, r)
				return
			} else if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				gzipHandler(fn, w, r)
				return
			} else if strings.Contains(r.Header.Get("Accept-Encoding"), "deflate") {
				deflateHandler(fn, w, r)
				return
			}
		}

		fn(w, r)
	}
}

func reverseProxy(backEndUrl string) func(http.ResponseWriter, *http.Request) {
	url, err := url.Parse(backEndUrl)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	return proxy.ServeHTTP
}

func main() {
	var (
		bg             = flag.String("b", "http://localhost:9000", "background")
		port           = flag.Uint("p", 8080, "front port")
		useHttp3       = flag.Bool("enable_http3", false, "use http3")
		certPath       = flag.String("c", "./testdata/cert.pem", "cert file path")
		keyPath        = flag.String("k", "./testdata/priv.key", "key file path")
		enableCompress = flag.Bool("enable_compress", false, "enable compress")
	)
	flag.Parse()

	h := handler{
		enableCompress: *enableCompress,
	}

	proxyServer := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
		Handler: h.makeCompressionHandler(reverseProxy(*bg)),
	}

	if *useHttp3 {
		fmt.Println("http3")
		if err := http3.ListenAndServe(proxyServer.Addr, *certPath, *keyPath, proxyServer.Handler); err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Println("http")
		if err := proxyServer.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}
}
