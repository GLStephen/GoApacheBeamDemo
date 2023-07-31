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
	"fmt"
	"os"

	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/options/gcpopts"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/util/pubsubx"
)

var (
	input = flag.String("input", os.ExpandEnv("$USER-wordcap"), "Pubsub input topic.")
)

var (
	data = []string{
		"foo",
		"bar",
		"bazy",
		"bazsdf",
		"bazss",
		"bazsd",
		"bazsd",
		"bazsaaaa",
	}
)

func main() {
	ctx := context.Background()

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	if len(os.Args) > 1 {
		arg := os.Args[3]

		fmt.Println(argsWithProg)
		fmt.Println(argsWithoutProg)
		fmt.Println(arg)
	} else {
		displayArgError(ctx)
	}
}

func displayArgError(ctx context.Context) {
	log.Info(ctx, "Publishing %v messages to: %v", len(data), *input)
}

func sendPubSubData(ctx context.Context) {
	flag.Parse()

	project := gcpopts.GetProject(ctx)

	log.Infof(ctx, "Publishing %v messages to: %v", len(data), *input)

	defer pubsubx.CleanupTopic(ctx, project, *input)
	_, err := pubsubx.Publish(ctx, project, *input, data...)
	if err != nil {
		log.Fatal(ctx, err)
	}
}
