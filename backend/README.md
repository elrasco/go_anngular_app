# README #
This README documents the necessary steps to get your application up and running.  

### What is this repository for? ###
This is the backend of empatica homework. 
The backend is written in GO using the standard libray plus gin and gin-cors libraries for rest call handling. I used an offline country reverse geocoder (github.com/AsGz/geo/georeverse). 

The purpose of the project is to expose 2 calls in order to let the frontend access "informations" about the downloads of app:   
1. `GET /api/downloads?from=YYYY-MM-DD&to=YYYY-MM-DD`  
This call returns the list of the downloaded mobile app within a interval. from e to are not mandatory   
2. `GET /api/statistics?from=YYYY-MM-DD&to=YYYY-MM-DD`  
This call returns the statistics of the downloaded mobile app within a interval. from e to are not mandatory

#### Steps for getting up and running (if you don't have docker)

##### You need to install GO to run the solution outsite docker.   

1. copy the application in your GOPATH folder  
`mkdir $GOPATH/src/backeand_app;
cp . $GOPATH/src/backeand_app`
2. stay into the application workdir  
cd $GOPATH/src/backeand_app
3. get the dependencies  
`go get`
4. install  
`go install`
5. run  
`backeand_app`

#### Steps for getting up and running (with docker)
1. Build the image  
`docker build -t empatica_backend .`
2. Run the container     
`docker run -d -p 8080:8080 empatica_backend`

http://localhost:8080/api/downloads   
http://localhost:8080/api/statistics