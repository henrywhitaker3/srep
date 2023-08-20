#!/bin/bash

mode=$(cat /etc/redis/redis.conf | grep protected-mode | cut -d' ' -f 2)

if [ $mode == "no" ]; then
    echo OK
else
    echo NO
fi
