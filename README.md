# go-api-boilerplate
 
## Overview
A web application boilerplate built with go and clean architecture.

## Architecture
```bash
go-api-boilerplate
├── config                   //application config
├── docker                   //Dockerfile
│   ├── api
│   │   └── Dockerfile
│   └── db
│       ├── Dockerfile
│       ├── conf.d           //mysql config
│       └── initdb.d         //seed Data
└── src
    └── api
        ├── Gopkg.lock
        ├── Gopkg.toml       //dependencies
        ├── config
        │   ├── config.go    //load config file
        ├── domain           //Entities
        │   └── user.go      
        ├── infrastracture   //Frameworks & Drivers
        │   ├── db
        │   ├── env
        │   ├── http
        │   ├── log
        │   └── router
        ├── interfaces       //Interface
        │   ├── context.go
        │   ├── error.go
        │   ├── sqlhandler.go
        │   └── user
        ├── main.go
        ├── usecases         //Use cases
        │   ├── logger.go
        │   └── user
        └── vendor
└── Makefile                 // docker-compose operation
└── docker-compose.yml       // docker operation

```


## Requirement
 - Docker
 - Go

## Usage

### Start
```bash
make start
```

### Stop
```bash
make stop
```

### List process
```bash
make ps
```

### view api log
```bash
make logs
```

### Run a command in a running container
```bash
make exec-api
```

## LICENCE
MIT License

## References
- https://blog.tai2.net/the_clean_architecture.html
- https://qiita.com/hirotakan/items/698c1f5773a3cca6193e
- https://gist.github.com/mpppk/609d592f25cab9312654b39f1b357c60

