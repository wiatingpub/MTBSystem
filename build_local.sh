#!/bin/bash
ProjectName="mtbsystem"

# 将supervisor的conf文件复制到指定目录并且重启supervisorctl
cp /data/deploy/$ProjectName/dockerbase/supervisor/*conf /etc/supervisor/conf.d/

if [ $1 == "all" ]; then
    for srv in `ls src`; do
        if [[ ${srv:0-4} == "-srv" ]]; then
            echo "开始更新$srv"
            GOROOT=/data/services/go GOBIN=/data/goapp/$ProjectName/bin GOPATH=`pwd`:`pwd`/vendor /data/services/go/bin/go install $srv && sudo supervisorctl restart $srv:*
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