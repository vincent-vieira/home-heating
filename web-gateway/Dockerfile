FROM easypi/alpine-arm
MAINTAINER Vincent Vieira <vincent.vieira@carbon-it.fr>

# curl for debugging
RUN apk add --no-cache nginx curl

# forward request and error logs to docker log collector
RUN ln -sf /dev/stdout /var/log/nginx/access.log
RUN ln -sf /dev/stderr /var/log/nginx/error.log

RUN mkdir -p /run/nginx
COPY config/nginx.conf /etc/nginx/nginx.conf
COPY config/nginx.vh.default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
