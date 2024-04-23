# Comment Thread Package

> **Please note:** This exercise is aimed at practicing Go language constructs, data structures, and design patterns.

The `comment_thread` package provides a simple implementation of a TCP server who uses a proprietary string based protocol to manage threaded comments. Each comment is represented by a `Comment` struct and a comment thread is represented by a `Discussion` struct.

## Design Choices

1. **Immutable comments**: In this implementation, once a comment is added, it cannot be deleted.
1. **No spam protection**: The system does not check for identical comments, meaning that a user can add the same comment multiple times.
1. **No concurrency control**: The system assumes that users can only update their own comments. As such, there will not be concurrent attempts to update the same comment.

## Request Format

Clients send requests in the following format:

<request_id>|<data>

The `request_id` is specified by the client and allows them to reconcile responses that they receive with requests that they send. The `data` field depends on the request type.

## Authentication

The system supports the following authentication requests:

|SIGN_IN| |WHOAMI |SIGN_OUT

