#!/bin/sh
goose postgres 'postgres://postgres:@localhost:5432/blogator' $1
