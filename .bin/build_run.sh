#!/bin/bash
# if [ "$BASH" != "/bin/bash" ] ; then echo "Bash Only"; exit 1; fi

ENV=$1
GOPATH=$(echo $GOPATH | cut -d ':' -f 1)
PROJECTPATH="$GOPATH/src/github.com/therudite/api"
. ${PROJECTPATH}/.bin/util.sh

print_task_name "change dir to project path"
echo_and_eval "cd ${PROJECTPATH}"

print_task_name "Fix code style"
echo_and_eval "go fmt ${PROJECTPATH}..."

print_task_name "Load config file for develop"
echo_and_eval "rm ${PROJECTPATH}/config/config.ini"

print_task_name "Environment ${ENV}"
case ${ENV} in
  develop)
    echo_and_eval "cp ${PROJECTPATH}/config/dev.config.ini ${PROJECTPATH}/config/config.ini"
    ;;
  production)
    echo_and_eval "cp ${PROJECTPATH}/config/prod.config.ini ${PROJECTPATH}/config/config.ini"
    ;;
  *)
    echo_and_eval "cp ${PROJECTPATH}/config/dev.config.ini ${PROJECTPATH}/config/config.ini"
    ;;
esac

print_task_name "sync packages"
echo_and_eval "dep ensure"

print_task_name "Build project"
echo_and_eval "go build ./"

print_task_name "Run"
echo_and_eval "./api"
