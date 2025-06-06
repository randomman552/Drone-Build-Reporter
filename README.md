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
    # Your settings here
    variable:
      from_secret: variable
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
  image: randomman552/drone-build-reporter
  settings:
    # Your settings here
    variable:
      from_secret: variable
  when:
    status:
      - failure
```

## Available platforms
All configuration options should be passed through the step settings as shown for the example above.

Values can be passed from secrets, or directly as strings. Care should be taken to use secrets where appropriate (API tokens, etc)

Global settings:
- `notify_mode` - Controls when in the process the CI this plugin runs
  - `finished` - Notify at the conclusion of the CI process (default)
  - `started` - Notify at the start of the CI process

Platform specific:
- Gotify
  - `gotify_token` the token from the gotify instance
  - `gotify_url` Url of the gotify instance, including scheme (e.g. `https://gotify.example.com`)
- Discord
  - `discord_webhook` The webhook to send the report to
