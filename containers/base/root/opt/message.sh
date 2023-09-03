#!/bin/bash

echo CiAgICBfX19fX19fX19fXyBfX19fX19fX19fXwogICAgLyAgX19ffCBfX18gXCAgX19ffCBfX18gXAogICAgXCBgLS0ufCB8Xy8gLyB8X18gfCB8Xy8gLwogICAgYC0tLiBcICAgIC98ICBfX3x8ICBfXy8KICAgIC9cX18vIC8gfFwgXHwgfF9fX3wgfAogICAgXF9fX18vXF98IFxfXF9fX18vXF98CgoK | base64 -d

if [ -f "/opt/metadata.json" ]; then
    description=$(cat /opt/metadata.json | jq -r '.description')
    printf "$description\n\n"
fi
