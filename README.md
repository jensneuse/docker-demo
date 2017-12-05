# Build fat docker image
docker build -f Dockerfile_fat .

# Analzye
docker images

# Build small docker image
docker build .

# Run application
docker run -p 8080:8080 -e PORT="8080" %image_name%