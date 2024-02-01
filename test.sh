#!/bin/bash

if targets=$(wwweb notes tgt enc tst/notes tst/txt) &&[ -n "${targets}" ]
then
    if wwweb notes enc tst/notes tst/txt
    then
	for tgt in ${targets}
	do
	    if [ -f "${tgt}" ]
	    then
		echo "# ${tgt}"
		cat -n ${tgt}
	    else
		2>&1 echo "$0 error: missing '${tgt}'."
		exit 1
	    fi
	done
	exit 0
    else
	2>&1 echo "$0 error from 'wwweb notes enc tst/notes tst/txt'."
	exit 1
    fi
else
    2>&1 echo "$0 error from 'wwweb notes tgt enc tst/notes tst/txt'."
    exit 1
fi
