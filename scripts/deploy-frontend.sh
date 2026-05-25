#!/bin/bash

set -e

echo "Building frontend..."

cd Client

npm install
npm run build

echo "Deploying to S3..."

aws s3 sync dist/ s3://YOUR_BUCKET_NAME --delete

echo "Invalidating CloudFront cache..."

aws cloudfront create-invalidation \
  --distribution-id YOUR_DISTRIBUTION_ID \
  --paths "/*"

echo "Frontend deployment complete."
