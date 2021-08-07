# Intel Ark API

<img src="https://img.shields.io/github/license/issy/intel-ark-api?&color=brightgreen&logo=github"/>

This is a RESTful API for the Intel product index.

NOTE: In order to run this on your own machine, you will need a copy of the Ark sqlite database, which can be extracted from the Ark android app.

## Endpoints

### `GET /search`

<details>
  <summary>Query parameters</summary>

  #### `query`

  The product you're searching for
  ```
  Type: string
  Array: false
  Required: true
  ```

  #### `count`

  Determines the max amount of results to be returned by the API
  ```
  Type: integer
  Default: 50
  Min: 1
  Max: 50
  Array: false
  Required: false
  ```

  #### `firstRecord`

  Determines the index of the first record to be returned by the API, relative to the length of the results from the DB
  ```
  Type: integer
  Default: 0
  Min: 0
  Max: 9223372036854775807
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
            "minItems": 3,
            "maxItems": 3,
            "description": "Each array within the array contains three string values. The first value is the spec category - the second value is the spec key - and the third value is the spec value",
            "items": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "type": "object",
    "properties": {
      "totalRecords": {
        "type": "integer",
        "description": "How many records were retrieved by the whole search. Note that the amount of records contained in the products array may be smaller. This value is present for pagination purposes"
      },
      "firstRecord": {
        "type": "integer",
        "description": "The index of the first record in the products array, relative to the whole search query. This value is present for pagination purposes"
      },
      "products": {
        "type": "array",
        "default": [],
        "items": {
          "$ref": "#/$defs/product"
        }
      }
    }
  }
  ```

</details>

<details>
  <summary>Example response</summary>

  ```json
  {
    "totalRecords": 10,
    "firstRecord": 3,
    "products": [
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
  }
  ```

</details>