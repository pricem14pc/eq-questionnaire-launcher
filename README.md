# EQ Launcher

This project was copied from https://github.com/ONSdigital/go-launch-a-survey and should be used for v3 of runner.

### Building and Running
Install Go and ensure that your `GOPATH` env variable is set (usually it's `~/go`).

```
go get
go build
./eq-questionnaire-launcher

go run launch.go (Does both the build and run cmd above)
```

Open http://localhost:8000/

### Docker
The dockerfile is a multistage dockerfile which can be built using:

```
docker build -t eq-questionnaire-launcher:latest .
```

You can then run the image using `SURVEY_RUNNER_SCHEMA_URL` to point it at an instance of survey runner.

```
docker run -e SURVEY_RUNNER_SCHEMA_URL=http://localhost:5000 -it -p 8000:8000 onsdigital/eq-questionnaire-launcher:latest
```

The syntax for this will be slightly different on Mac

```
docker run -e SURVEY_RUNNER_SCHEMA_URL=http://host.docker.internal:5000 -it -p 8000:8000 onsdigital/eq-questionnaire-launcher:latest
```

You should then be able to access go launcher at `localhost:8000`

You can also run a Survey Register for launcher to load Schemas from

```
docker run -it -p 8080:8080 onsdigital/eq-survey-register:simple-rest-api
```

### Run Quick-Launch
If the schema specifies a `schema_name` field, that will be used as the schema_name claim. If not, the filename from the URL (before `.`) will be used.

Run Questionnaire Launcher
```
scripts/run_app.sh
```
Now run Go launcher and navigate to "http://localhost:8000/quick-launch?url=" passing the url of the JSON
```
e.g."http://localhost:8000/quick-launch?url=http://localhost:7777/1_0001.json"
```

### Deploying

For deploying with Concourse see the [CI README](./ci/README.md).

### Deployment with [gcloud](https://cloud.google.com/sdk/gcloud)

To deploy this application with gcloud, you must be logged in using `gcloud auth login` and `gcloud auth application-default login`.

Make sure required [environmental variables](./README.md#environment-variables) are set.

Then call the following command with environment variables set:
```sh
./ci/deploy_app.sh
```
## Environment Variables

The following environment variables must be set when deploying the app:
- PROJECT_ID
- DOCKER_REGISTRY
- RUNNER_URL

The following are optional variables that can also be set if needed:
- SERVICE_ACCOUNT_JSON
- REGION
- IMAGE_TAG
- CONCURRENCY
- MIN_INSTANCES
- MAX_INSTANCES
- CPU
- MEMORY

### Notes
* There are no unit tests yet
* JWT spec based on http://ons-schema-definitions.readthedocs.io/en/latest/jwt_profile.html

### Settings
Environment Variable | Meaning | Default
---------------------|---------|--------
GO_LAUNCH_A_SURVEY_LISTEN_HOST|Host address  to listen on|0.0.0.0
GO_LAUNCH_A_SURVEY_LISTEN_PORT|Host port to listen on|8000
SURVEY_RUNNER_URL|URL of Survey Runner to re-direct to when launching a survey|http://localhost:5000
SURVEY_REGISTER_URL|URL of eq-survey-register to load schema list from |http://localhost:8080
JWT_ENCRYPTION_KEY_PATH|Path to the JWT Encryption Key (PEM format)|jwt-test-keys/sdc-user-authentication-encryption-sr-public-key.pem
JWT_SIGNING_KEY_PATH|Path to the JWT Signing Key (PEM format)|jwt-test-keys/sdc-user-authentication-signing-launcher-private-key.pem
