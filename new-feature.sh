#!/bin/bash
FEATURE=$1

mkdir -p internal/domain/$FEATURE
mkdir -p internal/application/$FEATURE
mkdir -p internal/adapters/inbound/http

touch \
internal/domain/$FEATURE/entity.go \
internal/domain/$FEATURE/repository.go \
internal/application/$FEATURE/create.go

