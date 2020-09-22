## Voting App

Endpoints Specification

```
Server is deployed on IP:142.93.221.156 PORT:80
```

click [HERE](http://142.93.221.156) to visit

```
GET /
```

- page for create new poll

```
GET /signup
```

- page for register a user

```
GET /login
```

- page for login a user

```
GET /polls/view/{poll_id}
```

- page for given poll_id to vote
- if already voted or vote is greater then 13 it won't render the poll

```
GET /polls/my/{page_number}
```

- shows 5 poll name per page created by user on given page number

```
GET /polls/others/{page_number}
```

- shows 5 poll per page created by any user

#### Content-Type: Application/json

```

POST /api/v1/user/signup

```

```json
{ "email": "mail@name.io:", "password": "qwerty" }
```

```json
response {"ok": false, "msg": "if ok is false"}
```

- register new user

```
POST /api/v1/user/login
```

- login new user
- cookie wil be set

```json
response {"ok": false, "msg": "if ok is false"}
```

```json
{ "email": "mail@name.io:", "password": "qwerty" }
```

````json
response {"ok": false:, "msg": "if ok is false"}
```

````

POST /api/v1/polls/create

````

- create a new poll
- auth required

```json
{
  "title": "what is you fev food?",
  "option1": "chocolate",
  "option2": "banana",
  "option3": "meggi",
  "option4": "biriyani",
  "option5": "cake"
}
````

```json
response {"ok": false, "msg": "if ok is false"}
```

```
POST /api/v1/polls/vote/{poll_id}/{option1-option5}
```

```json
response {"ok": false, "msg":"if ok is false"}
```
