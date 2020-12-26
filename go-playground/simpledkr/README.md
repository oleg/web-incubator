```shell
docker build -t my-golang-app .
docker run -it --rm --name my-running-app -p 8080:8080 my-golang-app 
```