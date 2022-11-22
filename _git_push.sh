#!/bin/bash

VERSION="0.0.2"



git add .
git commit -m "DEBUG $VERSION"
# create tag
git tag "v$VERSION"

# git remote add gitee git@gitee.com:lyhuilin/openapi_setting.git
# git remote add github git@github.com:clin003/openapi_setting.git
git push -u github main
git push -u gitee main
# git push
# push commit and tag to remote
git push --tags  -u github main
git push --tags  -u gitee main

