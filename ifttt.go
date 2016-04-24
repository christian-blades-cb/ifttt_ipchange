package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

const IFTTT_URL_FMT = "https://maker.ifttt.com/trigger/%s/with/key/%s"

type IFTTTEvent struct {
	Value1 string `json:"value1,omitempty"`
	Value2 string `json:"value2,omitempty"`
	Value3 string `json:"value3,omitempty"`
}

func (e IFTTTEvent) SendEvent(ctx context.Context, client *http.Client, name, key string) error {
	buf := bytes.Buffer{}

	enc := json.NewEncoder(&buf)
	err := enc.Encode(e)
	if err != nil {
		log.WithError(err).Fatal("unable to encode event")
	}

	eventUrl := fmt.Sprintf(IFTTT_URL_FMT, name, key)
	resp, err := ctxhttp.Post(ctx, client, eventUrl, "application/json", &buf)
	defer resp.Body.Close()
	if err != nil {
		log.WithError(err).WithField("name", name).Error("unable to send event")
		return err
	}

	log.WithFields(log.Fields{
		"name":       name,
		"statuscode": resp.StatusCode,
	}).Info("published event")
	return nil
}
