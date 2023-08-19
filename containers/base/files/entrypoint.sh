#!/bin/bash

for file in /opt/run/*; do
    $file &
done

tail -f /dev/null
