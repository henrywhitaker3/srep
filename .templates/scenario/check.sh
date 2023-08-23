#!/bin/bash

rc=$?

if [ $rc -eq 0 ]; then
    echo -n "OK"
else
    echo "NO"
fi
