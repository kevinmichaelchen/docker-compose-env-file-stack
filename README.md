A demonstration of possibly unintuitive behavior.

Here we have a simple Go program that logs its `PORT` env var and spins up a
simple HTTP server on that port.

We also have a Docker Compose file that stacks 2 different `env_file` options,
and does a port-forwarding.
```yaml
    env_file:
      - .env
      - app.env
    ports:
      - "${PORT}:${PORT}"
```

## Expectations
`.env` has a port of 8080, and `app.env` has a port of 9090. Because `app.env`
is stacked on top of `.env`, its 9090 that is ultimately set as the value.
This is as you'd expect.

However, the port forwarding is where things are surprising. It's forwarded as
```
8080:8080
```