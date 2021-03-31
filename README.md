# go-restfulapi-swagger

Esempio di creazione di restful API strutturata basata su tecnologia ORM e standard Swagger con restituzione controllata degli errori in formato JSON.

localhost:8080/swagger/index.html

localhost:8080/api/users

localhost:8080/api/users/{id}


# Install ORM database libraries
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get -u github.com/lib/pq

# Install Swagger API libraries 
go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template

# Build Swagger documents in project dir (~/GOPATH/src/go-restfulapi-swagger normally)
#RUN swag init --parseDependency --parseInternal -g main.go

# Build the Go app
#RUN go build -o main 
