# Example api with Gin

### Motivation
### Configuration Manage
#### ENV Manage

- Default ENV Configuration Manage from `.env`. sample file `.env.example`


#### Server Configuration
- Use [Gin](https://github.com/gin-gonic/gin) Web Framework

#### Database Configuration
- Use [GORM](https://github.com/go-gorm/gorm) as an ORM
- Use database `MASTER_DB_HOST` value set as `localhost` for local development, and use `postgres_db` for docker development 

### Installation
#### Local Setup Instruction
- Make sure install the latest version of docker and docker-compose
- Pull the necessary tools
```
make init-tools
```

#### Run Application with live reload
```
make dev
```
#### Generate mocks
```
make gen-mocks
```

#### Seed data
Let's seed some data to test and run the application.
Please run the SQL file from `test-data/seed.sql` 
<!-- TODO: Create seed command -->

### Project Structure
<pre>├── <font color="#3465A4"><b>config</b></font>
│   ├── config.go
│   ├── db.go
│   └── server.go
├── <font color="#3465A4"><b>cmd/</b></font>: Contains the main file to start the application
│   └── <font color="#3465A4"><b>server</b></font>
│      └── <font color="#3465A4"><b>routers</b></font>: Contains the routes, endpoints, and middlewares
├── <font color="#3465A4"><b>internal</b></font>
│   └── <font color="#3465A4"><b>models</b></font>: contain models for the application
│   └── <font color="#3465A4"><b>handlers</b></font>: Parse/validate request params/body. Define the response structure
│   └── <font color="#3465A4"><b>controller</b></font>: Business logic layer
│   └── <font color="#3465A4"><b>repository</b></font>: Data access 
layer
├── docker-compose-dev.yml
├── docker-compose-prod.yml
├── Dockerfile
├── Dockerfile-dev
├── go.mod
├── go.sum
├── <font color="#3465A4"><b>infra</b></font>
│   └── <font color="#3465A4"><b>logger</b></font>
│       └── logger.go
├── LICENSE
├── Makefile
├── README.md
</pre>

### Reference
- Inspired from the [gin-boilerplate-examples](https://github.com/akmamun/gin-boilerplate-examples)


### Useful Commands
- `make dev`: make dev for development work
- `make build`: make build container
- `make production`: docker production build and up
- `clean`: clean for all clear docker images

### Use Packages
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)
