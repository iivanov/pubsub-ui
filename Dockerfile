FROM golang:1.19.1 as build
COPY . /app
RUN cd /app && unset GOPATH && CGO_ENABLED=0 go build -a -installsuffix cgo -o /go/bin/pubsub-ui ./main.go

FROM alpine:3.8 as prod
RUN apk add --no-cache tini
COPY --from=build /go/bin/pubsub-ui /pubsub-ui
COPY --from=build /app/frontend/dist /frontend/dist
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/pubsub-ui"]
EXPOSE 8780