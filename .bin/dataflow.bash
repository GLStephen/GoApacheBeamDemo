go run -C ./src/processor/ main.go \
    --runner dataflow \
    --project $GCP_PROJECT \
    --region $GCP_REGION \
    --staging_location gs://$GCP_BUCKET/binaries/ \
    --autoscaling_algorithm THROUGHPUT_BASED \
    --max_num_workers $DATAFLOW_MAX_NUM_WORKERS \
    --num_workers 2 \
    --async