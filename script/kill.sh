#! /bin/bash

cd /root/foolish/script/bin
./pid

cd /root/foolish/script
cat ./pid/user.pid | xargs kill -9
cat ./pid/artical.pid | xargs kill -9
cat ./pid/search.pid | xargs kill -9
cat ./pid/notify.pid | xargs kill -9
cat ./pid/action.pid | xargs kill -9
cat ./pid/comment.pid | xargs kill -9
cat ./pid/main.pid | xargs kill -9