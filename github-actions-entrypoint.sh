#!/bin/bash

set -euo pipefail

echo "bc compile --file "${INPUT_FILE}" $INPUT_ARGS"

bc compile --file "${INPUT_FILE}" $INPUT_ARGS

