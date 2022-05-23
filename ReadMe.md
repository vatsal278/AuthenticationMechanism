## Rest API with authentication mechanism

#### This Rest API was created using Golang and mongodb. It features functions like sign up login and getting user details from database.

#### This api has used authentication mechanism, middle ware and different clean code principles.

### In order to use this Api:

* clone the repo : `git clone https://github.com/vatsal278/Go-Api-with-Authentication-Mechanism.git`
* Pull official mongodb image from docker hub : `docker pull mongo`
* build the docker file :  `docker build --name Api-image .`
* Create a new isolated docker network : `docker network create api_net`
* Start a mongo container instance : `docker run --network api_net --name mongo-container --publish 9091:27017 -d mongo`
* Check the ip address of mongo container in my case IP was `172.0.0.2` keep the port number as `:27017` as it is default port for mongo db : `docker network inspect api_net`
* #### Running the Api locally: 
* *`go run -e DBADDRESS=172.0.0.2:27017 main.go`

* #### Running the Api in container :
* * `docker run --network api_net --name api-container --e DBADDRESS=172.0.0.2:27017 --publish 9092:8080 Api-image`

