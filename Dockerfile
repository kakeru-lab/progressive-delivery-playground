# build
FROM golang:1.22 AS build
WORKDIR /src
COPY app/go.mod ./
RUN go mod download
COPY app/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/app ./...

# runtime
FROM gcr.io/distroless/static:nonroot
ENV PORT=8080
COPY --from=build /out/app /app
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app"]
