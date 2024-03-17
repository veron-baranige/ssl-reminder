FROM golang:1.22-alpine3.18 as build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /dist/ssl-reminder
RUN	chmod +x /dist/ssl-reminder

FROM alpine:3.14 as app
COPY --from=build /dist/ssl-reminder /bin/
ENTRYPOINT [ "/bin/ssl-reminder" ] 