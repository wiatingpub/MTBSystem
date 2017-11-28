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
        docker run --name=$Container -p 18089:8084 -p 18505:8500 -d -v `pwd`:/data/deploy/$ProjectName ricoder/go-env
    ;;
    init)
        docker exec $Container bash /data/deploy/$ProjectName/dockerbase/init.sh
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
        echo "Usage: build|run|init|start|stop|remove|login|info|port"
        exit 3
    ;;
esac
