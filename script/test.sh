#!/bin/bash
#
# Test script for gh-contrib
#
# Usage: ./test.sh

./gh-contrib -u chaspy -s 180824 -e 180830
ret="$?"

if [[ "$ret" == "0" ]]; then
  echo "Test Passed"
else
  echo "Test Failed"
fi
