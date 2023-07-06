# GoApacheBeamDemo

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

### basic setup

* put strings into pubsub
* read strings into the system
* if X word occurs count and write out to a file|DB|??

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
