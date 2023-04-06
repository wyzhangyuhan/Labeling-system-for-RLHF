FROM golang:latest

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
ENV GOOS=linux 
ENV GOARCH=amd64

WORKDIR /usr/local/

RUN mkdir go_backend

COPY . /usr/local/go_backend/

WORKDIR /usr/local/go_backend/

# RUN go mod init 
RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "."]