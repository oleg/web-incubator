docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest
docker run -it --network some-network --rm mysql mysql -hsome-mysql -uexample-user -p
