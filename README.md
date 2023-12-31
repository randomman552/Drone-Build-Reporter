# Drone-Build-Reporter
[![Build Status](https://drone.ggrainger.uk/api/badges/randomman552/Drone-Build-Reporter/status.svg?ref=refs/heads/main)](https://drone.ggrainger.uk/randomman552/Drone-Build-Reporter)

Plugin for [Drone CI](https://www.drone.io/) to send information about builds to multiple sources.

## Usage
To use this plugin, add the following step to your pipline
```yaml
# Report all builds
- name: build-reporter
  image: randomman552/drone-build-reporter
  settings:
    discord_webhook:
      from_secret: discord-webhook
  when:
    status:
      - failure
      - success
```
The above will trigger regardless of a build success or failure

To only trigger the pipeline on a failure, you can remove the success case from the condition as shown below
```yaml
# Report failures only
- name: build-reporter
  image: randomman552/drone-build-reporter:${DRONE_COMMIT_BRANCH}
  settings:
    discord_webhook:
      from_secret: discord-webhook
  when:
    status:
      - failure
```

## Available platforms
All configuration options should be passed through the step settings as shown for the example above.

This plugin currently supports the following platforms:
- Gotify
  - `gotify_token` the token from the gotify instance
  - `gotify_url` Url of the gotify instance, including scheme (e.g. `https://gotify.example.com`)
- Discord
  - `discord_webhook` The webhook to send the report tos
