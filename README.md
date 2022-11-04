# deconz-ctl

## base url

- <http:////pi-aeng6eje.speedport.ip/api/$APIKEY/>

/api/<apikey>/lights/<id>/state

```shell
source .env
curl -X PUT http://pi-aeng6eje.speedport.ip/api/$DECONZ_API_KEY/lights/5/state -d '{"on":true}'
curl -X PUT http://pi-aeng6eje.speedport.ip/api/$DECONZ_API_KEY/lights/5/state -d '{"on":false}'
```

funtions:

- get configuration: <http://pi-aeng6eje.speedport.ip/api/$DECONZ_API_KEY/config>
- list all lights: <http://pi-aeng6eje.speedport.ip/api/$DECONZ_API_KEY/lights>
- list ligth: <http://pi-aeng6eje.speedport.ip/api/$DECONZ_API_KEY/lights/$id>
- power on ligth: <http://pi-aeng6eje.speedport.ip/api/$DECONZ_API_KEY/lights/5/state> -d '{"on":true}'
- power off light: <http://pi-aeng6eje.speedport.ip/api/$DECONZ_API_KEY/lights/5/state> -d '{"on":false}'

implementaions:

- implement table output
- parse response
- display functions, version commit date etc
- build for all architectures

references:

- <https://github.com/lukibahr/go-conbee/>
- <https://dev.to/plutov/writing-rest-api-client-in-go-3fkg>
