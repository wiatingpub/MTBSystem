#!/bin/sh
Container="mtbsystem"
ProjectName="mtbsystem"

case "$1" in
    build)
        cd dockerbase
        docker build -t ricoder/go-env .
        cd ..
    ;;
    run)
        docker run --name=$Container -p 13306:3306 -p 16379:6379 -p 18500:8500 -p 19999:9999 -p 15324:5324 -p 18082:8082 -p 18083:8083 -p 18084:8084 -p 18085:8085 -p 18090:8090 -d -v `pwd`:/data/deploy/$ProjectName ricoder/go-env
    ;;
    init)
        docker exec $Container bash /data/deploy/$ProjectName/dockerbase/init.sh $2
    ;;
    start)
        docker start $Container
    ;;
    stop)
        docker stop $Container
    ;;
    remove)
        docker stop $Container
        docker rm $Container
    ;;
    login)
        docker exec -it $Container /bin/bash
    ;;
    info)
        docker ps -a -f name=$Container
    ;;
    port)
        docker port $Container
    ;;
    *)
        echo "Usage: build|run|init[chmod,conf]|start|stop|remove|login|info|port"
        exit 3
    ;;
esac
