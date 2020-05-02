# goHTTPForwarder
A Go developed webhook/HTTP forwarder.

## Overview
This is a radically simple go http forwarder. The problem this solves is a service generating webhooks where a HTTP Proxy cannot be specified.

I used this opportunity to write my first Go application. So simple, yet so effective. I wasn't expecting it to be this simple in Go.
## Problem Statement
Problem I was solving was sending webhooks to Microsoft Teams from a secured internet restricted Gitlab instance that we could not configure a host wide http proxy.

### Quick note:
The Go net library auto inherits the 'http_proxy' configuration, setting this environment variable will proxy requests easily.


## Example
The app will listen on all interfaces we can use localhost in the example.
Original webhook URL 
```
https://outlook.office.com/webhook/178c7fe7-0188-449a-95b0-9f3530808115@97e20f09-abc4-442b-91b5-f08a261718a8/IncomingWebhook/403bb4ef667642289dc9a2c32520cc60/d9ee2ca5-8fb6-43a8-8fb9-5b76d487083f
```
To forward it through the http forwarder simply add the FQDN of the forwarder e.g.
```
http://localhost:8000/outlook.office.com/webhook/178c7fe7-0188-449a-95b0-9f3530808115@97e20f09-abc4-442b-91b5-f08a261718a8/IncomingWebhook/403bb4ef667642289dc9a2c32520cc60/d9ee2ca5-8fb6-43a8-8fb9-5b76d487083f
```
