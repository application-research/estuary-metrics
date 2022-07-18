export PROJ_PATH=estuary-metrics

all: all
	cd cmd && go build -o metrics-cmd
	cd rest && go build -o metrics-rest-api

