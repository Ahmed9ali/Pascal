# Changelog

## v0.1.1
- Rebranded from Crush to Pascal: name, logo, exit banner, tool names, environment markers, config paths and built-in guides.
- Telemetry and provider-list updates are now opt-in (`PASCAL_ENABLE_METRICS`, `PASCAL_ENABLE_PROVIDER_UPDATE`).
- Yolo and sysadmin modes require `PASCAL_ALLOW_YOLO=1`.
- API keys load from `~/.config/pascal/env` (must be `chmod 600`); the `.env` autoloader was removed.
- The agent confirms the stack with the user before starting a new project.
- Greetings and general questions are answered without tool calls.
- Added a CI workflow (build, vet, test) and `build.sh`.

## v0.1.0
- First working fork of Crush, renamed to Pascal.
