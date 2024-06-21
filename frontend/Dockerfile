
FROM node:14-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

EXPOSE 8080

# Define the command to run your app using npm start (replace with your start command if different)
CMD ["npm", "dev"]
