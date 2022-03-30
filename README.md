# aiVLE CLI

Course administration utility for [aiVLE](https://github.com/edu-ai/aivle-web) (AI Virtual Learning Environment).

## Getting started

Before running the executable, you need to create a `.env` file in the same directory with the following content:
```
API_ROOT=http://192.168.3.51:8000
```

1. `API_ROOT`: root of aiVLE backend API, for example, `http://127.0.0.1:8000` or `https://aivle-api.leotan.cn/api/v1`

## Features

1. Download submissions
2. Download evaluation results as CSV file
3. Upload LumiNUS student roster Excel to aiVLE course whitelist
4. Get API token from username and password

## Note

To export student roster from LumiNUS, please use the following config:

![You should only include email and select "No Photos" for the "Photo Size".](/images/luminus_export_example.png)
