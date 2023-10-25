#!/bin/sh -e
set -x

# Sort imports one per line, so autoflake can remove unused imports
isort --force-single-line-imports app
autoflake --remove-all-unused-imports --recursive --remove-unused-variables --in-place app  --exclude=__init__.py
black app
isort app
mypy app
