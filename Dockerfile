FROM alpine:3.7
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ARG env
ARG config
ARG version
ENV env=${env}
ENV config=${config}
ENV version=${version}
CMD ["/app/app-exe"]