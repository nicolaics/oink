Deployment command:
`docker build -f Dockerfiles/Dockerfile_backend -t oink_backend .`
`docker build -f Dockerfiles/Dockerfile_frontend -t oink_frontend .`
`docker stack deploy -c docker-compose.yaml --detach=false oink`
