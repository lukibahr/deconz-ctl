# deconz-ctl

deconz-ctl is a command line utility to control a deconz device via the [deconz REST API](https://dresden-elektronik.github.io/deconz-rest-doc/).
Using deconz-ctl assumes you have properly set up a deconz device and installed lights and switches. Setting up the authentication for the REST API is documented [here](https://dresden-elektronik.github.io/deconz-rest-doc/getting_started/).

> Note: This project is in early development.

## Installing the client

You can either:

- download pre-built binaries from the releases page and move it to your `$PATH` variable
- build the client by yourself using the `Taskfile.yaml`
- use the supplied docker image
- or install the client using `go install github.com/lukibahr/deconz-ctl` if you have golang installed

## Using the client (the docker way)

Using docker, first login to the ghcr.io registry:

`docker login ghcr.io --username=$GITHUB_USER --password=$GITHUB_REGISTRY_ACCESS_TOKEN`

and run:

`docker run -it --rm -e DECONZ_API_KEY=$DECONZ_API_KEY -e DECONZ_API_URL=$DECONZ_API_URL ghcr.io/lukibahr/deconz-ctl:1.0.4 switch-on -d 5`

The client is build with cobra. You can either use command line parameters or environment variables:

```zsh
export DECONZ_API_KEY=
export DECONZ_API_URL=http://<api-url>:port
export DECONZ_DEVICE=#optional, can be supplied with -d option because this option might be more dynamic
```

## Currently implemented functions

- switch on and off a light by device id
- show version

### Future functions

- get configuration: <http://$DECONZ_GATEWAY/api/$DECONZ_API_KEY/config>
- list all lights: <http://$DECONZ_GATEWAY.ip/api/$DECONZ_API_KEY/lights> with table output
- list lights: <http://$DECONZ_GATEWAY/api/$DECONZ_API_KEY/lights/$id>

This bit of software was built while watching cows :cow: at my neighbours farm.
