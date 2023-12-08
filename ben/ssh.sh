cd ./firecracker-containerd
git checkout firecracker-v1.4.1-vhive-integration-debug
cd ./tools/image-builder
ssh -vvv -i ./firecracker_rsa root@172.16.0.1