FROM golang:latest
WORKDIR /usr/src/app/
COPY . /usr/src/app/
RUN go mod download
RUN go build -o aiqap .
EXPOSE 4444
ENV TZ Asia/Almaty
CMD ["./aiqap"]



#sudo docker build -t aiqap .
# docker run -d -p 4444:4444 -v ~/go/aiqap-back/audio:/usr/src/app/audio  aiqap
# docker run -d -p 4444:4444 -v ~/workspace/aiqap-api/audio:/usr/src/app/audio  aiqap