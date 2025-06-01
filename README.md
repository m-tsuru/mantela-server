# Mantela Server

Mantela Server is a simple server that serves and proxies mantela.json files.

Some Features:
- Serve mantela.json from a file.
- Proxy mantela.json from a URL.
- Override mantela.json with a source mantela file.

> [!IMPORTANT]
> Mantela Server is a sample implementation.
> It is unofficial and not affiliated with the mantela project.

## Usage

1. copy `config.sample.toml` to `config.toml` and change it as you like.

2. save mantela file that you want to override to source mantela file.

3. run mantela-server.

## Docker

You can run mantela-server with Docker.

1. You need to do Step 1, 2 in Usage section before running Docker.

2. Run the following command:

```bash
docker compose up -d
```

> [!NOTE]
>
> however, In the future, the official tkytel may provide a compose.yaml that allows quick start of customized MikoPBX.
> At that time, this repository may follow suit and provide an additional mantela-adder with customized MikoPBX.

## Tips

if the destination of `diff` is not set in `config.toml`, mantela-server serves `source` mantela simply.

If mantela.json is fetched from the web, it can be treated as a proxy, and if mantela.json is served from a file, it can be treated as mantela.json hosting software.

This means that it can be treated as a mantela.json server that can be easily hosted in a virtual environment.
