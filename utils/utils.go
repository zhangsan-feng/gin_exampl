package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetParams(r *gin.Context, keyParams string) interface{} {
	contentType := r.GetHeader("Content-Type")
	params := make(map[string]interface{})

	for key, values := range r.Request.URL.Query() {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	if r.Request.Method == "POST" || r.Request.Method == "PUT" || r.Request.Method == "PATCH" {
		if strings.Contains(contentType, "application/json") {
			var jsonData map[string]interface{}
			if err := r.ShouldBindJSON(&jsonData); err == nil {
				for k, v := range jsonData {
					params[k] = v
				}
			}
		}

		if strings.Contains(contentType, "application/x-www-form-urlencoded") ||
			strings.Contains(contentType, "multipart/form-data") {
			if err := r.Request.ParseForm(); err == nil {
				for k, v := range r.Request.PostForm {
					if len(v) > 0 {
						params[k] = v[0]
					}
				}
			}

			if strings.Contains(contentType, "multipart/form-data") {
				if err := r.Request.ParseMultipartForm(32 << 20); err == nil {
					for k, v := range r.Request.MultipartForm.Value {
						if len(v) > 0 {
							params[k] = v[0]
						}
					}
				}
			}
		}
	}
	return params[keyParams]
}
