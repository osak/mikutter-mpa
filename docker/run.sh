#!/bin/bash

NETWORK_NAME=mpa_nw
MPA_IMAGE_NAME=mpa
MPA_CONTAINER_NAME=mpa
MONGO_IMAGE_NAME=mongo:3.4
MONGO_CONTAINER_NAME=mpa-mongo

script_dir_rel=$(dirname $0)
base_dir=$(cd "${script_dir_rel}/.."; pwd)

echo "base directory: ${base_dir}"

docker network ls | awk '{print $2}' | grep -e "^${NETWORK_NAME}$" > /dev/null
network_exists=$?
if [ ${network_exists} -ne 0 ]; then
    echo "Network ${mpa_nw} doesn't exist. Creating..."
    docker network create --driver=bridge --subnet=172.18.0.0/24 "${NETWORK_NAME}"
    echo "Done."
fi

mysql_running=$(docker ps -f name=${MONGO_CONTAINER_NAME} -q)
if [ -z "${mysql_running}" ]; then
    echo "Mongo container is not running. Starting..."
    docker kill ${MONGO_CONTAINER_NAME} || true
    docker rm ${MONGO_CONTAINER_NAME} || true
    docker run -d -v ${base_dir}/mongo:/data/db --network=${NETWORK_NAME} --name=${MONGO_CONTAINER_NAME} ${MONGO_IMAGE_NAME}
    echo "Done."
fi

ip_path=".NetworkSettings.Networks.${NETWORK_NAME}.IPAddress"
mongo_ip=$(docker inspect -f "{{${ip_path}}}" ${MONGO_CONTAINER_NAME})
echo "Mongo server's IP is ${mongo_ip}"

docker kill ${MPA_CONTAINER_NAME} || true
docker rm ${MPA_CONTAINER_NAME} || true
bin_dir="${base_dir}/bin"
web_dir="${base_dir}/web-build"
storage_dir="${base_dir}/storage"
if [ ! -e ${storage_dir} ]; then
    mkdir ${storage_dir}
fi
docker run -d -v "${bin_dir}:/app/bin" -v "${web_dir}:/app/web" -v "${storage_dir}:/app/storage" -p 127.0.0.1:3939:8080 --network=${NETWORK_NAME} --name=${MPA_CONTAINER_NAME} ${MPA_IMAGE_NAME} ${mongo_ip}
