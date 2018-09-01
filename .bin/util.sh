#!/bin/bash
# if [ "$BASH" != "/bin/bash" ] ; then echo "Bash Only"; exit 1; fi

function print_task_name() {
    echo ""
    echo "---"
    echo "[$(date +"%T")] > $1"
}

function echo_and_eval() {
    echo "[$(date +"%T")] $ $1"
    eval "${1}"
}
