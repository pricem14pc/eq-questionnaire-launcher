# Deployment with [Concourse](https://concourse-ci.org/)

To deploy this application with Concourse, you must have a Kubernetes cluster and be logged in to a Concourse instance.

---

The following environment variables must be set when deploying the app:
- PROJECT_ID
- DOCKER_REGISTRY
- RUNNER_URL

The following are optional variables that can also be set if needed:
- SERVICE_ACCOUNT_JSON
- REGION
- IMAGE_TAG

	To deploy the app to the cluster via Concourse, use the following task command, specifying the `image_registry` and the `build_image_version` variables:

```sh
PROJECT_ID=<project_id> \
DOCKER_REGISTRY=<docker_registry> \
RUNNER_URL=<runner_instance_url> \
fly -t <target_concourse_instance> execute \
  --config ci/deploy.yaml
  -v image_registry=<docker-registry> \
  -v build_image_version=<image-tag>
```
