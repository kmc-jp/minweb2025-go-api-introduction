# Part3: 実践的な機能拡張ハンズオン

## セットアップ

### 1. Docker Composeで起動
```shell
docker-compose up --build
```

### 2. 別ターミナルで動作確認

## 動作確認手順

### Step 1: ユーザー登録
```shell
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```

### Step 2: ログイン（JWTトークン取得）
```shell
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```

レスポンス例:
```json
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}
```

**以降のコマンドで使用するため、トークンを環境変数に保存:**
```shell
export TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

---

## Task CRUD操作

### Task作成（複数）
```shell
# Task 1
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"スライド作成1","description":"API講座①のスライドを作成する"}'

# Task 2
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"スライド作成2","description":"API講座②のスライドを作成する"}'

# Task 3
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"スライド作成3","description":"API講座③のスライドを作成する"}'
```

### Task一覧取得
```shell
curl http://localhost:8080/tasks
```

### Task詳細取得（ID: 1）
```shell
curl http://localhost:8080/tasks/1
```

### Task更新（ID: 1を完了状態に）
```shell
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"スライド作成1 (完了)","completed":true}'
```

### Task削除（ID: 3）
```shell
curl -X DELETE http://localhost:8080/tasks/3
```

### 更新後のTask一覧確認
```shell
curl http://localhost:8080/tasks
```

---

## Schedule CRUD操作（認証必須）

### Schedule作成（Task 1用）
```shell
curl -X POST http://localhost:8080/schedules/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "task_id": 1,
    "start_at": "2025-01-20T10:00:00Z",
    "end_at": "2025-01-20T12:00:00Z"
  }'
```

### Schedule作成（Task 2用）
```shell
curl -X POST http://localhost:8080/schedules/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "task_id": 2,
    "start_at": "2025-01-21T14:00:00Z",
    "end_at": "2025-01-21T16:00:00Z"
  }'
```

### Schedule一覧取得
```shell
curl http://localhost:8080/schedules/ \
  -H "Authorization: Bearer $TOKEN"
```

### Schedule詳細取得（ID: 1）
```shell
curl http://localhost:8080/schedules/1 \
  -H "Authorization: Bearer $TOKEN"
```

### 特定TaskのSchedule取得（Task ID: 1）
```shell
curl http://localhost:8080/schedules/tasks/1/schedules \
  -H "Authorization: Bearer $TOKEN"
```

### Schedule更新（ID: 1の時間を変更）
```shell
curl -X PUT http://localhost:8080/schedules/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "start_at": "2025-01-20T09:00:00Z",
    "end_at": "2025-01-20T11:00:00Z"
  }'
```

### Schedule削除（ID: 2）
```shell
curl -X DELETE http://localhost:8080/schedules/2 \
  -H "Authorization: Bearer $TOKEN"
```

### 更新後のSchedule一覧確認
```shell
curl http://localhost:8080/schedules/ \
  -H "Authorization: Bearer $TOKEN"
```

---

## 認証エラーのテスト

### トークンなしでScheduleにアクセス（401エラー）
```shell
curl http://localhost:8080/schedules/
```

### 無効なトークンでScheduleにアクセス（401エラー）
```shell
curl http://localhost:8080/schedules/ \
  -H "Authorization: Bearer invalid_token"
```

---

## テスト実行

### ユニットテスト実行
```shell
go test -v ./...
```

### 特定パッケージのテスト
```shell
go test -v ./internal/service
go test -v ./internal/handler
```

---

## クリーンアップ

### コンテナ停止・削除
```shell
docker-compose down
```

### ボリュームも含めて削除
```shell
docker-compose down -v
```
