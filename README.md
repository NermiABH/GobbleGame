## <a href=#content id=routes>Routes</a>:

| <p id="№">№<p> | url                    | methods |
|:--------------:|:-----------------------|:--------|
|       1        | http: /user/register   | POST    |
|       2        | http: /user/login      | POST    |
|       3        | http: /user/refresh    | POST    |
|       4        | http: /unique/username | POST    |
|       5        | http: /unique/email    | POST    |


<p id="1"></p>

## <a href="#№">Register User</a>
```
POST http: /user/register
```
Request body:
```
{
    "username": "user",
    "email": "user@email",
    "password": "password"
}
```
Аутентификация не обязательна, возвращает <a href="#User">User</a>

Обязательные поля: `username, email, password`

## <a href="#№">Login</a>
Авторизация на основе jwt
```
POST http: /user/login 
```
Request body:
```
{
    "username": "user",
    "email": "user@email",
    "password": "password"
}
```
Аутентификация не обязательна, возвращает <a href="#tokens">Tokens</a>

Обязательные поля: `username или email, password` 

## <a href="#№">Refresh</a>
Авторизация на основе jwt
```
POST http: /user/refresh
```
Request body:
```
{
    "refresh": "something token"
}
```
Аутентификация не обязательна, возвращает <a href="#refresh">Refresh</a>

Обязательные поля: `refresh`

## <a href="#№">Delete User</a>
```
DELETE .../user/{id}
```
Аутентификация обязательна.


[//]: #                                         (ФОРМАТ ОТВЕТОВ)
# <a href="#content" id=response>Формат ответов</a>

## <p id=tokens>Tokens</p>
```
{
    "type"
    "attributes": {
        "access": "something token",
        "refresh": "something token"
    }
}

```

## <p id=user>User</p>
```
{
    "type": "user",
    "id": 1,
    "attributes": {
        "username": "nermiAbh",
        "email: "nermi@gmail.com"
    }
}
```

[//]: #                                         (ОШИБКИ)

# Коды статуса 


### 200 StatusOK
### 201 StatusCreated
Успешное создание ресурса(вместо 200)
