package handler_test

import (
	"net/http"
	"net/http/httptest"
	"nike_store_api/internal/handler"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	v1 := r.Group("/api/v1/")
	{
		v1.GET("/ping", handler.Ping)
	}

	req := httptest.NewRequest("GET", "/api/v1/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("انتظار کد 200 داشتیم ولی %d دریافت شد", w.Code)
	}

	expectedBody := `{"message":"pong"}`
	if w.Body.String() != expectedBody {
		t.Errorf("انتظار %s داشتیم ولی %s دریافت شد", expectedBody, w.Body.String())
	}
}
