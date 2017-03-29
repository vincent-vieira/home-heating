FROM easypi/alpine-arm
MAINTAINER Vincent Vieira <vincent.vieira@supinfo.com>

# curl for debugging, nginx for... you know.
RUN apk add --no-cache nginx curl
RUN adduser -D -u 1000 -g 'www' www

# Forward request and error logs to docker logs console
RUN ln -sf /dev/stdout /var/log/nginx/access.log
RUN ln -sf /dev/stderr /var/log/nginx/error.log

# Permissions-related stuff
RUN chown -R www:www /var/lib/nginx && chown -R www:www /var/www
RUN mkdir -p /var/www/web-gateway /run/nginx

COPY config/nginx.conf /etc/nginx/nginx.conf
COPY config/nginx.vh.default.conf /etc/nginx/conf.d/default.conf

# Launching
EXPOSE 80 443
RUN rc-service nginx start && rc-update add nginx default

ENTRYPOINT ["/bin/bash"]