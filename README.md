# go computer science playground

## Algorithms
### Distributed Computing

[doc](./pkg/distributed)

#### Database

- [ ] MultiVersion Concurrency Control
- [ ] Gossip Protocol
- [ ] Atomic Commit

#### Model

- [ ] message passing model

#### Leader Election

#### Consensus

- [ ] Paxos
- [ ] Raft

#### Mutex (Mutual exclution)
- [x] spin lock
- [x] ticket lock

#### Cache
- [x] write through/around/back

##### Cache Eviction Policies
- [x] fifo (Firstly In, First Out)
- [x] lifo (Last In, First out)
- [x] lru (Least Recently Used)
- [x] tlru (Time aware Least Recently Used)
- [ ] mru (Most Recently Used)
- [ ] plru (Pseudo LRU)
- [ ] rr (Random Replacement)
- [ ] slru (Segmented LRU)
- [ ] lfu (Least Frequently Used)
- [ ] lfuda (LFU with Dynamic Aging)
- [ ] lirs (Low Inter-reference Recency Set)
- [ ] clockpro
- [ ] arc (Adaptive Replacement Cache)
- [ ] ac (Adaptive Climb)
- [ ] car (Adaptive Climb)
- [ ] mq (Multi Queue)
- [ ] lfd (Longest Forward Distance)
- [ ] fwf (Flush When Full)

#### Rate Limiter
- [x] array push
- [x] increment
- see https://redis.io/commands/incr

#### Semaphore
- [x] counting
- [ ] weighted

### Sort
- [x] bubble sort
- [x] bucket sort
- [x] heap sort
- [x] merge sort
- [x] quick sort
- [x] insertion sort
- [x] introspective sort
- [ ] tim sort

### Search
- [x] Linear Search
- [x] Binary Search
- [x] Jump Search

### Coding
- [x] lz77
- [x] huffman
- [x] runlength
- [ ] lz78
- [ ] lzw
- [x] base64

### Diff
- [ ] dp
- [ ] myers's diff

### Math

basic functions
- [x] factorial
- [x] combination
- [x] permutation
- [x] absolute
- [x] square
- [x] ceil
- [x] floor
- [x] sqrt
- [x] power
- [x] logarithm
- [x] mod
- [x] exp
- [x] sin
- [x] cos
- [x] tan
- [x] sinc
- [x] arcsin
- [x] arccos
- [x] arctan
- [x] almost equal
  - [x] epsilon comparison
  - [ ] unit in the last place: https://en.wikipedia.org/wiki/Unit_in_the_last_place

complex number
- [ ] test
- [x] add, sub, mul, div
- [x] abs, arg
- [x] reciprocal, conjugate
- [x] euler
- [x] equal

interpolation
- [x] linear
- [x] cosine

numerical calculation
- [x] newton
- [ ] quasi newton
- [ ] BFGS: Broyden–Fletcher–Goldfarb–Shanno
- [ ] Jacobi method
- [ ] Gauss-Seidel method
- [ ] Successive Over-Relaxation
- [ ] gaussian elimination
- [ ] euler
- [ ] Leap-Frog
- [ ] Runge–Kutta

descrete logarithm problem (dlp)
- [x] greedy

elliptic curve dlp (ecdlp)
- [ ] hoge

### Optimization
- [ ] Simulated Annealing
- [ ] genetic algorithm
- [ ] ant colony optimization
- [ ] Particle Swarm Optimization

### Statistics
- [ ] Metropolis-Hastings
- [ ] Hamiltonian MC
- [ ] Gibbs sampling

### Pseudorandom
- [x] midsquare
- [x] linear congruential generators
- [x] gfsr (generalized feedback shift register)
- [x] twisted gfsr
- [ ] mersenne twister
- [x] xorshift
- [ ] Blum-Blum-Shub
- [x] perlin noise

### Version solving
- [ ] PubGrub

### Cryptology
- [ ] enigma

### Load balancing
- [ ] round robin
- [ ] weighted round robin
- [ ] least connection
- [ ] Agent-Based Adaptive Load Balancing
- [ ] Chained Failover
- [ ] Weighted Response Time
- [ ] Source IP Hash

## Hash

uncryptographic hash
- [x] fnv

cryptographic hash
- [x] md5
- [ ] sha-1
- [ ] ripemd-160
- [ ] bcrypt
- [ ] whirlpool
- [ ] sha-2
- [ ] sha-3

## Data structures

### List
- [x] doubly linked list
- [x] queue
- [x] stack

### Trees
- [x] binary search tree
- [x] heap
- [x] treap
- [ ] btree
- [ ] Log-structured merge-tree

### hashtable
- [x] robin hood hashing
- [x] consistent-hashing with treap (get operation: O(lon N))
- [x] cockoo hashing

## Succinct Data Struture
- [x] bit vector
- [ ] wavelet tree
- [ ] wavelet matrix

## Probabilistic Data Structures
- [x] bloom filter
  - wiredtiger: https://github.com/wiredtiger/wiredtiger/wiki/LSMTrees-Bloom
  - Building a Better Bloom Filter: https://www.eecs.harvard.edu/~michaelm/postscripts/rsa2008.pdf
- [x] counting filter
- [ ] quotient filter
- [ ] cuckoo filter
- [ ] xor filter

## Etc
- [x] bits

## References
- https://en.wikipedia.org/wiki/List_of_data_structures
- https://en.wikipedia.org/wiki/List_of_algorithms
- https://en.wikipedia.org/wiki/List_of_algorithm_general_topics
- https://en.wikipedia.org/wiki/List_of_terms_relating_to_algorithms_and_data_structures
- https://en.wikipedia.org/wiki/List_of_computability_and_complexity_topics
- https://www.geeksforgeeks.org/
