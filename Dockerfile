FROM golang:latest

#RUN DEBIAN_FRONTEND=noninteractive apt update && apt install build-essential make wget libncurses-dev bison flex libssl-dev libelf-dev bc -y

WORKDIR "/src"

ENTRYPOINT [ "make" ]
