FROM node:22-bookworm 

WORKDIR /app/src

COPY package*.json ./

RUN npm install --production

COPY . .

CMD ["npm", "run", "dev"]
