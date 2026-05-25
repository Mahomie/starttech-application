#!/bin/bash

URL="http://localhost"

STATUS=$(curl -o /dev/null -s -w "%{http_code}" $URL)

if [ "$STATUS" -eq 200 ]; then
  echo "Application healthy."
  exit 0
else
  echo "Application unhealthy. Status code: $STATUS"
  exit 1
fi
