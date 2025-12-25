#set text(font: ("Roboto", "Noto Sans CJK JP"))
#set text(lang: "ja")

#set heading(numbering: none)

#import "@preview/slydst:0.1.5" : *
#import "@preview/fletcher:0.5.8" : *

#show: slides.with(
  title: "Web API講座 Part 3",
  subtitle: "実践的な機能拡張ハンズオン",
  authors: "48th irom",
)

#outline(depth: 2)

= Part 2の復習

== Part 2の復習
- REST APIの設計原則 (リソース, HTTPメソッド)
- Go言語とGinフレームワークによる実装
- レイヤードアーキテクチャ (Handler -> Service -> Repository)
- データ構造の分離 (Model vs DTO)

= 1. Scheduleモデルの追加 (リレーション)

== リレーショナルデータベースの基礎
- テーブル間の関係性
  - 1対1 (User - Profile)
  - 1対多 (Task - Schedules)
  - 多対多 (Student - Class)
- 外部キー (Foreign Key)
  - 子テーブルが親テーブルを参照するためのID
- GORMのPreload
  - 関連データを効率的に取得する (Eager Loading)

== 実装: Scheduleモデルの作成
- `internal/model/schedule.go` を作成
- TaskID を外部キーとして持つ
```go
package model
import "time"

type Schedule struct {
  ID      uint      `gorm:"primaryKey" json:"id"`
  TaskID  uint      `gorm:"not null" json:"task_id"` // FK
  StartAt time.Time `gorm:"not null" json:"start_at"`
  EndAt   time.Time `gorm:"not null" json:"end_at"`
  // リレーション定義
  Task    Task      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
```

== 実装: Taskモデルの更新
- internal/model/task.go
  - TaskからScheduleへの参照を追加 (1対多)
```go
type Task struct {
  // ...既存フィールド...
  
  // 1対多のリレーション
  Schedules []Schedule `json:"schedules,omitempty"`
}
```

== マイグレーションの更新
- cmd/api/main.go
  - AutoMigrateにScheduleモデルを追加
```go
db.AutoMigrate(&model.Task{}, &model.Schedule{})
```

= 2. テストコードの実装
== テストの種類
- ユニットテスト (単体テスト)
  - 関数やメソッド単位での検証
  - 外部依存 (DBなど) はモック化する

- 統合テスト (結合テスト)
  - 複数のモジュールを連携させた検証
  - 実際にDBやHTTPリクエストを使うこともある

== モックとDI
- 依存性注入 (DI) の恩恵
- Repositoryの実装をモックに差し替えることで、DBなしでService層のロジックをテスト可能
- GoMockやTestifyなどのライブラリを活用

```bash
go install github.com/golang/mock/mockgen@latest
``` 
== モックの生成
- internal/repository/task.go のインターフェースをモック化
```bash
mockgen -source=internal/repository/task.go \
  -destination=internal/repository/mock_task.go \
  -package=repository
```

== Service層のユニットテスト例
- internal/service/task_service_test.go
```go
func TestCreateTask(t *testing.T) {
  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  mockRepo := repository.NewMockTaskRepository(ctrl)
  // "Create"が呼ばれたらnil(エラーなし)を返す
  mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

  service := NewTaskService(mockRepo)
  task := &model.Task{Title: "Test Task"}
  err := service.CreateTask(task)
  if err != nil {
    t.Fatalf("expected no error, got %v", err)
  }
  
  mockRepo.Finish()
}
```

== テストの実行
```bash
go test -v ./... 
```

== Handler層の統合テスト例
```bash
mockgen -source=internal/service/task.go \
  -destination=internal/service/mock_task.go \
  -package=service
```

- internal/handler/task_test.go

= 3. 認証・認可の実装
== ステートレス認証とJWT
- ステートレス認証
  - サーバー側でセッション情報を保持しない
- JSON Web Token (JWT)
  - クライアントにトークンを発行し、各リクエストで送信
  - トークンにユーザー情報や権限を含める
  - 構成：ヘッダー、ペイロード、署名

== Middlewareの実装
- ハンドラの実行前に共通処理を挟む
- internal/middleware/auth.go
```go
func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    // ... トークン検証ロジック ...
    if err != nil {
      c.AbortWithStatus(401)
      return
    }
    c.Next() // 次のハンドラへ
  }
}
```

== ルーティングへの適用
- cmd/api/main.go
```go
// main.go
r := gin.Default()

// 公開API
r.POST("/login", loginHandler)

// 認証必須API
authGroup := r.Group("/api/v1")
authGroup.Use(middleware.AuthMiddleware())
{
  authGroup.POST("/tasks", taskHandler.CreateTask)
  // ...
}
```

== ユーザー登録とログイン
- internal/handler/auth.go
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"user1","password":"pass123"}'
```

= 4. DB永続化(Docker Compose)
== PostgreSQLの導入
- Docker ComposeでPostgreSQLコンテナを起動
- `docker-compose.yml`
```yaml
services:
  app:
    depends_on: [db]
    environment:
      - DB_HOST=db
      - DB_USER=user
      - DB_PASSWORD=password
  db:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
      - POSTGRES_DB=taskdb
    ports: ["5432:5432"]
volumes:
  db-data:
    driver: local
```

== DB接続の変更
- cmd/api/main.go
```go
import "gorm.io/driver/postgres"

// ...
dsn := "host=db user=user password=password dbname=todo port=5432 sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
  log.Fatal("failed to connect database:", err)
}
```

= 5. Swaggerドキュメントの生成
== Documentation as Code
- API仕様をコード内に記述
- 自動生成ツールでドキュメントを生成
- メンテナンス性の向上
- 実装と仕様の乖離を防止

== Swaggoの導入
- Swaggoを使ってSwaggerドキュメントを生成
```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```
== コメントによるAPIドキュメントの記述
- ハンドラ関数にコメントを追加
```go
// ListTasks godoc
// @Summary タスク一覧取得
// @Description 登録されているタスクを全件取得します
// @Tags tasks
// @Success 200 {array} dto.ListTasksResponse
// @Router /tasks [get]
func (h *TaskHandler) ListTasks(c *gin.Context) { ... }
```

== ドキュメントの生成
```bash
swag init -g cmd/api/main.go -o docs
```