#!/bin/bash
# push.sh

VERSION=$1
if [ -z "$VERSION" ]; then
    echo "Error: No version specified."
    exit 1
fi

docker build -t account-app-backend:$VERSION .
docker tag account-app-backend:$VERSION gsemir/account-app-backend:$VERSION
docker push gsemir/account-app-backend:$VERSION