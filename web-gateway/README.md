# NGINX webproxy

## Purpose of this module
This module is responsible for proxying all incoming webtraffic to the respective instances up all around this project.

## Security
TODO 

## Proxy rules
2 proxy rules are configured :
- `/portainer` : 
    - Proxies traffic to the web interface of the [Portainer](http://portainer.io/) container declared in the `docker-compose.yml` file in the parent directory.
    - Used for Docker administration tasks
- `/rethink` :
    - Proxies traffic to the web interface of the [RethinkDB](https://rethinkdb.com/) container declared in the `docker-compose.yml` file in the parent directory.
    - Used for administration tasks
