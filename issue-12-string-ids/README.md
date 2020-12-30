This is a reproduction for string id's for to test and resolve this issue: https://github.com/web-ridge/gqlgen-sqlboiler/issues/12
Not a real example!

## Generated files in this example

- graphql_models/\*
- models/\*
- helpers/\*
- resolver.go

## Steps to make this project

1. Start MariaDB container with social network.sql  
   `docker-compose up -d`

2. Generate /models/\* files  
   `sqlboiler mysql --no-back-referencing -d`

   --no-back-referencing if you don't set it your program will crash with recursive converts!

3. Generate schema.grapql, converts (helpers/*) and resolver.go
   `go run convert_plugin.go`

4. ...Let's go!  
   `go run server.go resolver.go database.go`
