FROM node:22-bookworm 

COPY package*.json ./

RUN npm install --production

COPY . .

CMD ["npm", "run", "dev"]
