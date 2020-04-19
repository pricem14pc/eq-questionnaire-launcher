# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: apps/v1
kind: Deployment
metadata:
  name: eq-questionnaire-launcher
  labels:
    app: eq-questionnaire-launcher
spec:
  replicas: 1
  selector:
    matchLabels:
      app: eq-questionnaire-launcher
  template:
    metadata:
      labels:
        app: eq-questionnaire-launcher
    spec:
      containers:
      - name: eq-questionnaire-launcher
        image: eu.gcr.io/GOOGLE_CLOUD_PROJECT/eq-questionnaire-launcher:COMMIT_SHA
        ports:
        - containerPort: 8000
        env:
          - name: SURVEY_RUNNER_URL
            valueFrom:
              secretKeyRef:
                name: launcher-secrets
                key: SURVEY_RUNNER_URL
---
kind: Service
apiVersion: v1
metadata:
  name: eq-questionnaire-launcher
spec:
  selector:
    app: eq-questionnaire-launcher
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8000
  type: LoadBalancer