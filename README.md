Deployment command:

1. `docker build -f Dockerfiles/Dockerfile_backend -t oink_backend .`
2. `docker build -f Dockerfiles/Dockerfile_frontend -t oink_frontend .`
3. `docker stack deploy -c docker-compose.yaml --detach=false oink`
