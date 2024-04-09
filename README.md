# Comments Package

> **Please note:** Exercise aimed at practicing Go language constructs, data structures, and design patterns.
 

The `comments` package provides a simple implementation of a comment thread system. Each comment is represented by a `CommentNode` and the entire comment thread for a video is represented by a `CommentThread`.

## Design Choices

1. **Immutable comments**: In this implementation, once a comment is added, it cannot be deleted.
1. **No spam protection**: The system does not check for identical comments, meaning that a user can add the same comment multiple times.
1. **No concurrency control**: The system assumes that users can only update their own comments. As such, there will not be concurrent attempts to update the same comment.
1. **Storage Adapter Interface**: The system uses a `StorageAdapter` interface to abstract the storage mechanism.
1. **InMemoryStorage Implementation**: The `InMemoryStorage` struct is an implementation of the `StorageAdapter` interface. It stores comments in a map, with the comment ID as the key.
1. **Parent ID Approach**: The system uses the Parent ID approach for managing comment relationships. Each `CommentNode` has a `parentID` field, which is the ID of its parent comment. An older commit showed a more academic recursive approach, where each comment node had a list of its child comments.