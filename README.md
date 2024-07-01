# webc - One file command line WEB server
![webc builder](https://github.com/midstar/webc/actions/workflows/build.yml/badge.svg)

webc is a mini command line WEB server. It consists of one executable 
and no external dependencies.

All majort platforms supported such as Windows, Linux (x86 and ARM)
and Mac.

No installation required, just download the suitable executable 
[here on GitHub](https://github.com/midstar/webc/releases). Or build
the software yourself according to the instructions below.

## Usage

Copy the webc to a directory of your choice and add it to your PATH
environment variable.

    Usage: webc [options] [<path>]

    <path> is the directory where your web files are located.
    Default is current directory.

    Supported options:
    -c string
            TLS certificate file (default "cert.pem")
    -k string
            TLS key file (default "key.pem")
    -p int
            Network port to listen to (default 8080)
    -s	Use secure connection (TLS/HTTPS)
    -v	Display version

OpenSSL can be used to generate the public and private key required for TLS/HTTPS:

    openssl genrsa -out key.pem 2048
    openssl req -new -x509 -sha256 -key key.pem -out cert.pem -days 3650

## Build from source (any platform)

To build from source on any platform you need to:

* Install Golang 
* Set the GOPATH environment variable

Then run:

    go get github.com/midstar/webc
    go install github.com/midstar/webc


## Author and license

This application is written by Joel Midstjärna and is licensed under the MIT License.