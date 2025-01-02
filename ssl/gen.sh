#!/bin/bash

set -e

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

docker run --rm -it -p 80:80 -v "$SCRIPT_DIR:/etc/letsencrypt" -v "$SCRIPT_DIR:/var/lib/letsencrypt" certbot/certbot certonly --standalone -m juanbiabdon@gmail.com --agree-tos -d jbabdon.dev -d www.jbabdon.dev
