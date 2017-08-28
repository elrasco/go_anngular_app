# README #
This README documents the necessary steps to get your application up and running.  

### What is this repository for? ###
This is the frontend of empatica homework. 
The frontend is written in angularjs 1.6 with ECMA6 syntax
Along apart the angular application, the project contains also the configuration for the nginx server supposed to serve


#### Features
1. Bundled with webpack
2. ECMA6 syntax, compiled by Babel
2. Test written with mocha-chai syntax
3. Test executed Karma Test Runner
4. Support for sass template
5. Hot module replacement plugin (you can see changes without stopping the server)
6. Multi environment configuration

#### Build and Test
##### Development
`npm run build:dev` will create the distribution with a configuration suited for a development environment (see __Configuration__ for details).  
`npm run build:prod` will create the distribution with a configuration suited for a production environment (see __Configuration__ for details).  
The dist folder will containd the distribution.

##### Test
`npm test` will execute all the test.

#### Steps for getting up and running (with docker)
1. Build the image
`docker build -t --name empatica_frontend .` 
2. Run the container
`docker run -d -p 85:80 empatica_frontend`

#### Configuration
This is the frontend side of a single page web application, this means it supposed to have access to a backend service.
The location of this service can be easily configured setting the property API_URL.  

##### client/config/local.js
This file contains the configuration of your local development.   
Create the file and add the API_URL property if you want to change the behavioour defined in client/config/development.js file. 
This file will affect the configuration only when you launch the application with `npm start` (your local development)  

##### client/config/production.js
This file affect the configuration only when you run `npm run build:prod` 

##### client/config/development.js
This file affect the configuration only when you run `npm run build:dev` 

#### Development
1. install nodejs (v6.11.2) + npm (v3.10.10)
2. `npm install`
3. `npm test`
4. `npm run` to run locally 

##### Everything is a component
This is a component based application, this means that I followed the component pattern as defined in https://docs.angularjs.org/guide/component  
The app component is the root component.

##### Folder structure
The `server` folder contains the Nginx server configuration will be used by Docker in order to run the frontend application.
The `client` contained the angular application application (Html, javascript, css)
The `client/index.html` is the page will be served. This file contains the root component (app).
The `client/app.html` is the page will be served
The `client/app.js` load all the module (external and not) required by the application and define the app component
