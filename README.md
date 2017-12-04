[![Build Status](https://travis-ci.org/roundpartner/tempus.svg?branch=master)](https://travis-ci.org/roundpartner/tempus)
# Tempus
Micro Service for generating time based keys
## Abstract
This service provides an end point for generating a new token that can be used to authenticate a user.
# Usage
```bash
curl localhost:7373 -d "{\"user_id\":\"1\",\"scenario\":\"test\"}"
```
```json
{"user_id":"1","scenario":"test","token":"2b73bb3213da2f0627f492b741bdeb491a43d0f392176981e4138924147ca0d7"}
```
```bash
curl localhost:7373/1/2b73bb3213da2f0627f492b741bdeb491a43d0f392176981e4138924147ca0d7
```

```json
{"user_id":"1","scenario":"test","token":"BAa-gqpESzJMy6k-oxiokPk-oJ-sejqzgmdSQ0pHXdz3dwTg9wXImWT3_hBKyhgS"}
```
# Testing
You will need redis to run tests
## Starting Redis
Redis can be run locally in a container with
```bash
docker run --rm --name tempus-redis -p 127.0.0.1:6379:6379 -d redis
```
You can then monitor the server with
```bash
docker exec tempus-redis redis-cli monitor
```
# Benchmarking
```bash
go test -bench=.
```
