## What
A dead simple utility for turning strings into bcrypt hashes

## How

`bcrypt_hash_password hunter2`

Gives you:

```
hunter2
$2a$10$z.otAeBGgGYL6dtYAM.W/u5T2vv6U5ii9vksCage6/KKgqdSxKxx2
```

## Why

Sometimes when developing an application I want to set a user password directly in the DB without going through the app's password reset mechanism.

## Gimmie

`go get github.com/davidbanham/bcrypt_hash_password`

## Not bad!

https://notbad.software
