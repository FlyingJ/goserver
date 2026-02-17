# start with lightweight linux goodness
FROM debian:stable-slim
# drop in a simple go http server
COPY goserver /bin/goserver
# set necessary environment variables
ENV PORT="8991"
# do the thing
CMD ["/bin/goserver"]
