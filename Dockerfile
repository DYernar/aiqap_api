FROM golang:latest
WORKDIR /usr/src/app/
COPY . /usr/src/app/
RUN go mod download
RUN go build -o aiqap .
EXPOSE 8080
ENV TZ Asia/Almaty
CMD ["./aiqap"]

