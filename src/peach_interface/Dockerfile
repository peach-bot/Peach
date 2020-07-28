FROM node:13.12.0-alpine
COPY package.json .
COPY package-lock.json .
RUN npm install --silent
RUN npm install -g serve
COPY . .
RUN npm run build
CMD serve -s build -l tcp://0.0.0.0:5000
EXPOSE 5000
