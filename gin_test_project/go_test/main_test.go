package main

import (
"net/http"
"net/http/httptest"
"testing"

"github.com/stretchr/testify/assert"
)

// https://gin-gonic.com/zh-cn/docs/testing/

// 怎样编写 Gin 的测试用例, HTTP 测试首选 net/http/httptest 包

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}