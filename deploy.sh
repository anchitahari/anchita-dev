#!/bin/bash

set -e

echo "🔐 Authenticating Docker to ECR..."
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 261142222458.dkr.ecr.us-east-1.amazonaws.com

echo "🏗️ Building image for linux/amd64..."
docker build --platform linux/amd64 -t anchita-dev:latest .

echo "🏷️ Tagging image..."
docker tag anchita-dev:latest 261142222458.dkr.ecr.us-east-1.amazonaws.com/anchita-dev:latest

echo "📤 Pushing to ECR..."
docker push 261142222458.dkr.ecr.us-east-1.amazonaws.com/anchita-dev:latest

echo "🚀 Deploying to EKS..."
kubectl apply -f k8s/
kubectl rollout restart deployment anchita-dev
kubectl rollout status deployment anchita-dev

echo "✅ Deployment complete!"