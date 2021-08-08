# Intel Ark API

<img src="https://img.shields.io/github/license/issy/intel-ark-api?&color=brightgreen&logo=github"/>

This is a RESTful API for the Intel product index.

NOTE: In order to run this on your own machine, you will need a copy of the Ark sqlite database, which can be extracted from the Ark android app.

## Endpoints

### `GET /search`

<details>
  <summary>Query parameters</summary>

  #### `query`

  The name of the product you're searching for
  ```
  Type: string
  Array: false
  Required: true
  ```

  #### `page`

  The page number to fetch. Pages are 50 items long
  ```
  Type: integer
  Default: 1
  Min: 1
  Max: max int64
  Array: false
  Required: false
  ```

</details>

<details>
  <summary>Response schema</summary>

  ```json
  {
    "$schema": "http://json-schema.org/schema",
    "$defs": {
      "product": {
        "type": "object",
        "required": [
          "id",
          "name",
          "specs"
        ],
        "properties": {
          "name": {
            "type": "string"
          },
          "id": {
            "type": "integer",
            "description": "The ID of the product, as represented on the Ark website. This can be used to generate a URL for this product on the Ark website"
          },
          "specs": {
            "type": "array",
            "description": "Each array within the array contains three string values. The first value is the spec category - the second value is the spec key - and the third value is the spec value",
            "items": {
              "type": "array",
              "minItems": 3,
              "maxItems": 3,
              "items": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "type": "array",
    "default": [],
    "items": {
      "$ref": "#/$defs/product"
    }
  }
  ```

</details>

<details>
  <summary>Example response</summary>

  ```json
  [
    {
      "id": 52210,
      "name": "Intel® Core™ i5-2500K Processor (6M Cache, up to 3.70 GHz)",
      "specs": [
        [
          "Advanced Technologies",
          "Intel® 64 <small><sup>‡</sup></small>",
          "Yes"
        ],
        [
          "Advanced Technologies",
          "Intel® Hyper-Threading Technology <small><sup>‡</sup></small>",
          "No"
        ],
        [
          "Advanced Technologies",
          "Intel® Virtualization Technology for Directed I/O (VT-d) <small><sup>‡</sup></small>",
          "No"
        ],
        [
          "CPU Specifications",
          "# of Cores",
          "4"
        ],
        [
          "CPU Specifications",
          "# of Threads",
          "4"
        ],
        [
          "CPU Specifications",
          "Cache",
          "6 MB Intel® Smart Cache"
        ],
        [
          "CPU Specifications",
          "Intel® Turbo Boost Technology 2.0 Frequency<small><sup>‡</sup></small>",
          "3.70 GHz"
        ],
        [
          "CPU Specifications",
          "Max Turbo Frequency",
          "3.70 GHz"
        ],
        [
          "CPU Specifications",
          "Processor Base Frequency",
          "3.30 GHz"
        ],
        [
          "CPU Specifications",
          "TDP",
          "95 W"
        ],
        [
          "Essentials",
          "Code Name",
          "Products formerly Sandy Bridge"
        ],
        [
          "Essentials",
          "Expected Discontinuance",
          "Q1'13"
        ],
        [
          "Essentials",
          "Launch Date",
          "Q1'11"
        ],
        [
          "Essentials",
          "Lithography",
          "32 nm"
        ]
      ]
    }
  ]
  ```

</details>

## Pagination

On endpoints that support pagination, you can specify a page number by using the `page` query parameter.
Pages are numbered starting from 1. If you pass in a number outside of the bounds of the amount of pages in the result, it will default to the maximum or minimum page, depending on which direction your page number was out of bounds by.

Pagination information is contained within the [link header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Link) in the API response. No link header will be sent if a search only yields 1 page