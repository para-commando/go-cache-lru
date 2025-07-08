# Go LRU Cache

This project implements a simple Least Recently Used (LRU) cache in Go using a combination of a doubly linked list (queue) and a hash map for efficient O(1) operations. The cache now supports a configurable maximum size and automatic eviction of the least recently used item when full.

## Features
- O(1) time complexity for cache insertions and lookups
- Doubly linked list to maintain usage order
- Hash map for fast key-to-node lookups
- **Configurable maximum cache size** (default: 5)
- **Automatic eviction** of least recently used items when the cache exceeds its size
- **Display method** to visualize the current cache contents

## Structure
- **Node**: Represents an entry in the cache (doubly linked list node)
- **Queue**: Doubly linked list to track the order of usage
- **Hash**: Map from string keys to nodes for fast access
- **Cache**: Combines the queue and hash map, and provides methods to check, add, remove, and display entries

## Usage

1. **Clone the repository**
2. **Run the program**:
   ```sh
   go run main.go
   ```
   You should see output like:
   ```
   Adding Radio-Check to cache
   Cache: {Radio-Check}
   Cache size: 1
   Adding marine to cache
   Cache: {marine}<-->{Radio-Check}
   Cache size: 2
   ...
   ```

## Example
The main function demonstrates how to use the cache:

```go
cache := NewCache()
for _, word := range []string{"Radio-Check", "marine", "commando", "Do you copy ?", "Over"} {
    cache.Check(word)
    cache.Display()
}
```

### Display Output
The `Display` method prints the cache contents from most recently used to least recently used, along with the current cache size.

Example output:
```
Cache: {Over}<-->{Do you copy ?}<-->{commando}<-->{marine}<-->{Radio-Check}
Cache size: 5
```

If you add more than 5 items, the least recently used item is evicted automatically.

## Customization
- To change the cache size, modify the `SIZE` constant in `main.go`:
  ```go
  const SIZE = 5 // maximum size of the cache
  ```

## Notes
- The cache uses a doubly linked list for efficient reordering and eviction.
- The `Check` method moves accessed items to the front (most recently used position).
- The `Display` method helps visualize the cache state after each operation.

## License
MIT 