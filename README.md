# Excel To Grafana

There was a need to show data from an updating `.xlsx` file in [Grafana](https://grafana.com/). But sadly there is no free plugin for it.

Converting this:

![.xlsx file](/docs/images/excel.PNG)

To this:

![.xlsx file](/docs/images/grafana-dashboard.PNG)

## Our Solution

A [golang](https://go.dev/) µsertvice which will convert a `.xlsx` file into a `.csv` format every `1 sec` and host this file locally to `:8081` so the [CSV](https://grafana.com/grafana/plugins/marcusolsson-csv-datasource/) plugin could read it in real time and show the data inside a [Grafana](https://grafana.com/) dashboard.

## How to run it

### Run excel-to-csv 

First we have to run the container using the following `CLI` command:

```cmd
docker run -p 8081:8081 -v /path/to/.xlsx/folder:/app/metrics zigelboimmisha/excel-to-csv  
```

### Create new Data Source

In [Grafana](https://grafana.com/) go to `Administration` > `Data sources`.
In there create a new `CSV` data source.

Please make sure you install this [CSV](https://grafana.com/grafana/plugins/marcusolsson-csv-datasource/) plugin first.

There are 2 variables you could change while running this µsertvice:

| Name | Default | Description |
| --- | --- | --- |
| `FILE_NAME` | `Book` | The `.xlsx` file to export |
| `SPREADSHEET` | `Sheet1` | The spread sheet you want to export |

### Creating a new Dashboard

Now you can create a dashboard using the new self updating data source, or use this `Panel JSON`:

```json
{
  "datasource": {
    "type": "marcusolsson-csv-datasource",
    "uid": "bd26b10d-5db3-4927-987d-77accda6fd66"
  },
  "description": "Freq Per User",
  "fieldConfig": {
    "defaults": {
      "custom": {
        "drawStyle": "line",
        "lineInterpolation": "stepAfter",
        "barAlignment": 0,
        "lineWidth": 1,
        "fillOpacity": 0,
        "gradientMode": "none",
        "spanNulls": false,
        "showPoints": "auto",
        "pointSize": 5,
        "stacking": {
          "mode": "none",
          "group": "A"
        },
        "axisPlacement": "auto",
        "axisLabel": "",
        "axisColorMode": "text",
        "scaleDistribution": {
          "type": "linear"
        },
        "axisCenteredZero": false,
        "hideFrom": {
          "tooltip": false,
          "viz": false,
          "legend": false
        },
        "thresholdsStyle": {
          "mode": "off"
        }
      },
      "color": {
        "mode": "palette-classic"
      },
      "mappings": [],
      "thresholds": {
        "mode": "absolute",
        "steps": [
          {
            "color": "green",
            "value": null
          }
        ]
      }
    },
    "overrides": []
  },
  "gridPos": {
    "h": 10,
    "w": 24,
    "x": 0,
    "y": 0
  },
  "id": 1,
  "options": {
    "tooltip": {
      "mode": "single",
      "sort": "none"
    },
    "legend": {
      "showLegend": true,
      "displayMode": "list",
      "placement": "bottom",
      "calcs": []
    }
  },
  "pluginVersion": "10.0.0",
  "targets": [
    {
      "datasource": {
        "type": "marcusolsson-csv-datasource",
        "uid": "bd26b10d-5db3-4927-987d-77accda6fd66"
      },
      "decimalSeparator": ".",
      "delimiter": ",",
      "header": true,
      "hide": false,
      "ignoreUnknown": false,
      "refId": "A",
      "schema": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "skipRows": 0
    }
  ],
  "title": "Freq Per User",
  "transformations": [
    {
      "id": "convertFieldType",
      "options": {
        "conversions": [
          {
            "destinationType": "string",
            "targetField": "Name"
          },
          {
            "destinationType": "time",
            "targetField": "Timestamp"
          },
          {
            "destinationType": "number",
            "targetField": "Freq"
          }
        ],
        "fields": {}
      }
    },
    {
      "id": "partitionByValues",
      "options": {
        "fields": [
          "Name"
        ]
      }
    }
  ],
  "type": "timeseries"
}
```