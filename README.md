[![Build Status](https://travis-ci.org/roundpartner/tempus.svg?branch=master)](https://travis-ci.org/roundpartner/tempus)

# Tempus
Micro Service for generating time based keys

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
