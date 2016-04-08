# Fleet Unit Creator

## Creates and prints (to standard output) the fleet unit files corresponding to the `services.yaml` file found on the `rootURI`.

## Build & Run app

`go build`

`./fleet-unit-creator[.exe] -rootURI="https://raw.githubusercontent.com/Financial-Times/up-service-files/debug-services/"`

## Build & Run image

`docker build -t my/image .`

```
export ROOT_URI=https://raw.githubusercontent.com/Financial-Times/up-service-files/debug-services/
export $FLEET_REVISION=master
docker run --env="ROOT_URI=$ROOT_URI" --env="REVISION=$REVISION" imageName
```
