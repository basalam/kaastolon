#!/usr/bin/env bash
version="$1"
target="kaastolon-${version}-linux-amd64"
srcdir="$(pwd)"
make build

mkdir -p "${target}"/bin

for bin in stolon-keeper stolon-sentinel stolon-proxy stolonctl; do
    cp "${srcdir}"/bin/${bin} "${target}"/bin
done

cp "${srcdir}"/README.md "${target}"/README.md

cp -R "${srcdir}"/doc "${target}"/doc
cp -R "${srcdir}"/examples "${target}"/examples
rm -rf "${target}"/examples/kubernetes/image/docker/bin
rm -rf "${target}"/examples/docker/bin

tar cvfz ${target}.tar.gz ${target}/
echo "Wrote release/${target}.tar.gz"