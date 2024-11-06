package main

import (
	"Course-Management/routes"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLoginHandler(t *testing.T) {
	router := gin.Default()
	routes.Init(router)

	payload := []byte(`{"username": "admin", "password": "123456"}`)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 %d，实际状态码 %d", http.StatusOK, w.Code)
	}

	t.Logf("响应数据: %s", w.Body.String())
}
