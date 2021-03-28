package middlewares

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"reverse/log"
	"strings"
	"time"
)

type response struct {
	http.ResponseWriter
	Status int
	Body   []byte
}

func (r *response) Write(data []byte) (int, error) {
	r.Body = data
	return r.ResponseWriter.Write(data)
}

func (r *response) WriteHeader(statusCode int) {
	r.Status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now().UnixNano()

		headers := []string{}
		for k, v := range r.Header {
			headers = append(headers, fmt.Sprintf("%v: %v", k, v[0]))
		}

		reqHead := fmt.Sprintf("Request headers: %s", strings.Join(headers, ", "))
		body, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		reqBody := fmt.Sprintf("Request body:%s", body)

		// Call the next handler in the chain.
		res := response{ResponseWriter: w}
		next.ServeHTTP(&res, r)

		headers = []string{}
		for k, v := range res.Header() {
			headers = append(headers, fmt.Sprintf("%v: %v", k, v[0]))
		}

		elapse := float64(time.Now().UnixNano()-begin) / 1000000.0

		logEntry := log.WithFields(logrus.Fields{
			"access_time":     time.Now(),
			"ip":              r.RemoteAddr,
			"method":          r.Method,
			"path":            r.URL.Path,
			"query":           r.URL.RawQuery,
			"request_header":  reqHead,
			"request_body":    reqBody,
			"server":          r.URL.Host,
			"response_header": headers,
			"response_body":   string(res.Body),
			"status":          res.Status,
			"elapse":          elapse,
		})
		logEntry.Info()

	})
}
