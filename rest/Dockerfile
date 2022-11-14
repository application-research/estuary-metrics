FROM golang:1.18-stretch AS builder
USER root
RUN apt-get update && \
    apt-get install -y wget jq hwloc ocl-icd-opencl-dev git libhwloc-dev pkg-config make && \
    apt-get install -y cargo \

WORKDIR /app/
COPY . .
RUN make all
RUN ./metrics-rest-api

