# ifttt_ipchange

Send IFTTT (If This Then That) events when the external ip changes for the local machine

# Building

Requires >= go 1.5

## Binary

If you're running go 1.5, `export GO15VENDOREXPERMENT=1`

`go build`

## Docker image

Set your `GOARCH` and `GOOS` to match your target (defaults to `arm/linux` because raspberry pi.)

`make`

# Usage

```
Usage:
  ifttt_ipchange [OPTIONS]

Application Options:
  -i, --interval=  how often (in seconds) to check for an ip change (default: 1200) [$CHECK_INTERVAL]
  -t, --timeout=   maximum time to wait for checking the ip and sending the event (default: 10) [$LOOP_TIMEOUT]
  -k, --key=       Key to use for sending and IFTTT event [$IFTTT_KEY]
  -n, --eventname= event name to use when sending IFTTT event (default: newhomeip) [$IFTTT_EVENTNAME]
      --debug      debug log messages

Help Options:
  -h, --help       Show this help message
```

