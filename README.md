# heartbreak 💔

Use the Twitter API to unlike your Twitter Likes.

## Install

Run: `go install github.com/victoriadrake/heartbreak`

Alternatively, clone this repository, then build the program:

```bash
cd heartbreak
go build
```

## Usage

Expects a `.env` file with the following:

```env
export TWITTER_ACCESS_TOKEN= ...
export TWITTER_ACCESS_TOKEN_SECRET= ...
export TWITTER_CONSUMER_KEY= ...
export TWITTER_CONSUMER_SECRET= ...
export TWITTER_USERNAME= ...
```

This program fails safely. When run without the `-unlike` flag, or any arguments at all, it prints usage instructions:

```txt
Usage: heartbreak [options]

  -filename string
        Optional name of the file to store archived Tweets. (default "likes-archive.txt")
  -no-archive
        Do not archive unliked Tweets.
  -unlike
        Required to unlike all liked Tweets. Logs Tweet text to a file in the current directory.
```

To unlike all your likes and by default, log the Tweets to a file called `likes-archive.txt`:

```bash
heartbreak -unlike
```

Specify your own filename with `-filename`, for example:

```bash
heartbreak -unlike -filename unlikes.log
```

Not one for nostalgia? I get it. Skip logging with `-no-archive`:

```bash
heartbreak -unlike -no-archive
```

New day, new you.
