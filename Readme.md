# Go CRUD Project with GORM (net/http and MySQL)
This is a simple Go project demonstrating CRUD operations using GORM as the entity framework with net/http and MySQL.

## Project Setup
Ensure you have Go installed on your system. You can download it from https://golang.org/dl/.

Install the necessary Go packages and dependencies by running:
```
go get -u gorm.io/gorm
```

Use the following command to clone the repository to your local machine.
   ```
   git clone https://github.com/your-username/golang-crud.git
   cd golang-crud
   ```

## Create a MySQL database and update the main.go file with your database credentials:

```
const (
	DBUser     = "your_database_username"
	DBPassword = "your_database_password"
	DBName     = "your_database_name"
)
```

## Run the project:

```
go run main.go
```


## Database Setup
Make sure you have MySQL installed on your system. You can download it from https://dev.mysql.com/downloads/installer/.

Connect to MySQL using your preferred MySQL client (e.g., MySQL Workbench or MySQL CLI).

Create a new database (if it doesn't exist) for the project:

```
CREATE DATABASE your_database_name;
```

The GORM will automatically create the users table based on the User struct defined in the main.go file.

## CURL Commands
### Create a new user:

curl -X POST -d '{"name":"John Doe"}' http://localhost:8080/create

### Read all users:

curl http://localhost:8080/read

### Update a user:

curl -X PUT -d '{"id":1,"name":"Updated Name"}' http://localhost:8080/update

### Delete a user:

curl -X DELETE -d '{"id":1}' http://localhost:8080/delete
