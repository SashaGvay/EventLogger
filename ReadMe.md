## EventLogger - log weather in postgres table with kafka
build on: golang, kafka, postgres, golang-migrate

---

first copy .env

    cp .env.dist .env



add kafka alias in host

```
echo "127.0.0.1 kafka" | sudo tee -a /etc/hosts > /dev/null
```

migrate database

```
make migrate
```

run app:

    make run

---

redpanda-console is available at http://localhost:8081/