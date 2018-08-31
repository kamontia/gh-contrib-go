#!/bin/bash
curl -v "https://api.github.com/search/issues?q=type:pr+in:body+is:merged+merged:2018-08-24..2018-08-30+author:chaspy"
