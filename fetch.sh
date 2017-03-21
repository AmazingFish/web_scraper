#!/bin/bash -eux

readonly base='http://ejje.weblio.jp/content/'

cat wordList.txt | while read word; do
    target=`echo $base$word | tr -d '\r\n'`
    curl -Lso- $target | grep -oP "content-explanation>\K(.*?)(?=<)" 
done
