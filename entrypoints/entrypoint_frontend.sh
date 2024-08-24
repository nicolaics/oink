#!/bin/bash
echo "var BACKEND_ROOT = '$BACKEND_HOST:$BACKEND_PORT';" > game.js
python3 -m http.server 1337
