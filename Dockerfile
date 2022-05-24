FROM golang:alpine as build
COPY . /project
WORKDIR /project
RUN go build ./main.go

FROM alpine:latest
COPY --from=build /project/main .
EXPOSE 4000
CMD ["./main"]