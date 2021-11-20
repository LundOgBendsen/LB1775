#!/bin/sh

[ -d tmp/ ] && sudo rm -rf tmp/*
[ ! -d tmp/ ] && mkdir tmp

cp -pr src tmp/src

#echo 'Run go version'

PRJDIR="/go"
VOLMOUNT=$(pwd)/tmp:${PRJDIR}:rw

CMD=""
if [ "$1" != "" ] ; then
	CMD="go run $*"
fi
sudo docker run -it --rm -v ${VOLMOUNT} -w ${PRJDIR} golang:alpine $CMD
