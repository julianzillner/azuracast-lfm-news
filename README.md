# azuracast-lfm-news 🎧📰

**azuracast-lfm-news** is a lightweight Docker-based service that automatically downloads the latest [laut.fm](https://laut.fm) news and uploads them to your AzuraCast instance — ideal for keeping your radio station up to date with fresh content.

## 🔧 Features

- Automatically fetches the latest laut.fm news audio
- Uploads the file to your AzuraCast server via SFTP
- Runs every hour at **xx:55**
- Easy configuration via `.env` file

---

## ⚙️ Environment Variables

You need to set the following environment variables in a `.env` file (example configuration):

```env
STATION_NAME=example-station
LIVE_PASSWORD=live-password
SFTP_HOST=sftp.example.com:2022
SFTP_USER=janedoe
SFTP_PASSWORD=janedoe1234
```
