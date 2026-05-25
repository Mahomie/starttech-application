# StartTech DevOps Architecture

## Overview

This project deploys a full-stack application using AWS infrastructure provisioned with Terraform.

## Components

### Frontend
- React application
- Hosted on Amazon S3
- Distributed with CloudFront CDN

### Backend
- Golang API
- Dockerized application
- Running on EC2 instances

### Load Balancing
- Application Load Balancer (ALB)

### Database
- MongoDB Atlas

### Caching
- Redis / ElastiCache

### Monitoring
- Amazon CloudWatch
- CloudWatch Alarms
- CloudWatch Logs

### CI/CD
- GitHub Actions
- DockerHub image registry

## Data Flow

User → CloudFront → S3 Frontend  
Frontend → ALB → EC2 Backend  
Backend → MongoDB Atlas  
Backend → Redis Cache

## Infrastructure as Code

Terraform provisions:
- VPC
- Subnets
- Security Groups
- EC2
- ALB
- IAM Roles
- Monitoring resources
