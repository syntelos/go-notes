#!/bin/bash

function usage {
    cat<<EOF>&2
Synopsis

    $0 all

    $0 notes

Description

    Perform test validation over files in 'tst' directory.

EOF
    exit 1
}

function notes_encode {

    if targets=$(wwweb notes tgt enc tst/notes tst/txt) &&[ -n "${targets}" ]
    then
	if wwweb notes enc tst/notes tst/txt
	then
	    for tgt in ${targets}
	    do
		if [ -f "${tgt}" ] # [TODO] (vector)
		then

		    echo "# ${tgt}"
		    cat -n ${tgt}
		else
		    2>&1 echo "$0 error: missing '${tgt}'."
		    return 1
		fi
	    done
	    return 0
	else
	    2>&1 echo "$0 error from 'wwweb notes enc tst/notes tst/txt'."
	    return 1
	fi
    else
	2>&1 echo "$0 error from 'wwweb notes tgt enc tst/notes tst/txt'."
	return 1
    fi
}

function notes_update {

    if targets=$(wwweb notes tgt upd tst/notes) &&[ -n "${targets}" ]
    then
	if wwweb notes upd tst/notes
	then
	    for tgt in ${targets}
	    do
		if [ -f "${tgt}" ] # [TODO] (vector)
		then

		    echo "# ${tgt}"
		    cat -n ${tgt}
		else
		    2>&1 echo "$0 error: missing '${tgt}'."
		    return 1
		fi
	    done
	    return 0
	else
	    2>&1 echo "$0 error from 'wwweb notes upd tst/notes'."
	    return 1
	fi
    else
	2>&1 echo "$0 error from 'wwweb notes tgt upd tst/notes'."
	return 1
    fi
}

function test_notes {
    rm -rf tst/notes

    if notes_encode
    then
	if notes_update
	then
	    exit 0
	else
	    exit 1
	fi
    else
	exit 1
    fi
}

#
if [ 1 -eq $# ]
then
    case ${1} in
	notes|all)
	    if test_notes
	    then
		exit 0
	    else
		exit 1
	    fi
	    ;;

	*)
	    usage $0
	    ;;
    esac
else
    usage $0
fi
