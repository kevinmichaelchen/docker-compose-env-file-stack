A demonstration of how to use the `.env` file properly.

You can provide the `.env` file as an `env_file`, but its
[primary purpose](https://docs.docker.com/compose/environment-variables/#the-env-file)
is to be used for variables in the Compose file itself, not in the container.

To make this more clear, I've created a `base.env` for container env vars.
```yaml
    env_file:
      - base.env
      - svc-a.env
    ports:
      - "${SVC_A_HOST_PORT}:9090"
```

The `SVC_A_HOST_PORT` comes from the `.env` file. It only affects the host's
port, not the actual port our container runs its HTTP server on.

9090 needs to be hard-coded, as that as what port gets used internally.

This is an informative [StackOverflow post](https://stackoverflow.com/questions/52664673/how-to-get-port-of-docker-compose-from-env-file):
> The env_file option will only set environment variables in the Docker
> container itself. Not on the host which is used during the Compose 'build'.