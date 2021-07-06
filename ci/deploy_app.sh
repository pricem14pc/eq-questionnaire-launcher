#!/usr/bin/env bash
set -euxo pipefail


gcloud beta run deploy eq-questionnaire-launcher \
    --project="$PROJECT_ID" --region="$REGION" --concurrency=250 --min-instances="$MIN_INSTANCES" --max-instances="$MAX_INSTANCES" \
    --port=8000 --cpu=1 --memory=256M \
    --image="${DOCKER_REGISTRY}/eq-questionnaire-launcher:${IMAGE_TAG}" --platform=managed --allow-unauthenticated \
    --set-env-vars SURVEY_RUNNER_URL="https://${RUNNER_URL}"
