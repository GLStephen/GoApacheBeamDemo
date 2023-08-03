// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// streaming_wordcap is a toy streaming pipeline that uses PubSub. It
// does the following:
//
//	(1) create a topic and publish a few messages to it
//	(2) start a streaming pipeline that converts the messages to
//	    upper case and logs the result.
//
// NOTE: it only runs on Dataflow and must be manually cancelled.
package main

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"

	"cloud.google.com/go/pubsub"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/options/gcpopts"
)

var (
	input   = flag.String("input", os.ExpandEnv("$USER-wordcap"), "Pubsub input topic.")
	command = flag.String("command", "generate", "Default command")
	count   = flag.Int("count", 10, "Default number of total strings")
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func main() {
	flag.Parse()

	ctx := context.Background()

	log.Infof(ctx, "%v %v %v %v", *input, *command, *count, utf8.RuneCountInString(*command))

	if utf8.RuneCountInString(*command) > 0 {
		switch *command {
		case "generate":
			log.Infof(ctx, "Running Generate Command for %d strings", *count)
			sendPubSubData(ctx, *count)
		}
	} else {
		displayArgError(ctx)
	}
}

func displayArgError(ctx context.Context) {
	log.Info(ctx, "Error with Input Data")
}

func sendPubSubData(ctx context.Context, totalStrings int) {
	project := gcpopts.GetProject(ctx)

	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		log.Fatal(ctx, err)
	}
	defer client.Close()

	topic, err := client.CreateTopic(ctx, *input)
	if err != nil {
		log.Info(ctx, "Topic Reused")
		topic = client.Topic(*input)
	} else {
		// topic created
		log.Info(ctx, "New Topic Created")
	}

	sub, err := client.CreateSubscription(ctx, topic.ID(), pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 600 * time.Second,
	})
	if err != nil {
		// this is informational for our purposes
		log.Infof(ctx, "Subcription Exists %v", err)
		sub = client.Subscription(topic.ID())
	} else {
		// topic created
		log.Info(ctx, "New Subscription Created")
	}

	log.Infof(ctx, "Subscription %s created for %s", sub.ID(), topic.ID())

	defer topic.Stop()
	var results []*pubsub.PublishResult
	for i := 0; i < totalStrings; i++ {
		// data = append(data, randStringRunes(rand.Intn(8)+1))
		message := randStringRunes(rand.Intn(8) + 1)
		res := topic.Publish(ctx, &pubsub.Message{Data: []byte(message)})
		results = append(results, res)
		log.Infof(ctx, "Publishing %s\n", message)
	}

	//log.Infof(ctx, "Published %v messages to: %v %v", len(results), *input, results)
}

// func sendPubSubDataOld(ctx context.Context, totalStrings int) {
// 	var data []string

// 	for i := 0; i < totalStrings; i++ {
// 		data = append(data, randStringRunes(rand.Intn(8)+1))
// 	}

// 	log.Infof(ctx, "Publishing %v messages to: %v %v", len(data), *input, data)

// 	project := gcpopts.GetProject(ctx)

// 	//defer pubsubx.CleanupTopic(ctx, project, *input)
// 	sub, err := pubsubx.Publish(ctx, project, *input, data...)
// 	if err != nil {
// 		log.Fatal(ctx, err)
// 	}
// 	log.Infof(ctx, "Subscription ID %v", sub.ID())
// }

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
