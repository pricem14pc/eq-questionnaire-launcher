# Deployment with [Concourse](https://concourse-ci.org/)

To deploy this application with Concourse, you must be logged in to a Concourse instance.

---

Make sure the required [environment variables](../README.md#environment-variables) are set.
IMAGE_TAG can be passed as an environment variable or as a file input located at `./image-tag/tag`.

To deploy the app via Concourse, use the following task command:

```sh
PROJECT_ID=<project_id> \
DOCKER_REGISTRY=<docker_registry> \
RUNNER_URL=<runner_instance_url> \
fly -t <target_concourse_instance> execute \
  --config ci/deploy.yaml
```
