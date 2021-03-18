FROM golang:1.16.2-alpine as build

WORKDIR /home/app
COPY . /home/app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./server app/main.go

FROM scratch
COPY --from=build /home/app .
CMD [ "./server" ]
EXPOSE 9999
