## 1. Create a Google Cloud Platform (GCP) account, AND name the project 'notely'.
```text
+ We'll keep everything for Notely in a single project, and when done you can simply delete the project to clean everything up in one place.

* Next, you'll need to create a billing account. This is where you'll provide your credit card information. You can find the billing section in the GCP console by clicking the hamburger menu in the top left, then "Billing".
* Ensure your billing account is linked to your project, and you are able to see the billing information for your project in the GCP console.
```

## 2. Install and setup via Google Cloud SDK
### Install 'gcloud' CLI tool from here: https://cloud.google.com/sdk/docs/install

### Initialized it:
```text
Initialize it by running gcloud init in your terminal. If you are using WSL, use gcloud init --console-only instead.
    It will prompt you to login by opening a browser window. Login with the same account you used to create your GCP project.
    Select your notely project.
```

```bash
gcloud init
```

### Run the following command to verify your authenticated account and project settings:
```bash
gcloud config list
```

### Troubleshooting: if you already authenticated after running 'gcloud init'. If not, run:
```bash
gcloud auth login
```

## 3. Create 'Artifact Registry' in GCP UI.
```text
1) In the GCP console, search for and enable the Cloud Build API.
2) Within Artifact Registry in the GCP console, enable the Artifact Registry API.
3) Click Create Repository:
    * Name: notely-ar-repo
    * Format: Docker
    * Mode: Standard
    * Location type: Region
    * Region for deployment: us-central1
    * Leave "Google-managed encryption key" selected
    * Disable vulnerability scanning if the option appears; we don't need paid scans for this course
```

### Troubleshoot: check if repository can be listed:
```bash
gcloud artifacts repositories list
```

## 4. Build and push the Docker image to Google's (GCP) Artifact Registry
### example commands:
```bash
gcloud builds submit --tag REGION-docker.pkg.dev/PROJECT_ID/REPOSITORY/IMAGE:TAG .
```
### final command:
```bash
gcloud builds submit --tag us-central1-docker.pkg.dev/notely-504508/notely-ar-repo/notely:latest .
```

### Troubleshooting:
```text
Note: If you run into that same Cloud Build storage permission error again,
you would need to grant the default Cloud Build service account (559262895345-compute@developer.gserviceaccount.com)
the Storage Object Viewer (roles/storage.objectViewer) role in IAM, or grant the Cloud Build Service Account role permissions to access your Cloud Storage bucket.
```