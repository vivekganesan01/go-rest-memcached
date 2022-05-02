# Caching in Go: memcached

## Running

Download the required packages:

```
go mod download
```

Then you can run `server.go` with the PostgreSQL and memcached configuration:

```
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable \
MEMCACHED=localhost:11211 \
  go run .
```

### PostgreSQL

#### Using Docker

```
docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5432:5432 \
  postgres:12.5-alpine
```

#### Migrations

Make sure the `migrate` tool is installed:

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
```

Then, assuming you used docker, you can run:

```
migrate -path db/migrations/ -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
```

### Memcached

```
docker run \
  -d \
  -p 11211:11211 \
  memcached:1.6.9-alpine
```

### Load testing data

(https://mariocarrion.com/2020/08/27/go-implementing-complex-pipelines-part-5.html), install it using:

```
go install github.com/MarioCarrion/complex-pipelines/part5
```

Then run it using:

```
DATABASE_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable" part5
```
