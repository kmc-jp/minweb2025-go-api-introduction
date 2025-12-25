package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"part3/internal/dto"
	"part3/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTask(t *testing.T) {
	// 1. 準備
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Serviceのモックを作成
	mockService := service.NewMockTaskService(ctrl)

	// Handlerにモックを注入
	h := NewTaskHandler(mockService)

	// Ginのテストモード設定
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/tasks", h.CreateTask)

	// 2. 期待する振る舞いの定義
	reqBody := dto.CreateTaskRequest{
		Title: "New Task",
	}
	expectedResponse := &dto.TaskResponse{
		ID:    1,
		Title: "New Task",
	}

	// Service.CreateTaskが呼ばれたら、成功レスポンスを返すように設定
	mockService.EXPECT().
		CreateTask(gomock.Any()).
		Return(expectedResponse, nil)

	// 3. リクエストの作成と実行
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// レスポンスを記録するレコーダー
	w := httptest.NewRecorder()

	// ハンドラ実行
	r.ServeHTTP(w, req)

	// 4. 検証
	assert.Equal(t, http.StatusCreated, w.Code) // ステータスコード 201

	var res dto.TaskResponse
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, "New Task", res.Title) // レスポンスの中身
}
