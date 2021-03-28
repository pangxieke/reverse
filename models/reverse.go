package models

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"reverse/config"
	"reverse/log"
)

// NewReverseProxy 创建反向代理处理方法
func NewReverseProxy() *httputil.ReverseProxy {
	//创建Director
	director := reverseDirector()

	resp := reverseModifyResp()

	return &httputil.ReverseProxy{Director: director, ModifyResponse: resp}
}

func reverseDirector() func(r *http.Request) {
	//创建Director
	director := func(r *http.Request) {
		//查询原始请求路径
		reqPath := r.URL.Path
		if reqPath == "" {
			return
		}

		//设置代理服务地址信息
		scheme := "http"
		r.URL.Scheme = scheme
		r.URL.Host = fmt.Sprintf("%s:%d", config.OpenApi.Host, config.OpenApi.Port)

		if r.Body == nil {
			log.Info(r.URL.Path + "request body is empty")
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.ErrorLog.Info(r.URL.Path + "error" + err.Error())
		}
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	return director
}

func reverseModifyResp() func(res *http.Response) error {
	resp := func(res *http.Response) error {

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		res.Body.Close()
		res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return nil
	}
	return resp
}
