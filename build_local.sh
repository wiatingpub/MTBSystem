#!/bin/bash
ProjectName="mtbsystem"


if [ $1 == "all" ]; then
    for srv in `ls src`; do
        if [[ ${srv:0-4} == "-srv" ]]; then
            if [[ ${srv} != "api-srv" ]]; then
                echo "开始更新$srv"
                GOROOT=/data/services/go GOBIN=/data/goapp/$ProjectName/bin GOPATH=`pwd`:`pwd`/vendor /data/services/go/bin/go install $srv && sudo supervisorctl restart $srv:*
            fi
        fi
    done
else
    for srv in "$@"
    do
        if [[ "${srv:0-4}" != "-srv" ]]; then
            srv="${srv}-srv"
        fi
        echo "开始更新$srv"
        GOROOT=/data/services/go GOBIN=/data/goapp/$ProjectName/bin GOPATH=`pwd`:`pwd`/vendor /data/services/go/bin/go install $srv && sudo supervisorctl restart $srv:*
    done
fi