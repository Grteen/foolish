#! /bin/bash

./getpid

cat ./pid/user.pid | xargs kill -9
cat ./pid/artical.pid | xargs kill -9
cat ./pid/search.pid | xargs kill -9
cat ./pid/main.pid | xargs kill -9