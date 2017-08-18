FROM sinet/nginx-node:latest

# Install and build the application
COPY . /usr/src/app
WORKDIR /usr/src/app

RUN apt-get install yarn
RUN rm -rf node_modules \
    && yarn install \
    && npm run build:prod

COPY default.conf /etc/nginx/conf.d/

CMD ["nginx", "-g", "daemon off;"]  