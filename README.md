# TCP Server #

## Introduction ##

This is a basic TCP server that handles HTTP requests and responses. This reads and responds to requests that adhere to the HTTP.

## Behind the scenes ##

* The webserver is written in go using net package from the standard library. 

* Splits the request string to fetch the URI and the METHOD.

* Responds according to the URI and METHOD obtained from the request.

## What did I Learn ##

* A better understanding of TCP and HTTP protocols.

* Structure of HTTP requests.

* Learned about IETF and RFC 7230.

* Worked with the net package of the Go standard library.
