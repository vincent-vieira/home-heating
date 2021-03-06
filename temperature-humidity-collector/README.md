# Temperature & humidity collector/publisher

## About the software
This software is written in Go and is designed to periodically read values from a Unix socket exposing measurements. All of them are then forwarded to anyone listening for them through a [Socket.io](http://socket.io/) server, allowing data streaming to anyone connected.

## Launch options
- `listen`
    - The port to listen on for Websocket connections
    - Not mandatory, defaults to *5000*
- `refresh`
    - The time interval between two temperature/humidity measurements, in seconds
    - Not mandatory, defaults to *60*
- `gpio`
   - The GPIO pin number where the sensor is connected
   - Not mandatory, defaults to GPIO pin *4*
   
## Build information 
The project is following the recommended Go layout for a project. Please just make sure that this folder is accessible from the *GOPATH*.

### External dependencies
This project requires the *bcm2835* library. Make sure it's installed before compiling.

## Containerization & delivery methods
The project also includes a Dockerfile, embedding a very lightweight Linux runtime, with the go runtime installed. 
Building the image for your target architecture will also compile the project.

*You can pass directory the program's arguments to the Docker container. For instance, this command works :*
`docker run home-heating/collector -listen 6000`