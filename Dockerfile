FROM centurylink/ca-certs

ADD ifttt_ipchange /
ENTRYPOINT ["/ifttt_ipchange"]
