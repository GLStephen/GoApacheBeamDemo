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
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/options/gcpopts"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/transforms/stats"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/util/pubsubx"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/debug"
)

var (
	input = flag.String("input", os.ExpandEnv("$USER-wordcap"), "Pubsub input topic.")
)

func main() {
	flag.Parse()
	beam.Init()

	ctx := context.Background()
	project := gcpopts.GetProject(ctx)

	// creating the client
	pubSubClient, err := pubsub.NewClient(ctx, project)
	if err != nil {
		log.Fatal(ctx, err)
	}

	// ensure the topic exists, so we don't unecessarily blow up
	sub, err := pubsubx.EnsureTopic(ctx, pubSubClient, *input)
	if err != nil {
		log.Fatal(ctx, err)
	}

	log.Infof(ctx, "Running streaming wordcap with subscription: %v", sub.ID())

	p := beam.NewPipeline()
	s := p.Root()

	// get data
	col := pubsubio.Read(s, project, *input, &pubsubio.ReadOptions{Subscription: sub.ID()})

	// capitalize
	str := beam.ParDo(s, func(b []byte) string {
		return (string)(b)
	}, col)
	cap := beam.ParDo(s, strings.ToUpper, str)
	debug.Print(s, stats.CountElms(s, cap))

	// count
	counted := stats.Count(s, cap)
	formatted := beam.ParDo(s, func(w string, c int) string {
		return fmt.Sprintf("Counts %s: %v", w, c)
	}, counted)

	// output the data
	debug.Print(s, formatted)

	//log.Infof(ctx, "Coverted %s to %s.", s, cap)

	if err := beamx.Run(context.Background(), p); err != nil {
		log.Exitf(ctx, "Failed to execute job: %v", err)
	}
}
