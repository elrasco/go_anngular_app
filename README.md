# README #
This README documents the necessary steps to get the empatica homework get up and running. 

### What is this repository for? ###

This repository contains both backend and frontend of Empatica Homework.  
You can find frontend application in __frontend__ folder and the backend application in the __backend__ folder.   

This solution can be deployed as multi container docker application, this means that you can use doccker-compose tool to get it up and running.   
You can deploy without docker as well (check the READMEs of the single projects).

The frontend is written in angularjs (1.6) using ES6 syntax.  
The backend is writtend in GO.  
There are no database in the stack.

For the details of each projects you can refer to the READMEs file of the project.  

###### You need to install docker and docker-compose tool to run the solution.

### Steps for getting up and running

1. clone the repo 
`git clone git@bitbucket.org:luca_rasconi/empa_homework.git`

2. enter the project folder 
`cd empa_homework`

3. run the solution
`docker-compose up -d`  
to reload use `docker-compose up -d --no-deps --build`

4. go to http://localhost:85

##### Architecture
This is a single page application composed by a front-end application and a api backend.   
The front end side is served by an nginx server that acts as proxy to the api backend. In this way I can call the api using relative url without paying attention to cors problem.    
For the homework I choose to not use database because, imho, it wasn't so important. Of course I would choose one in a real application.   

##### Expansion
In order to fullfill the real time requirement, the first options I can see is to create a comminication channel from server to the client, this means that as soon as a change happens, clients will be noticed about it.  
In this scenario the WebSocket techonlogy can help us implementing this solutions.  
