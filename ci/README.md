# Deployment with [Concourse](https://concourse-ci.org/)

To deploy this application with Concourse, you must have a Kubernetes cluster and be logged in to a Concourse instance.

---

The following environment variables must be set when deploying the app:
- RUNNER_URL
- PROJECT_ID
- DOCKER_REGISTRY

The following are optional variables that can also be set if needed:
- SERVICE_ACCOUNT_JSON
- IMAGE_TAG
- REGION

To deploy to a cluster you can run the following command

```sh
RUNNER_URL=<runner_instance_url> \
PROJECT_ID=<project_id> \
DOCKER_REGISTRY=<docker_registry> \
fly -t <target_concourse_instance> execute \
  --config ci/deploy.yaml
```
