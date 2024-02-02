#!/bin/bash
set -x

function usage {
    cat<<EOF>&2
Synopsis

    $0 notes encode

    $0 notes update

Description

    Perform test validation over 'tst' directory.

    When the 'tst/notes' directory is not present, Notes
    Encode is performed before Notes Update to populate the
    directory in a successful performance.

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
		    exit 1
		fi
	    done
	    exit 0
	else
	    2>&1 echo "$0 error from 'wwweb notes upd tst/notes'."
	    exit 1
	fi
    else
	2>&1 echo "$0 error from 'wwweb notes tgt upd tst/notes'."
	exit 1
    fi
}

#
if [ 2 -eq $# ]
then
    case ${1} in
	notes)
	    case ${2} in
		encode)
		    if notes_encode
		    then
			exit 0
		    else
			exit 1
		    fi
		    ;;
		update)
		    if notes_update
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
	;;

	*)
	    usage $0
	    ;;
    esac
else
    usage $0
fi
