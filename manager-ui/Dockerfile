FROM node:lts

WORKDIR /usr/src/app/api

COPY package*.json ./

COPY . .

RUN npm install

EXPOSE 3000 

CMD ["npm", "run", "dev"]
