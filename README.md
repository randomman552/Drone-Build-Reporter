# Drone-Build-Reporter
[![Build Status](https://drone.ggrainger.uk/api/badges/randomman552/Drone-Build-Reporter/status.svg?ref=refs/heads/main)](https://drone.ggrainger.uk/randomman552/Drone-Build-Reporter)

Plugin for [Drone CI](https://www.drone.io/) to send information about builds to multiple sources.

## Usage
To use this plugin, add the following step to your pipline
```yaml
# Reporting
- name: build-reporter
  image: randomman552/drone-build-reporter
  settings:
    gotify_token:
      from_secret: gotify-token
    gotify_url:
      from_secret: gotify-url
    discord_webhook:
      from_secret: discord-webhook

```

## Available sources
Currently this plugin only reports to a gotify instance, but I plan to add more sources when I need. E.g. Discord Webhook, MS Teams Webhook

## Configuration
| Variable          | Default | Description                                           |
|:-----------------:|:-------:|:-----------------------------------------------------:|
| gotify_token      |         | Token used when sending gotify messages               |
| gotify_url        |         |  URL of gotify instance, including scheme (https://)  |
| discord_webhook   |         |  Discord webhook to call, including scheme (https://) |
