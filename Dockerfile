FROM gcr.io/mitac-cust-gcp-1/base-onion:latest 

CMD ["/go/bin/onion","serve"]
EXPOSE 8080
