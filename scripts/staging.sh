#!/bin/bash

export SLUG=ghcr.io/awakari/consumer-log
export VERSION=latest
docker tag awakari/consumer-log "${SLUG}":"${VERSION}"
docker push "${SLUG}":"${VERSION}"
