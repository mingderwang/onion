FROM asia.gcr.io/winter-wonder-647/base-onion:latest 

CMD ["/go/bin/onion","serve"]
EXPOSE 8080
