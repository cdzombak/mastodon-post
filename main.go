package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mattn/go-mastodon"
)

var version = "<dev>"

// Environment variable names used to configure the program.
const (
	EnvVarServer       = "MASTODON_SERVER"
	EnvVarClientID     = "MASTODON_CLIENT_ID"
	EnvVarClientSecret = "MASTODON_CLIENT_SECRET"
	EnvVarAccessToken  = "MASTODON_ACCESS_TOKEN"
)

func eprintf(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
}

func usage() {
	eprintf("usage: %s -text \"message to post\" [OPTIONS]\n", filepath.Base(os.Args[0]))
	eprintf("\nOptions:\n")
	flag.PrintDefaults()
	eprintf("\nEnvironment variables:\n")
	eprintf("  %s\n    \tMastodon server's address.\n", EnvVarServer)
	eprintf("  %s\n    \tYour app's Client ID.\n", EnvVarClientID)
	eprintf("  %s\n    \tYour app's Client Secret.\n", EnvVarClientSecret)
	eprintf("  %s\n    \tYour Access Token.\n", EnvVarAccessToken)
	// TODO(cdzombak): instructions on getting client ID/secret and access token
	eprintf("\nVersion:\n  mastodon-post %s\n", version)
	eprintf("\nGitHub:\n  https://github.com/cdzombak/mastodon-post\n")
	eprintf("\nAuthor:\n  Chris Dzombak <https://www.dzombak.com>\n")
}

func main() {
	text := flag.String("text", "", "Text to post to Mastodon")
	visibility := flag.String("visibility", "public", "Visibility of the post (public, unlisted, private)")
	printVersion := flag.Bool("printVersion", false, "Print printVersion and exit")
	flag.Usage = usage
	flag.Parse()

	if *printVersion {
		fmt.Printf("%s %s", filepath.Base(os.Args[0]), version)
		os.Exit(0)
	}

	if *visibility != "public" && *visibility != "unlisted" && *visibility != "private" {
		eprintf("invalid post visibility '%s' given\n", *visibility)
		os.Exit(1)
	}

	if *text == "" {
		flag.Usage()
		os.Exit(1)
	}

	config := mastodon.Config{
		Server:       os.Getenv(EnvVarServer),
		ClientID:     os.Getenv(EnvVarClientID),
		ClientSecret: os.Getenv(EnvVarClientSecret),
		AccessToken:  os.Getenv(EnvVarAccessToken),
	}
	if config.Server == "" {
		eprintf("environment variable %s is missing\n", EnvVarServer)
		os.Exit(1)
	}
	if !strings.HasPrefix(strings.ToLower(config.Server), "http") {
		config.Server = "https://" + config.Server
	}
	if config.ClientID == "" {
		eprintf("environment variable %s is missing\n", EnvVarClientID)
		os.Exit(1)
	}
	if config.ClientSecret == "" {
		eprintf("environment variable %s is missing\n", EnvVarClientSecret)
		os.Exit(1)
	}
	if config.AccessToken == "" {
		eprintf("environment variable %s is missing\n", EnvVarAccessToken)
		os.Exit(1)
	}

	c := &mastodon.Client{
		Client:    *http.DefaultClient,
		Config:    &config,
		UserAgent: "",
	}

	result, err := c.PostStatus(context.Background(), &mastodon.Toot{
		Status:     *text,
		Visibility: *visibility,
		// Sensitive:   false, // TODO(cdzombak): file an issue
		// SpoilerText: "", // TODO(cdzombak): file an issue
	})
	if err != nil {
		fmt.Printf("error posting status: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("posted: %s\n", result.URL)
}
