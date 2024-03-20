# Movie Journal API

A go project to serve react frontend movie blog

## Installation

### Docker Based Installation

#### Step -01

- Git clone the app

#### Step -02

- Setup `docker/.env` file from `docker/.env.example`

  > Change the `DOCKER_BUILD_MODE=dev`  if needed.
- Setup `docker-compose.override.yml` file
  > duplicate the `docker-compose.override.example.yml` to `docker-compose.override.yml`.

 > if needed (for dev environment) add docker-compose.override.yml inside docker/.env add this in COMPOSE_FILE key 
- Setup `docker/.envs/app.env` file from `docker/.envs/app.env.example`
  > Note: Application env file is here, so any changes might require docker app service to restart

#### Step -03

- Goto `docker` directory. `cd docker/`
- Run `docker-compose build app` first inside for building app service
- Run `docker-compose build` for all services
- Run `docker-compose ps -a` for checking process status
- Run `docker-compose up` to run the application
