# anchita-dev

My personal portfolio site, rebuilt from scratch as a learning project.
The backend is written in Go, containerized with Docker, and deployed
on AWS EKS. It includes an AI-powered chatbot that answers questions
about my experience using RAG (Retrieval-Augmented Generation).

This repo is a work in progress — I'm building and documenting it
publicly as I go.

## Tech stack

- **Backend:** Go (net/http)
- **Frontend:** HTML, CSS → migrating to React (Phase 3)
- **Containerization:** Docker
- **Orchestration:** Kubernetes on AWS EKS
- **Infrastructure:** Terraform
- **AI:** LLM API + vector database (RAG)

## Project phases

- [x] Phase 1 — Go HTTP server with API skeleton
- [ ] Phase 2 — Dockerize and deploy to EKS
- [ ] Phase 3 — Frontend rebuild in React
- [ ] Phase 4 — AI chatbot feature
- [ ] Phase 5 — Terraform, CI/CD, observability

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
| POST | `/api/chat` | AI chatbot (coming in phase 4) |

## Why I built this

I'm a software engineer transitioning toward developer advocacy and
wanted a project that would let me learn Go, Kubernetes, and AI
integration in one place — while building something I actually have
a reason to maintain and improve.

I'm documenting the journey publicly on
[LinkedIn](https://www.linkedin.com/in/anchita-hn).

## Contact

[anchita.hari@gmail.com](mailto:anchita.hari@gmail.com)