docker build -t hello .
docker run -d -p 8080:8080 hello
curl localhost:8080/world
