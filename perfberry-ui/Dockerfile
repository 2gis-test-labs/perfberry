FROM node:12.13.0-alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package.json /usr/src/app/
COPY package-lock.json /usr/src/app/
RUN npm install --ignore-scripts
COPY . /usr/src/app
RUN npm run postinstall

EXPOSE 5000

CMD [ "npm", "start" ]
