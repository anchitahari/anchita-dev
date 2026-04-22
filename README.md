# anchita-dev

My personal portfolio site, rebuilt from scratch as a learning project.
The backend is written in Go, containerized with Docker, and originally
deployed on AWS EKS. It includes an AI-powered chatbot that answers questions
about my experience using RAG (Retrieval-Augmented Generation).

The backend has since been migrated from EKS to Railway to reduce infrastructure
costs. EKS was over-engineered for a personal project. The original EKS setup
(VPC, managed node groups, cluster config) is preserved in [`terraform/main.tf`](terraform/main.tf)
as a reference for how it was provisioned.

This repo is a work in progress that I'm building and documenting it
publicly as I go.

## Tech stack

- **Backend:** Go (net/http)
- **Frontend:** HTML, CSS → migrating to Astro (Phase 3)
- **Containerization:** Docker
- **Orchestration:** Kubernetes on AWS EKS (original deployment)
- **Deployment:** Railway (current)
- **Infrastructure:** Terraform
- **AI:** LLM API + vector database (RAG)

## Project phases

- [x] Phase 1 — Go HTTP server with API skeleton
- [x] Phase 2 — Dockerize and deploy to EKS
- [x] Phase 3 — Frontend rebuild in Astro
- [x] Phase 4 — AI chatbot feature
- [x] Phase 5 — Terraform, CI/CD, observability

## Running locally

**Prerequisites:** Go 1.22+
```bash
git clone https://github.com/your-username/anchita-dev.git
cd anchita-dev
go run cmd/main.go
```

Server runs at `http://localhost:8080`

**API endpoints:**

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/health` | Liveness check |
| POST | `/api/chat` | AI chatbot |

## Why I built this

I'm a software engineer who wanted a project that would let me learn Go, Kubernetes, and AI integration in one place 
while building something I actually have a reason to maintain and improve.

I'm documenting the journey publicly on
[LinkedIn](https://www.linkedin.com/in/anchita-hn).

## Contact

[anchita.hari@gmail.com](mailto:anchita.hari@gmail.com)