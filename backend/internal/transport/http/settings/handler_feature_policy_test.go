package settings

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DEEIX-AI/DEEIX-Chat/backend/internal/infra/config"
	"github.com/gin-gonic/gin"
)

func TestGetFeaturePolicyIncludesUserAboutEnabled(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := &Handler{runtime: config.NewRuntime(config.Config{
		KnowledgeBaseEnabled: true,
		UserAboutEnabled:     false,
	})}
	router := gin.New()
	router.GET("/settings/feature-policy", handler.GetFeaturePolicy)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/settings/feature-policy", nil))
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	var body struct {
		Data FeaturePolicyResponse `json:"data"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if !body.Data.KnowledgeBaseEnabled {
		t.Fatalf("expected knowledgeBaseEnabled true")
	}
	if body.Data.UserAboutEnabled {
		t.Fatalf("expected userAboutEnabled false by default")
	}
}
