#!/bin/bash
#
# Test script for gh-contrib
#
# Usage: ./test.sh

./gh-contrib -u kamontia -s "2018-08-24" -e "2018-08-30"
ret="$?"

if [[ "$ret" == "0" ]]; then
  echo "Test Passed"
else
  echo "Test Failed"
fi
