package main // import "github.com/christian-blades-cb/ifttt_ipchange"

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/cenkalti/backoff"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	CheckInterval  int64  `short:"i" long:"interval" default:"1200" env:"CHECK_INTERVAL" description:"how often (in seconds) to check for an ip change"`
	LoopTimeout    int64  `short:"t" long:"timeout" default:"10" env:"LOOP_TIMEOUT" description:"maximum time to wait for checking the ip and sending the event"`
	IftttKey       string `short:"k" long:"key" env:"IFTTT_KEY" description:"Key to use for sending and IFTTT event" required:"true"`
	IftttEventName string `short:"n" long:"eventname" default:"newhomeip" env:"IFTTT_EVENTNAME" description:"event name to use when sending IFTTT event"`
	Debug          func() `long:"debug" description:"debug log messages"`
}

func init() {
	opts.Debug = func() {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		log.WithError(err).Debug("unable to parse arguments")
		return
	}

	expBackoff := backoff.ExponentialBackOff{
		InitialInterval:     backoff.DefaultInitialInterval,
		RandomizationFactor: backoff.DefaultRandomizationFactor,
		Multiplier:          backoff.DefaultMultiplier,
		MaxInterval:         2 * time.Minute,
		MaxElapsedTime:      0,
		Clock:               backoff.SystemClock,
	}

	basectx := context.Background()
	looptimeout := time.Second * time.Duration(opts.LoopTimeout)
	currentIp := "initial state"

	for {
		log.Debug("main loop")

		loopctx, _ := context.WithTimeout(basectx, looptimeout)

		currentIp, err = compareAndNotify(loopctx, http.DefaultClient, opts.IftttEventName, opts.IftttKey, currentIp)
		if err != nil {
			log.WithError(err).WithField("retryduration", expBackoff.GetElapsedTime()).Warn("retrying")
			time.Sleep(expBackoff.NextBackOff())
			continue
		}

		expBackoff.Reset()
		time.Sleep(time.Duration(opts.CheckInterval) * time.Second)
	}
}

func compareAndNotify(ctx context.Context, client *http.Client, eventname string, key string, currentip string) (newip string, err error) {
	fip, err := GetIpV4(ctx, client, MY_FING_IPV4_URL)
	if err != nil {
		return currentip, err
	}
	newip = fip.IPAddress

	if newip != currentip {
		event := IFTTTEvent{Value1: newip, Value2: currentip}
		if err := event.SendEvent(ctx, http.DefaultClient, eventname, key); err != nil {
			return currentip, err
		}
	}

	return
}
