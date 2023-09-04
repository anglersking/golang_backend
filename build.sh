docker rm -f golang_dev
docker build  -t golang_dev . 
docker run -dit --net=host -v $(pwd)/go_project:/go_project --name=golang_dev golang_dev
# docker rm -f uboot_build