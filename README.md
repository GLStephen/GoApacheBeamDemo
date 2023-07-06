# GoApacheBeamDemo

## resources

* https://beam.apache.org/documentation/sdks/go/

## based on example

* https://github.com/apache/beam/blob/master/sdks/go/examples/streaming_wordcap/wordcap.go

## concept

* put strings into pubsub
* read strings into the system
* if X word occurs count and write out to a file|DB|??

### run process in 3 types of batches
* run simple one word entry then run batch
* input 1 million words
  * run larger version with capped dataflow workers in batch mode (slowish)
  * run larger version with more workers in batch mode (faster than previous)
