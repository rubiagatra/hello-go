# hello-go with docker

building the image

```bash
docker build -t go-web -f Dockerfile.go-web .
```

run the container

```bash
docker run -p 8000:8000 --name go-web go-web 
```