FROM ubuntu:latest

MAINTAINER Tige Phillips <tige@ignw.io>

RUN apt-get update

# Install some pre-requisite - Ubuntu
#####################################
RUN apt-get -y install git make wget

# Install GO - Ubuntu
#####################
RUN mkdir /root/go/
RUN apt-get install -y golang
ENV GOPATH "/root/go"
ENV PATH "$PATH:/usr/local/go/bin:$GOPATH/bin"

# Setup Terraform - golang
##########################
RUN apt-get update
RUN apt-get install unzip
WORKDIR /tmp
RUN wget https://releases.hashicorp.com/terraform/0.11.11/terraform_0.11.11_linux_amd64.zip
RUN unzip terraform_0.11.11_linux_amd64.zip
RUN mv terraform /usr/local/sbin
RUN rm terraform_0.11.11_linux_amd64.zip

# IGNW Cisco ACI Go SDK and Terraform provider install
######################################################
RUN mkdir -p $GOPATH/src/github.com/ignw
WORKDIR $GOPATH/src/github.com/ignw

# Clone everything
RUN git clone https://github.com/IGNW/cisco-aci-go-sdk.git
RUN git clone https://github.com/IGNW/terraform-provider-cisco-aci.git

# Install the SDK
WORKDIR $GOPATH/src/github.com/ignw/cisco-aci-go-sdk
RUN make install

# Install the Terraform provider
WORKDIR $GOPATH/src/github.com/ignw/terraform-provider-cisco-aci
RUN make tools
RUN dep init
RUN make deps
RUN make fmt
RUN make build

#RUN mkdir -p examples/example_1/terraform.d/plugins/linux_amd64
#RUN cp /root/go/bin/terraform-provider-aci examples/example_1/terraform.d/plugins/linux_amd64
