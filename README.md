# Todo API (Clean Architecuture)

## Docker環境セットアップ

### サーバー起動
```
docker-compose up -d
```

### UnitTest
* Mockの作成
```
mockgen -source domain/repository/todo.go -destination mock_todo/mock_repository_todo.go
```
* テストの実行
```
# usecase
go test -v ./mock_todo/usecase_todo_test.go
# repository
go test -v ./infrastructure/persistence/todo_test.go
```

## API 

### Todo
* 作成
```shell script
curl -X POST "http://localhost:8080/todos" -H "accept: application/json" --data-raw '{$jsonData}'
```

* 取得
```shell script
curl -X GET "http://localhost:8080/todos/{$id}" -H "accept: application/json"
```

* 更新
```shell script
curl -X PUT "http://localhost:8080/todos/{$id}" -H "accept: application/json" --data-raw '{$jsonData}'
```

* 削除
```shell script
curl -X DELETE "http://localhost:8080/todos/{$id}" -H "accept: application/json"
```

* リスト取得
```shell script
curl -X GET "http://localhost:8080/todos" -H "accept: application/json"
```