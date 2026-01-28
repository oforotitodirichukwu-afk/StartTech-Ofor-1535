StartTech Full-Stack Deployment

A containerized React frontend and Go/Node backend application deployed on AWS App Runner using Terraform and Docker.

Live Demo
- Frontend: https://mayd7wqi72.us-east-1.awsapprunner.com
- Backend API: https://93hrzaeur2.us-east-1.awsapprunner.com

Tech Stack
- Frontend: React (Dockerized)
- Infrastructure: Terraform (IaC)
- Cloud: AWS (ECR, App Runner)
- CI/CD: GitHub Actions

Deployment Process
1. Containerization: Optimized Dockerfiles for minimal image size.
2. Registry: Images pushed to Amazon Elastic Container Registry (ECR).
3. Orchestration: Terraform scripts used to provision App Runner services.
4. Environment: WSL2 (Ubuntu) integrated with Docker Desktop for Windows.

Project Structure
- /frontend: React application and Docker configuration.
- /backend: API service.
- /infra: Terraform configuration files.
