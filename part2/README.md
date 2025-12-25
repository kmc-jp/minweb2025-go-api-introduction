- Dockerイメージのビルド
```shell
docker build -t go-app .
```
- Dockerコンテナの実行
```shell
docker run -p 8080:8080 go-app
```

### 動作確認
- 新しいタスクの作成
```shell
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "スライド作成1", "description": "API講座①のスライドを作成する"}' \
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "スライド作成2", "description": "API講座②のスライドを作成する"}' \
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "スライド作成3", "description": "API講座③のスライドを作成する"}'
```

- タスク一覧の取得
```shell
curl http://localhost:8080/tasks
```

- タスク詳細の取得
```shell
curl http://localhost:8080/tasks/1
```

- タスクの更新
```shell
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "スライド作成1 (完了)", "completed": true}'
```

- タスクの削除
```shell
curl -X DELETE http://localhost:8080/tasks/3
```

- タスク一覧の取得
```shell
curl http://localhost:8080/tasks
```