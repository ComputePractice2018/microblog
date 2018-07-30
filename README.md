# microblog

## Usecases

1. Как пользователь, я хочу иметь возможность создать свой профиль (ник, имя, фамилия, email, github), чтобы иметь свой профиль.
1. Как пользователь, я хочу иметь возможность редактировать свой профиль, чтобы исправить ошибки при создании.
1. Как пользователь, я хочу иметь возможность просмотреть все публикации, чтобы использовать эту информацию. 
1. Как пользователь, я хочу иметь возможность искать публикации по названию, чтобы быстро получить нужную информацию.
1. Как пользователь, я хочу иметь возможность просмотреть комментарии публикации, чтобы использовать ее.
1. Как пользователь, я хочу иметь возможность комментировать публикацию, чтобы прокомментировать публикацию.
1. Как пользователь, я хочу иметь возможность редактировать комментарий, чтобы исправить ошибки при создании комментария.
1. Как пользователь, я хочу иметь возможность удалить комментарий, чтобы не хранить неактуальную информацию.
1. Как пользователь, я хочу иметь возможность добавить публикацию, чтобы добавить свою публикацию.
1. Как пользователь, я хочу иметь возможность редактировать публикацию, чтобы исправить ошибки при добавлении.
1. Как пользователь, я хочу иметь возможность удалить публикацию, чтобы не хранить неактуальную информацию.

## REST API

### POST /api/microblog/profiles

Тело запроса:
```json
    {
        "nikname": "Ник",
        "name": "Имя",
        "surname": "Фамилия",
        "email": "user@domain.ru",
        "github": "user"
    }
```
Ответ: 201 Created
Location: /api/microblog/profiles/1

### PUT /api/microblog/profiles/1

Тело запроса:
```json
    {
        "nikname": "Ник",
        "name": "Имя",
        "surname": "Фамилия",
        "email": "user@domain.ru",
        "github": "user"
    }
```
Ответ: 202 Accepted
Location: /api/microblog/profiles/1



### GET /api/microblog/profiles/1/publications

Ответ: 200 ОК
```json
    [{
        "namepub": "Название публикации",
        "time": "Время",
        "publication": "Публикация",
    }]
```

### GET /api/microblog/profiles/1/publications?namepub=Название публикации

Ответ: 200 ОК
```json
    [{
        "namepub": "Название публикации",
        "time": "Время",
        "publication": "Публикация",
    }]
```

### GET /api/microblog/profiles/1/publications/1/comments

Ответ: 200 ОК
```json
   [{
       "time": "Время",
        "comment": "Комментарий"
    }]
```

### POST /api/microblog/profiles/1/publications/1/comments

Тело запроса:
```json
    [{
        "time": "Время",
        "comment": "Комментарий"
    }]
```
Ответ: 201 Created
Location: /api/microblog/profiles/1/publications/1/comments/1

### PUT /api/microblog/profiles/1/publications/1/comments/1

Тело запроса:
```json
    [{
        "time": "Время",
        "comment": "Комментарий"
    }]
```
Ответ: 202 Accepted
Location: /api/microblog/profile/1/publications/1/comments/1

### DELETE /api/microblog/profiles/1/publications/1/comments/1

Ответ: 204 No Content

### POST /api/microblog/profiles/1/publications

Тело запроса:
```json
    {
        "namepub": "Название публикации",
        "time": "Время",
        "publication": "Публикация",
    }
```
Ответ: 201 Created
Location: /api/microblog/profile/1/publications/1

### PUT /api/microblog/profiles/1/publications/1

Тело запроса:
```json
    {
        "namepub": "Название публикации",
        "time": "Время",
        "publication": "Публикация",
    }
```
Ответ: 202 Accepted
Location: /api/microblog/profile/1/publications/1

### DELETE /api/microblog/profiles/1/publications/1

Ответ: 204 No Content

## как собрать и запустить 

Backend:

'''bat
cd backend
docker build -f Dockerfile -t microblogbackend:<имя ветки> .
docker run --rm --name microblogbackend -e NAME=<параметр приложения> microblogbackend:<имя ветки>
'''

Frontend:

'''bat
cd frontend
docker build -f Dockerfile -t microblogfrontend:<имя ветки> .
docker run -d --rm --name microblogfrontend -p 80:80 microblogfrontend:<имя ветки>
'''