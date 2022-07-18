# Estuary Metrics Core Module

Importable core module. This has all the generated data models which is then used by the rest of the components e.i rest api notification and cli.

## To generate the default models

Prepare the .env variables with the following DB connection
```
DB_NAME=DBNAME
DB_HOST=DBHOST
DB_USER=DBUSER
DB_PASS=DBPASS
DB_PORT=DBPORT
```

Run generate.sh
```
 ./generate.sh
```

This generates the models and the default querying functions.

## Usage

```
go get github.com/application-research/estuary-metrics/core
```
Initialize the module
```
//  initialize your database connection (estuary) - readonly
metricsCore,err := core.Init(DB,gorm)
metricsCore.GetRetrievalDealSuccessRate()
```

## Extending
This module comes with a model package. Anyone can build a new metric data that can be exposed or consumed by clients.
```
metricsCore,err := core.Init(DB,gorm)
u := query.Use(DB).Content
content, err := u.WithContext(ctx).Count() // get number of contents
```