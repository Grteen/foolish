#! /bin/bash

cd /root/BE/script/bin
./pid

cd /root/BE/script
cat ./pid/user.pid | xargs kill -9
cat ./pid/artical.pid | xargs kill -9
cat ./pid/search.pid | xargs kill -9
cat ./pid/notify.pid | xargs kill -9
cat ./pid/action.pid | xargs kill -9
cat ./pid/comment.pid | xargs kill -9
cat ./pid/main.pid | xargs kill -9