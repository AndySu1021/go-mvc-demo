package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mvc/utils/log"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		requestMap := make(map[string]string)
		if c.Request.Method == "GET" {
			queryParam, _ := json.Marshal(c.Request.URL.Query())
			requestMap["method"] = c.Request.Method
			requestMap["path"] = c.FullPath()
			requestMap["data"] = string(queryParam)
		}

		go func() {
			_ = log.Info(log.Fluentd{
				Tag:  "log.request",
				Data: requestMap,
			})
		}()

		c.Next()

		responseMap := make(map[string]string)
		responseMap["data"] = w.body.String()

		go func() {
			_ = log.Info(log.Fluentd{
				Tag:  "log.response",
				Data: responseMap,
			})
		}()
	}
}
