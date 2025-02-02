# Cinematik - GCP Architecture Documentation

## Overview

This document outlines the architecture for a streaming platform build on Google Cloud Platform (GCP). The system is designed to handle video content management, processing and streaming at scale.

## Event Flow

```mermaid
graph TD
    A[Upload Video] --> B[Raw Storage Bucket]
    B --> C{Upload Handler}
    C --> E[Publish Event]
    E --> D[Extract Metadata]
    E --> F[Transcoding Service]
    F --> G[Cloud Transcoder]
    G --> H[Processed Storage]
    D --> I[Metadata Service]
    I --> J[(Cloud SQL)]
    I --> K[(Firestore)]
    H --> L[Cloud CDN]
    L --> M[End Users]
```
