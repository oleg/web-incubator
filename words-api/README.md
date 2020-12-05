#run java
mvn clean spring-boot:run

#run js
cd src/main/resources/appjs
npm start

#database
docker run --name db12 -p 5432:5432 -e POSTGRES_PASSWORD=xxx -d postgres

docker exec -it db12 psql -U postgres