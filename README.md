# **PRJ CINEMATIK: Cloud Storage Event Processor with Pub/Sub Integration**

## **Overview**

This repository contains a **Cloud Function** written in **Go** that listens for `storage.finalized` events from **Google Cloud Storage (GCS)**. When a user uploads a video file, the function extracts its metadata and publishes it to a **Google Cloud Pub/Sub topic**. This enables downstream applications to process the event asynchronously.

## **Architecture**

```mermaid
graph TD;
    A[User Uploads Video] -->|storage.finalized Event| B[Cloud Storage Bucket]
    B -->|Triggers| C[Cloud Function (Go)]
    C -->|Extracts Metadata| D[Pub/Sub Topic]
    D -->|Publishes Message| E[Downstream Applications]
    E -->|Processes Event| F[AI Analysis / DB Storage / Notifications]
```

1. **User Uploads a Video** → A user uploads a video file to a designated **Cloud Storage bucket**.
2. **Cloud Function Trigger** → The **Cloud Function** is automatically triggered by a `storage.finalized` event.
3. **Metadata Extraction** → The function extracts metadata such as **bucket name, file path, content type, size, and creation timestamp**.
4. **Pub/Sub Publishing** → The extracted metadata is published to a **Google Cloud Pub/Sub topic**.
5. **Downstream Processing** → Other applications can subscribe to the topic and process the video file accordingly.

## **Features**

- **Automated Event Processing**: Triggers when a new file is uploaded to Cloud Storage.
- **Metadata Extraction**: Captures essential file attributes such as name, size, and content type.
- **Pub/Sub Integration**: Publishes structured event data to a Pub/Sub topic for further processing.
- **Scalable & Asynchronous**: Enables event-driven architectures by decoupling file uploads from downstream processing.

## **Prerequisites**

- **Google Cloud Platform (GCP) Account**
- **Cloud Storage Bucket** with `storage.finalized` event enabled
- **Google Cloud Pub/Sub Topic** for event publishing
- **IAM Permissions**:
  - `roles/storage.objectViewer` (for reading object metadata)
  - `roles/pubsub.publisher` (for publishing messages to Pub/Sub)
  - `roles/cloudfunctions.invoker` (for executing the function)

## **Deployment**

1. **Set Up the Environment**

   - Configure **GCP project**, **Cloud Storage bucket**, and **Pub/Sub topic**.
   - Ensure the correct **IAM roles** are assigned to the service account.

2. **Deploy the Cloud Function**

   - Use **Terraform** or **gcloud CLI** to deploy the function.
   - Specify the trigger as `storage.finalized` events.

3. **Validate Deployment**
   - Upload a test video to the **Cloud Storage bucket**.
   - Check **Pub/Sub messages** to confirm successful metadata extraction and publication.

## **Monitoring & Logging**

- **Google Cloud Logging**: Track function execution, errors, and performance metrics.
- **Google Cloud Monitoring**: Set up alerts for failures, performance degradation, or missing Pub/Sub messages.
- **Dead Letter Queue (DLQ) for Pub/Sub**: Configure DLQ for handling undelivered messages.

## **Security Best Practices**

- **Least Privilege IAM Policies**: Assign only necessary permissions to the function's service account.
- **Vulnerability Scanning**: Integrate **security scans** in CI/CD to identify dependency risks.
- **Data Encryption**: Ensure **Cloud Storage and Pub/Sub messages** use **encryption at rest and in transit**.

## **Use Cases**

- **Media Processing Pipelines**: Automate video ingestion workflows.
- **Metadata Indexing**: Store extracted metadata in databases for faster search and retrieval.
- **Machine Learning & AI**: Trigger AI-based video analysis models when new content is uploaded.
- **Compliance & Audit Logging**: Maintain logs of uploaded media for regulatory compliance.

## **Contributing**

- Follow the **branching strategy** (feature branches, PR reviews, and approvals).
- Ensure **all CI/CD checks pass** before submitting a pull request.
- Submit feature requests and bug reports via **GitHub Issues**.

## **License**

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.
