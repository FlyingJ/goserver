# start with lightweight linux goodness
FROM debian:stable-slim
# drop in a simple go http server
COPY goserver /bin/goserver
# configure the runtime environment
ENV GOSERVER_PORT="8080"
ENV GOSERVER_ROOT="/var/run/goserver/public"
# do the thing
CMD ["/bin/goserver"]
