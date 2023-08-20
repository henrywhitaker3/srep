#!/bin/bash

mode=$(cat /etc/redis/redis.conf | grep protected-mode | cut -d' ' -f 2)

if [ $mode == "no" ]; then
    echo -n "OK"
else
    echo -n "NO"
fi
