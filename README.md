# Image Resizer & Cropper API

This project is an image manipulation API built using [GoFiber](https://gofiber.io/) and [nfnt/resize](https://pkg.go.dev/github.com/nfnt/resize) for resizing and cropping images. The API allows you to resize or crop an image by specifying the dimensions and the image URL.

## Features

- **Resize Images**: Resize an image by specifying the desired width and height.
- **Crop Images**: Crop an image by providing the crop width, height, and starting coordinates (`x` and `y`).
- **Supported Formats**: PNG and JPEG image formats are supported.

## Installation

1. Clone the repository:
```bash
git clone https://github.com/KonyD/image-manipulation-api.git
cd image-manipulation-api
```

2. Install the dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

## Usage

You can use this API to resize or crop images by sending HTTP GET requests with the appropriate query parameters.

# Resizing an Image

To resize an image 

* `image` (required): The URL of the image to resize.
* `width` (required): The desired width of the image.
* `height` (required): The desired height of the image.
* `mode` (optional): If set to `"resize"`, the image will be resized (default behavior if mode is omitted).

Example Request:
```bash
http://localhost:3000/?image=https://example.com/image.png&width=200&height=100&mode=resize
```

# Cropping an Image

To crop an image, make a GET request with the following query parameters:

* `image` (required): The URL of the image to crop.
* `width` (required): The width of the cropped area.
* `height` (required): The height of the cropped area.
* `x` (required): The x-coordinate of the top-left corner where cropping should start.
* `y` (required): The y-coordinate of the top-left corner where cropping should start.
* `mode` (required): Set to `"crop"` for cropping mode.

Example Request:
```bash
http://localhost:3000/?image=https://example.com/image.png&width=100&height=100&x=50&y=50&mode=crop
```

# Supported Formats

The API supports both PNG and JPEG formats. Based on the input image format, the API will return the resized/cropped image in the same format.

# Error Handling

The API will return appropriate error messages for invalid or missing parameters:

* **400 Bad Request:** Returned when the query parameters (e.g., width, height, etc.) are invalid.
* **502 Bad Gateway:** Returned when the image URL is invalid or inaccessible.
* **415 Unsupported Media Type:** Returned if the image format is not supported.

## Dependencies

This project uses the following Go packages:

- **[GoFiber](https://gofiber.io/)**: A fast HTTP web framework for building web applications in Go.
- **[nfnt/resize](https://pkg.go.dev/github.com/nfnt/resize)**: A library for resizing images in Go.
- **[image/jpeg](https://pkg.go.dev/image/jpeg)**: The standard Go library for working with JPEG image formats.
- **[image/png](https://pkg.go.dev/image/png)**: The standard Go library for working with PNG image formats.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
