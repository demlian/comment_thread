# Comments Package

The `comments` package provides a simple implementation of a comment thread system. Each comment is represented by a `CommentNode` and the entire comment thread for a video is represented by a `CommentThread`.

## Design Choices

1. **Immutable comments**: In this implementation, once a comment is added, it cannot be deleted. This simplifies the design as there's no need to verify whether a comment exists before updating it.

2. **No spam protection**: The system does not check for identical comments, meaning that a user can add the same comment multiple times. This simplifies the design as there's no need for spam protection.

3. **No concurrency control**: The system assumes that users can only update their own comments. As such, there will not be concurrent attempts to update the same comment. This simplifies the design as there's no need for concurrency control mechanisms.