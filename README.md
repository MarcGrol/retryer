# Retying http forwarder

This HTTP-service will as a persistent and retrying queue.<br/>
Upon receipt of a HTTP POST-request, the service will asynchronously forward the received HTTP request to a remote host.<br/>
When the remote host does not return a success, the request will be retried untill success or 
untill the retry scheme is exhausted.<br/>
The remote host is indicated by:
- the HTTP query parameeter "HostToForwardTo" or
- the HTTP-request-header "X-HostToForwardTo"
   
## deploy
  
    gcloud auth login
    gcloud config set project retryer
    gcloud app deploy app.yaml queue.yaml --quiet --version 3
    
## test

Example to test the interaction:

    curl -vvv \
        -X POST \
        --data "This is expected to be sent back as part of response body." \
        "https://retryer-dot-retryer.appspot.com/post?HostToForwardTo=https://postman-echo.com"   

