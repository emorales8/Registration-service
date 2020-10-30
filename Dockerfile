FROM amd64/golang
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv
COPY post .
EXPOSE 8081

CMD ["/app/post"]