# Use a Node.js base image with the desired version
FROM node:20 AS build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the Vue.js application
RUN npm run build

# Stage 2: Serve the built application with Vite's development server
FROM node:20 AS production

# Set the working directory in the container
WORKDIR /app

# Copy the built application from the previous stage
COPY --from=build /app/dist .

# Expose port 3000 (the default port for Vite)
EXPOSE 3000

# Define the command to run the development server
CMD ["npm", "run", "serve"]