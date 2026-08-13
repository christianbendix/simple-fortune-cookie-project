FROM golang:1.21-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /out/backend .

FROM scratch
COPY --from=build /out/backend /backend
EXPOSE 9000
ENTRYPOINT ["/backend"]
