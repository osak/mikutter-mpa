#!/bin/bash

script_dir_rel=$(dirname $0)
base_dir=$(cd "${script_dir_rel}/.."; pwd)

echo "base directory: ${base_dir}"

bin_dir="${base_dir}/bin"
web_dir="${base_dir}/web-build"

docker run -d -v "${bin_dir}:/app/bin" -v "${web_dir}:/app/web" -p 127.0.0.1:3939:8080 --network=mpa_nw mpa 172.18.0.2
