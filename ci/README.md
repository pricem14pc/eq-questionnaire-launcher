# Deployment with [Concourse](https://concourse-ci.org/)

To deploy this application with Concourse, you must be logged in to a Concourse instance.

---

Make sure required [environmental variables](../README.md#environment-variables) are set.

To deploy the app via Concourse, use the following task command:

```sh
PROJECT_ID=<project_id> \
DOCKER_REGISTRY=<docker_registry> \
RUNNER_URL=<runner_instance_url> \
fly -t <target_concourse_instance> execute \
  --config ci/deploy.yaml
```
