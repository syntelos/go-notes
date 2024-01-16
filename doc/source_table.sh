#!/bin/bash

rm -f /tmp/list.txt

for each in name path link
do
    if ! go run source_table.go list ${each} >> /tmp/list.txt
    then
	exit 1
    fi
done

if ! go run source_table.go enumerate path > /tmp/path.txt
then
    exit 1
fi

if ! go run source_table.go enumerate link > /tmp/link.txt
then
    exit 1
fi
