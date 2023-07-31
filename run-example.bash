#! /bin/bash

source customize.env

printf 'Compiling and running on %s in region %s With %d Workers in bucket %s\n' $GCP_PROJECT $GCP_REGION $DATAFLOW_MAX_NUM_WORKERS $GCP_BUCKET

# run the generator
source .bin/generator.bash

# run the dataflow job
source .bin/dataflow.bash