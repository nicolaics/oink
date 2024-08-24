# Deployment command:

1. `docker build -f Dockerfiles/Dockerfile_backend -t oink_backend .`
2. `docker build -f Dockerfiles/Dockerfile_frontend -t oink_frontend .`
3. `docker stack deploy -c docker-compose.yaml --detach=false oink`


# For testing:

1. Generating dummy data run the Python script by using `python3 create_dummy_data.py`
2. Enter the backend `host:port` which is where the application is hosted at
   *Note: if it is in the localhost, you can do `localhost:1338`
