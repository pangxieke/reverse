package middlewares

import (
	"net/http"
	"reverse/log"
	"reverse/models"
	"strings"
	"time"
)

func Access(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		// 记录 Header中部分数据到mongodb
		go AccessLog(r)
	})
}

func AccessLog(r *http.Request) {
	appKey := AuthorizeKey(r.Header.Get("Authorization"))
	entId := AuthorizeKey(r.Header.Get("ent_id"))

	if appKey != "" {
		accessData := make(map[string]interface{})
		accessData["ent_id"] = entId
		accessData["app_key"] = appKey
		accessData["path"] = r.RequestURI
		accessData["time"] = time.Now()
		err := models.MgoAdd("access", accessData)
		if err != nil {
			log.Info("mgo saver error", err)
		}
	}
}

// 解析header中authorizetion字段数据
// 例如 Authorization: app_key:20210309;sign:fa323252cdbd3e7a3835a142a126adab
func AuthorizeKey(auth string) (appKey string) {
	if auth == "" {
		return
	}

	var sign string

	items := strings.Split(auth, ";")
	for _, item := range items {
		kv := strings.Split(item, ":")
		if len(kv) == 2 {
			key := kv[0]
			value := kv[1]
			if key == "app_key" {
				appKey = value
			} else if key == "sign" {
				sign = value
			}
		}
	}
	_ = sign
	return
}
