FROM node:18-alpine
WORKDIR /app
# These next lines will be instant because they are already CACHED
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Install a tiny web server to serve the files
RUN npm install -g serve

EXPOSE 80
# Use 'serve' to run the app on port 80
CMD ["serve", "-s", "build", "-l", "80"]