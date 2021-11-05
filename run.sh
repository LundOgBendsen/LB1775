echo 'Run go version'

PRJDIR="/go"
VOLMOUNT=$(pwd):${PRJDIR}:rw

sudo docker run -it --rm -v ${VOLMOUNT} -w ${PRJDIR} golang:alpine
