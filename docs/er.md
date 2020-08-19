```uml
@startuml
' ユーザのリソース
entity ユーザ {
    + ID [PK]
    --
    email
    password
    名前
}

' Todoのリソース
entity Todo {
    + ID [PK]
    --
    ＃ ユーザID [FK(ユーザ.ID)]
    締切
    TODO内容
}

ユーザ||--|{Todo

@enduml
```