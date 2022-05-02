docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5432:5432 \
  postgres:12.5-alpine


  docker run \
  -d \
  -p 11211:11211 \
  memcached:1.6.9-alpine

DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable \
MEMCACHED=localhost:11211 \
  go run .