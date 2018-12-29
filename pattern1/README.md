### Reference

https://qiita.com/hirotakan/items/698c1f5773a3cca6193e

```
domain→Entities
infrastructure→Frameworks & Drivers
interfaces→Interfaces
usecase→Use cases
```

```
$ docker run --name mysql -e MYSQL_ROOT_PASSWORD=mysql -d -p 13306:3306 mysql
$ mysql -h 127.0.0.1 -uroot -pmysql --port 13306
 → create database sample;
```


```
# ユーザー作成
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"FirstName": "Susan", "LastName": "Taylor"}' localhost:8080/users

# ユーザー一覧
$ curl -i -H 'Content-Type:application/json' localhost:8080/users

# ユーザー取得
$ curl -i -H 'Content-Type:application/json' localhost:8080/users/3
```
