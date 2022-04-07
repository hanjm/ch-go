#!/bin/bash

# Download static build .tgz by version.
# Example: _hack/dl.sh 22.3.2.2-lts

URL=$(curl -L -s "https://api.github.com/repos/ClickHouse/ClickHouse/releases/tags/v${1}" | grep -o -P "https:\/\/.*clickhouse-common-static-\d.*\.tgz")
wget -O /tmp/static.tgz "${URL}"