# GoApacheBeamDemo

## after container start

run the following to initialize your gcloud

```bash
gcloud init
```

then configure the projects for gcloud that you want to use

ensure dataflow API is enabled

ensure compute engine API is enabled

then run on the dataflow runner, as outlined here: https://cloud.google.com/dataflow/docs/quickstarts/create-pipeline-go

example
```bash
go run wordcount.go --input gs://dataflow-samples/shakespeare/kinglear.txt \
            --output gs://BUCKET_NAME/results/outputs \
            --runner dataflow \
            --project PROJECT_ID \
            --region DATAFLOW_REGION \
            --staging_location gs://BUCKET_NAME/binaries/
```

specific

```bash
go run main.go --runner dataflow \
            --project mulan-372517 \
            --region us-east1 \
            --staging_location gs://bucket-wordcount-example/binaries/
```

monorepo transition

```bash
go run -C ./src/processor/ main.go --runner dataflow --project mulan-372517 --region us-east1 --staging_location gs://bucket-wordcount-example/binaries/ --async
```

## other ideas and noes

https://beam.apache.org/get-started/wordcount-example/

# OLD

## resources

* https://beam.apache.org/documentation/sdks/go/

## based on example

* https://github.com/apache/beam/blob/master/sdks/go/examples/streaming_wordcap/wordcap.go

## concept

The value of the demo is to show the value of DataFlow and Apache Beam with Go not just for linear "fast" processing, but for parallel processing.

Can be done batch (think similar to hadoop) or stream.

The idea is that this is fundamental data processing. What you do with this is up to you.

Conceptually = Language + Parallel Processing Framework + Orchestration Environment for Runners
Specifics = Go + Beam + GCP Dataflow

Called Kappa Architecture vs Lambda - Kappa unifies streaming and batch

heavily focuses stored streaming

### basic setup

* put strings into pubsub
* read strings into the system
* lowercase them

### run process in 3 types of batches
* run simple one word entry then run batch
* input 1 million words
  * run larger version with capped dataflow workers in batch mode (slowish)
  * run larger version with more workers in batch mode (faster than previous)

## things to talk about

* batch sumaries
* windowed
* complex calculations

# for demo

write tool to enter items into pubsub

* pubsub A with 10 items in it
* pubsub B with 1,000,000 items in it
* pubsub C with 1,000,000 items in it

psgen <channel_name> <count>
psgen pubsub_demo_channelA 10
psgen pubsub_demo_channelB 1000000
psgen pubsub_demo_channelC 1000000

## data generated

random word strings like

aardvark must make money with wooty woot

## metrics created by app

count words, and count by first letter of each word

## concepts to cover in presetnation

* the need to process a set of items where the processing is unrelated is a scaling or unit of work problem
* when you need to process lots of items in parallel and the results are related to each other you need a framework like beam
* example: processing independent queued messages vs. processing, counting or analyzing a set of related messages in a queue


