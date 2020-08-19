```uml
@startuml

title Todoリスト
actor "User" as user

box "User Interface" #e9ffe9
  participant "UI" as ui
end box

box "Application" #d4e1f5
  participant "App" as app
end box

box "Domain" #d4e1f5
  participant "Service" as service
end box

box "Infrastracture" #d4e1f5
  database "DB" as db
end box


user -> ui: Todo作成
ui -> app: POST /todos/{userId}
app -> service: CreateTodo
service -> db: INSERT todos
service <-- db:
app <-- service:
ui <-- app:

user -> ui: Todo取得
ui -> app: GET /todos/{userId}/{id}
app -> service: GetTodoById
service -> db: SELECT todos
service <-- db:
app <-- service:
ui <-- app:

user -> ui: Todo更新
ui -> app: PUT /todos/{userId}/{id}
app -> service: PutTodoById
service -> db: UPDATE todos
service <-- db:
app <-- service:
ui <-- app:

user -> ui: Todo削除
ui -> app: DELETE /todos/{userId}/{id}
app -> service: DeleteTodoById
service -> db: DELETE todos
service <-- db:
app <-- service:
ui <-- app:

user -> ui: Todoリスト取得
ui -> app: GET /todos/{userId}
app -> service: ListTodos
service -> db: SELECT todos
service <-- db:
app <-- service:
ui <-- app:

@enduml
```