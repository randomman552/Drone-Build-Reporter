FROM golang:1.21.3-alpine3.18 as build
WORKDIR /build
COPY . .
RUN go build -o entrypoint .

FROM alpine:3.18 as final
COPY --from=build /build/entrypoint /
COPY --from=build /build/templates /templates

ENTRYPOINT [ "/entrypoint" ]