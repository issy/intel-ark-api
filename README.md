# Intel Ark API

<img src="https://img.shields.io/github/license/issy/intel-ark-api?&color=brightgreen&logo=github"/>

This is a RESTful API for the Intel product index.

NOTE: In order to run this on your own machine, you will need a copy of the Ark sqlite database, which can be extracted from the Ark android app.

## Endpoints

### `GET /search/:search-term`

<details>
  <summary>Response schema</summary>

  ```json
  {
    "type": "array",
    "default": [],
    "items": {
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
          "type": "integer"
        },
        "specs": {
          "type": "array",
          "items": {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      }
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
          "Enhanced Intel SpeedStep® Technology",
          "Yes"
        ],
        [
          "Advanced Technologies",
          "Idle States",
          "Yes"
        ],
        [
          "Advanced Technologies",
          "Instruction Set",
          "64-bit"
        ],
        [
          "Advanced Technologies",
          "Instruction Set Extensions",
          "Intel® SSE4.1, Intel® SSE4.2, Intel® AVX"
        ],
        [
          "Advanced Technologies",
          "Intel vPro® Platform Eligibility <small><sup>‡</sup></small>",
          "No"
        ],
        [
          "Advanced Technologies",
          "Intel® 64 <small><sup>‡</sup></small>",
          "Yes"
        ],
        [
          "Advanced Technologies",
          "Intel® Fast Memory Access",
          "Yes"
        ],
        [
          "Advanced Technologies",
          "Intel® Flex Memory Access",
          "Yes"
        ],
        [
          "Advanced Technologies",
          "Intel® Hyper-Threading Technology <small><sup>‡</sup></small>",
          "No"
        ],
        [
          "Advanced Technologies",
          "Intel® Identity Protection Technology <small><sup>‡</sup></small>",
          "Yes"
        ],
        [
          "Advanced Technologies",
          "Intel® Turbo Boost Technology <small><sup>‡</sup></small>",
          "2.0"
        ],
        [
          "Advanced Technologies",
          "Intel® Virtualization Technology (VT-x) <small><sup>‡</sup></small>",
          "Yes"
        ],
        [
          "Advanced Technologies",
          "Intel® Virtualization Technology for Directed I/O (VT-d) <small><sup>‡</sup></small>",
          "No"
        ],
        [
          "Advanced Technologies",
          "Intel® VT-x with Extended Page Tables (EPT) <small><sup>‡</sup></small>",
          "Yes"
        ],
        [
          "Advanced Technologies",
          "Thermal Monitoring Technologies",
          "Yes"
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
          "Bus Speed",
          "5 GT/s"
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
        ],
        [
          "Essentials",
          "Off Roadmap",
          "No"
        ],
        [
          "Essentials",
          "Processor Number",
          "i5-2500K"
        ],
        [
          "Essentials",
          "Product Collection",
          "Legacy Intel® Core™ Processors"
        ],
        [
          "Essentials",
          "Status",
          "Discontinued"
        ],
        [
          "Essentials",
          "Vertical Segment",
          "Desktop"
        ],
        [
          "Expansion Options",
          "Max # of PCI Express Lanes",
          "16"
        ],
        [
          "Expansion Options",
          "PCI Express Revision",
          "2.0"
        ],
        [
          "Memory Specifications",
          "ECC Memory Supported   <small><sup>‡</sup></small>",
          "No"
        ],
        [
          "Memory Specifications",
          "Max # of Memory Channels",
          "2"
        ],
        [
          "Memory Specifications",
          "Max Memory Bandwidth",
          "21 GB/s"
        ],
        [
          "Memory Specifications",
          "Max Memory Size (dependent on memory type)",
          "32 GB"
        ],
        [
          "Memory Specifications",
          "Memory Types",
          "DDR3 1066/1333"
        ],
        [
          "Package Specifications",
          "Max CPU Configuration",
          "1"
        ],
        [
          "Package Specifications",
          "Package Size",
          "37.5mm x 37.5mm"
        ],
        [
          "Package Specifications",
          "Sockets Supported",
          "LGA1155"
        ],
        [
          "Package Specifications",
          "T<sub>CASE</sub>",
          "72.6°C"
        ],
        [
          "Processor Graphics",
          "# of Displays Supported <small><sup>‡</sup></small>",
          "2"
        ],
        [
          "Processor Graphics",
          "Device ID",
          "0x112"
        ],
        [
          "Processor Graphics",
          "Graphics Base Frequency",
          "850 MHz"
        ],
        [
          "Processor Graphics",
          "Graphics Max Dynamic Frequency",
          "1.10 GHz"
        ],
        [
          "Processor Graphics",
          "Intel® Clear Video HD Technology",
          "Yes"
        ],
        [
          "Processor Graphics",
          "Intel® Flexible Display Interface (Intel® FDI)",
          "Yes"
        ],
        [
          "Processor Graphics",
          "Intel® InTru™ 3D Technology",
          "Yes"
        ],
        [
          "Processor Graphics",
          "Intel® Quick Sync Video",
          "Yes"
        ],
        [
          "Processor Graphics",
          "Processor Graphics <small><sup>‡</sup></small>",
          "Intel® HD Graphics 3000"
        ],
        [
          "Security & Reliability",
          "Execute Disable Bit <small><sup>‡</sup></small>",
          "Yes"
        ],
        [
          "Security & Reliability",
          "Intel® AES New Instructions",
          "Yes"
        ],
        [
          "Security & Reliability",
          "Intel® Trusted Execution Technology <small><sup>‡</sup></small>",
          "No"
        ],
        [
          "Supplemental Information",
          "Datasheet",
          "http://www.intel.com/content/www/us/en/processors/core/CoreTechnicalResources.html"
        ],
        [
          "Supplemental Information",
          "Embedded Options Available",
          "No"
        ]
      ]
    }
  ]
  ```

</details>