#!/bin/bash -eux

readonly base='http://ejje.weblio.jp/content/'

cat wordList.txt | while read word; do
    target=`echo $base$word | tr -d '\r\n'`
    definition=`curl -Lso- $target | grep -oP "content-explanation>\K(.*?)(?=<)"`
    echo "$word\t$definition" >> out/out.txt
done
