FROM golang:1.21-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /out/frontend .

FROM scratch
WORKDIR /app
COPY --from=build /out/frontend /app/frontend
COPY templates /app/templates
COPY static /app/static
EXPOSE 8080
ENTRYPOINT ["/app/frontend"]
