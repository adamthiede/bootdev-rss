#!/usr/bin/env bash
curl -X POST "http://localhost:8080/v1/users" -H "Content-Type: application/json" -d $@
