package main

import (
	"context"
	"log"
	"time"

	"github.com/fluent/fluent-bit-go/plugin"
)

func init() {
	plugin.RegisterOutput("go-test-output-plugin", "Golang output plugin for testing", dummyPlugin{})
}

type dummyPlugin struct{}

func (plug dummyPlugin) Init(ctx context.Context, conf plugin.ConfigLoader) error { return nil }

func (plug dummyPlugin) Collect(ctx context.Context, tag string, ch <-chan plugin.Message) error {
	for msg := range ch {
		log.Printf("message=\"got record\" tag=%s time=%s record=%+v\n", tag, msg.Time.Format(time.RFC3339), msg.Record)
	}

	return nil
}

func main() {}