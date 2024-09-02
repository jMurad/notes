#!/bin/bash

images=$(docker images | grep "<none>" | awk '{print $3}')

if [ -z "$images" ]; then
    echo "No images with tag <none> found."
else
    # Удаляем каждый образ по его ID
    for id in $images; do
        docker rmi $id
    done
    echo "Deleted all images with tag <none>."
fi