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
The `dist` folder will contain the distribution.

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
2. `npm install` to install all the dependencies
3. `npm test` to run the test 
4. `npm run` to run locally 

##### Folder structure
The `server` folder contains the Nginx server configuration will be used by Docker in order to run the frontend application.
The `client` contained the angular application application (Html, javascript, css)

###### Everything is a component
This is a component based application, this means that I followed the component pattern as defined in https://docs.angularjs.org/guide/component  
The `app` component is the root component and you can find his definition spread out across the files: 
1. `client/app.component.js` is the component definition
2. `client/app.controller.js` is the controller
3. `client/app.scss` 
4. `client/app.html` is the template 
5. `client/app.js` load the module that define the component (the last line of the code). Actually this file load all the module required by the application

__Let's take a pause...__
Basically, in the application, I used 2 different type of component: the one you can find in the __view__ folder and the one in the __component__ folder.  
The first one are connected with routing; this means that the url http://localhost:85/downloads will show what you wrote in the `__views/downloads__` folder.  
The second one are the "more ordinary" component, the one you can use within views or other components.  
Check out already defined component to create new ones.


__Let's continue with the folder...__
`client/index.html` is the page will be served. This file use the root component `<app></app>`.
`client/components/index.js` this file is in charge to load into the application the component modules. This means that if create a new compoent we have to update this file.   
`client/config` contains the configurations ofr the different environment
`client/config` contains the configurations ofr the different environment
