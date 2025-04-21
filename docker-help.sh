#!/bin/bash

build_image() {
    docker build -t mp3-converter .
}

run_container() {
    docker run --name mp3-converter-container -d mp3-converter
}

copy_output() {
    docker cp mp3-converter-container:/output ./output
}

# Example usage
# Uncomment the following lines to execute the functions or manually call each function as needed :)
# build_image
# run_container
# copy_output