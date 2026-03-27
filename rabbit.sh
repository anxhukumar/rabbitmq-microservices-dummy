#!/bin/bash

start_or_run () {
    docker inspect e-commerce-rabbitmq > /dev/null 2>&1

    if [ $? -eq 0 ]; then
        echo "Starting E-commerce RabbitMQ container..."
        docker start e-commerce-rabbitmq
    else
        echo "E-commerce RabbitMQ container not found, creating a new one..."
        docker run -d --name e-commerce-rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management
    fi
}

case "$1" in
    start)
        start_or_run
        ;;
    stop)
        echo "Stopping E-commerce RabbitMQ container..."
        docker stop e-commerce-rabbitmq
        ;;
    logs)
        echo "Fetching logs for E-commerce RabbitMQ container..."
        docker logs -f e-commerce-rabbitmq
        ;;
    *)
        echo "Usage: $0 {start|stop|logs}"
        exit 1
esac