# Resize images with configuration file

Provided a filename for a jpeg image the app uses a sizes.json file to detemine the resizing of the images

## Structure of json configuration file

An example of a json configuration file:

    {
        "medidas": [{
                "imageType": "androidFullHdpi",
                "width": 474,
                "height": 299
            },
            {
                "imageType": "androidFullMdpi",
                "width": 316,
                "height": 199
            }
        ]
    }

## Usage

    ./imagecardart originalimage
    originalimage:  it's the original image to resize
