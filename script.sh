docker build -t go-app .
docker run -d -p 8080:8080 --name c1 go-app
