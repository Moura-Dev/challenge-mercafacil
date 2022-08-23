
# MercaFacil Challenge

## Run 

```bash
make services
```

```bash
go mod init mercafacil-challenge
```
Install Dependencies

```bash
  go mod tidy
```

Run Server Service

```bash
 make run
```

Endpoint Create User
#
```bash
url: localhost:5000/api/user

json:
{
    "login": "moura",
    "password": "123456",
    "customer": "macapa"
}
```

Endpoint Login


#
```bash
url: localhost:5000/api/login

json:
{
    "login": "moura",
    "password": "123456",
    "customer": "macapa"
}
```

using api with jwt
```bash
url: localhost:5000/api/

Authorization = Bearer "you-token-generate"
```


Collection Docs example in project:
docs/mercafacil.postman_collection.json
 