FROM golang:1.17 AS build

ARG PORT=9000

ENV PORT=${PORT}

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o api .

FROM alpine AS final

ARG DATABASE_URL=root:123456@tcp(host.docker.internal:3306)/api?parseTime=true
ARG JWT_SECRET=API-TASKS

ENV IS_PRODUCTION=true
ENV DATABASE_URL=${DATABASE_URL}
ENV JWT_SECRET=${JWT_SECRET}

COPY --from=build /app/api .

EXPOSE $PORT

ENTRYPOINT ./api
