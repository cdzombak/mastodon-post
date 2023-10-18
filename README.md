# mastodon-post

The simplest possible CLI tool for posting to Mastodon, ideally suited for use in simple bots.

## Usage

```text
mastodon-post [OPTIONS] -text "text to post"
```

### Options

- `-text`: Text to post. Required.
- `-visibility`: Visibility of the post. One of `public`, `unlisted`, or `private`. Defaults to `public`.
- `-help`: Print help and exit.
- `-version`: Print version and exit.

## Credentials and Server Configuration

Credentials and the server's address are provided via the environment variables `MASTODON_SERVER`, `MASTODON_CLIENT_ID`, `MASTODON_CLIENT_SECRET`, and `MASTODON_ACCESS_TOKEN`.

`MASTODON_SERVER` is the URL of your Mastodon server (for example, `https://mastodon.social`).

To get a Client ID/Secret and Access Token:

1. Navigate to your Mastodon Preferences
2. Select "Developer" in the Sidebar
3. Click "New Application" at the top of the page
4. Enter a name for your application (for example, "mastodon-post")
5. Use `urn:ietf:wg:oauth:2.0:oob` for the Redirect URI
6. Deselect all scopes, then select only `write:statuses`
7. Click "Submit"
8. The following page will show your Client ID, Client Secret, and Access Token

You can optionally provide these environment variables by placing a `.env` file in the working directory from which you run `mastodon-post`. The program will read the file if it exists and set the environment variables automatically.

If running the program via Docker, you can pass the same file to the `--env-file` option.

An empty `.env` file is included in this repository to help you get started: [`.env.template`](https://github.com/cdzombak/mastodon-post/blob/main/.env.template).

## Installation

### macOS via Homebrew

```shell
brew install cdzombak/oss/mastodon-post
```

### Debian via apt repository

Install my Debian repository if you haven't already:

```shell
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://dist.cdzombak.net/deb.key | sudo gpg --dearmor -o /etc/apt/keyrings/dist-cdzombak-net.gpg
sudo chmod 0644 /etc/apt/keyrings/dist-cdzombak-net.gpg
echo -e "deb [signed-by=/etc/apt/keyrings/dist-cdzombak-net.gpg] https://dist.cdzombak.net/deb/oss any oss\n" | sudo tee -a /etc/apt/sources.list.d/dist-cdzombak-net.list > /dev/null
sudo apt-get update
```

Then install `mastodon-post` via `apt-get`:

```shell
sudo apt-get install mastodon-post
```

### Manual installation from build artifacts

Pre-built binaries for Linux and macOS on various architectures are downloadable from each [GitHub Release](https://github.com/cdzombak/mastodon-post/releases). Debian packages for each release are available as well.

### Build and install locally

```shell
git clone https://github.com/cdzombak/mastodon-post.git
cd mastodon-post
make build

cp out/mastodon-post $INSTALL_DIR
```

## Docker images

Docker images are available for a variety of Linux architectures from [Docker Hub](https://hub.docker.com/r/cdzombak/mastodon-post) and [GHCR](https://github.com/cdzombak/unshorten/pkgs/container/mastodon-post). Images are based on the `scratch` image and are as small as possible.

Run them via, for example:

```shell
docker run --rm --env-file /path/to/.env cdzombak/mastodon-post:1 -text "message to post"
docker run --rm --env-file /path/to/.env ghcr.io/cdzombak/mastodon-post:1 -text "message to post"
```

## About

- Issues: [github.com/cdzombak/mastodon-post/issues](https://github.com/cdzombak/mastodon-post/issues)
- Author: [Chris Dzombak](https://www.dzombak.com)
  - [GitHub: @cdzombak](https://www.github.com/cdzombak)

## License

MIT; see `LICENSE` in this repository.
