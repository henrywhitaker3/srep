#!/usr/bin/python3

import json
import glob

combined = []

for file in glob.glob("containers/**/metadata.json"):
    if file == "containers/base/metadata.json":
        continue

    with open(file) as opened:
        data = json.load(opened)
        combined.append(data)

output = json.dumps(combined, indent=4)
print(output)
