#!/bin/bash -eux

readonly base='http://ejje.weblio.jp/content/'

cat wordList.txt | while read word; do
    target=`echo $base$word | tr -d '\r\n'`
    echo "$word\t`curl -Lso- $target | grep -oP "content-explanation>\K(.*?)(?=<)"`" >> out/out.txt
done
