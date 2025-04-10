# Traffic Light Simulation

Displays an ASCII image of a standard US traffic light in red, yellow and green colors. The lights will automatically change based on user provided durations for each color.

## Build

This will build the binary ready for distribution, debug and symbols are removed from the binary and file paths are trimmed. This will help keep the binary size as small as we can.

```shell
$ make build
```

## Running

To run the default settings you can use simply `make run`. To control the timing of lights use the following after building:

```shell
$ ./traffic-lights -red 8s -yellow 5s -green 10s
```
