# Image API
> A simple API that takes an image file and returns that same image in black and white.

## Getting started

1. Build the project 
    
    ```
    make build
    ```

2. Run the server

    ```
    ./main
    ```

3. Make a `POST` request to `http://localhost:8080/images/greyscale` with a JPEG file as a `multipart/form-data` body under the name `image`

4. Make a `POST` request to `http://localhost:8080/images/pixelate` with a JPEG file as a `multipart/form-data` body under the name `image`