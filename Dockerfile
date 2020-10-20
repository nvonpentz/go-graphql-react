FROM golang:alpine AS go_builder
WORKDIR /go-graphql-react
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN apk add --update make
RUN make server

FROM node:alpine AS node_builder
WORKDIR /go-graphql-react
COPY . ./
RUN apk add --update make
RUN make client

FROM alpine
WORKDIR /go-graphql-react
COPY --from=go_builder /go-graphql-react/ /go-graphql-react/
COPY --from=node_builder /go-graphql-react/frontend/build /go-graphql-react/build

CMD ["./bin/server"]
