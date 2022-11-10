# Estuary Metrics RESTAPI Module

This is a standalone api module that runs on its own web server. 

It uses the estuary authentication api (https://github.com/application-research/estuary-auth) to authenticate request.

Similar to estuary, it needs a Bearer token with an Admin level permission to access the endpoints.

# Access Key
In order to access the API, you need to have an Estuary API key. Please request an API key [here](https://docs.estuary.tech/get-invite-key)

# Database Access
Create a .env file on the root directory and fill in with Estuary DB details
```
DB_NAME=DBNAME
DB_HOST=DBHOST
DB_USER=DBUSER
DB_PASS=DBPASS
DB_PORT=DBPORT

CACHE_DEFAULT_TTL=10
```

## Build 
```
go build -o estuary-metrics-api api.go
```

# Run
```
./estuary-metrics-api
```
This opens up a web server on port [3030](http://localhost:3030)

# View the DDL to view all endpoints
```
curl -H "Authorization: Bearer <token>" http://localhost:3030/ddl
```