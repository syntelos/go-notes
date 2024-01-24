#!/bin/bash

rm -f /tmp/list.txt

for each in name path link
do
    if go run source_table.go list ${each} >> /tmp/list.txt
    then
	echo "# /tmp/list.txt"
	echo >> /tmp/list.txt
    else
	exit 1
    fi
done

if go run source_table.go enumerate path > /tmp/path.txt
then
    echo "# /tmp/path.txt"
else
    exit 1
fi

if go run source_table.go enumerate link > /tmp/link.txt
then
    echo "# /tmp/link.txt"
else
    exit 1
fi
