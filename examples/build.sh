#!/bin/bash
set -e

if [ "$(id -u)" != "0" ]; then
	echo "Sorry, you are not root."
	exit 1
fi

dir=$( dirname $0 )

buildAci() {
    echo -e "\n\n\033[0;32mBuilding aci : ${1}\033[0m\n\n"
    sleep 1
    dgr -W ${1} -L debug clean install
}

# base
buildAci ${dir}/aci-base

# none
buildAci ${dir}/none/aci-libc
buildAci ${dir}/none/aci-grafana
buildAci ${dir}/none/aci-prometheus
buildAci ${dir}/none/aci-rkt

# debian
buildAci ${dir}/debian/aci-debian
buildAci ${dir}/debian/aci-debian-cassandra
buildAci ${dir}/debian/aci-debian-openjdk7-jre
buildAci ${dir}/debian/aci-debian-prometheus-jmx-exporter

# alpine
buildAci ${dir}/alpine/aci-alpine-base
buildAci ${dir}/alpine/aci-alpine-nginx

# gentoo
buildAci ${dir}/gentoo/aci-gentoo-stage4
#buildAci ${dir}/gentoo/aci-gentoo-lighttpd # TODO build is too long for travis

# pod
buildAci ${dir}/pod/pod-cassandra

echo -e "\n\033[0;32mEverything looks good !\033[0m\n"
