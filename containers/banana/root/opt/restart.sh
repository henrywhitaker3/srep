#!/bin/bash

# Kill the currently running redis process
kill -s 9 $(ps -aux | grep redis-server | head -n 1 | awk '{print $2}')

# Start a new one
/opt/run/run.sh &
