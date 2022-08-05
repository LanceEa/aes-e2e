#!/bin/bash
set -euE -o pipefail

# capture Go psuedo version
version="v3.0.1-0.20220805143754-5480a49f4150"

# output the image tag version for use
echo "${version:1}"

