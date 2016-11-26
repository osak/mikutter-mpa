#!/bin/bash

NETWORK_NAME=mpa_nw
MPA_IMAGE_NAME=mpa
MPA_CONTAINER_NAME=mpa
MYSQL_IMAGE_NAME=mysql:5.7
MYSQL_CONTAINER_NAME=mpa-mysql

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

mysql_running=$(docker ps -f name=${MYSQL_CONTAINER_NAME} -q)
if [ -z "${mysql_running}" ]; then
    echo "MySQL container is not running. Starting..."
    docker kill ${MYSQL_CONTAINER_NAME} || true
    docker rm ${MYSQL_CONTAINER_NAME} || true
    docker run -d -v ${base_dir}/mysql:/var/lib/mysql --network=${NETWORK_NAME} --name=${MYSQL_CONTAINER_NAME} -e MYSQL_ALLOW_EMPTY_PASSWORD=yes ${MYSQL_IMAGE_NAME}
    echo "Done."
fi

ip_path=".NetworkSettings.Networks.${NETWORK_NAME}.IPAddress"
mysql_ip=$(docker inspect -f "{{${ip_path}}}" ${MYSQL_CONTAINER_NAME})
echo "MySQL server's IP is ${mysql_ip}"

docker kill ${MPA_CONTAINER_NAME} || true
docker rm ${MPA_CONTAINER_NAME} || true
bin_dir="${base_dir}/bin"
web_dir="${base_dir}/web-build"
docker run -d -v "${bin_dir}:/app/bin" -v "${web_dir}:/app/web" -p 127.0.0.1:3939:8080 --network=${NETWORK_NAME} --name=${MPA_CONTAINER_NAME} ${MPA_IMAGE_NAME} 172.18.0.2
