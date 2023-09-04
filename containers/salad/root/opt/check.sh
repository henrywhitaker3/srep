#!/bin/bash

curl https://echo.localhost --resolve echo.localhost:443:127.0.0.1 -k -I -s | grep -q 200
rc=$?

if [ $rc -eq 0 ]; then
    echo -n "OK"
else
    echo "NO"
fi
