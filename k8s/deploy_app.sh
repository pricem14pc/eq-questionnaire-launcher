#!/usr/bin/env bash
set -euxo pipefail

helm upgrade --install \
    questionnaire-launcher \
    k8s/helm \
    --set surveyRunnerUrl=https://${RUNNER_URL} \
    --set image.repository=${DOCKER_REGISTRY}/eq-questionnaire-launcher \
    --set image.tag=${IMAGE_TAG}
