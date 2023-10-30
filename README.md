# redis-database

## Simple Redis-like Server
Created a simplified version of Redis that supports a subset of its commands, 

The simplified server can handle basic SET, GET, DEL, EXPIRE, KEYS, TTL, ZADD and ZRANGE commands.

## Supported Commands

* SET: Store a key-value pair.
    * ```SET <key> <value>``` 
* GET: Fetch the value associated with a given key.
    * ```GET key``` 
* DEL: Delete a key-value pair.
    * ```DEL key``` 
* EXPIRE: Set expire time for a key-value pair.
    * ```EXPIRE key ttl``` 
* KEYS: Fetch all keys matching the regex.
    * ```KEYS filter``` 
* TTL: Check the expire time for a key-value pair.
    * ```TTL key``` 
* ZADD: Store a key in a sorted set.
    * ```ZADD key score value``` 
* ZRANGE: Fetch the score and value of a given key between min and max score.
    * ```ZRANGE key minscore maxscore``` 


## Getting Started

These instructions will help you get the project up and running on your local machine.

### Prerequisites

- Go (Golang) should be installed on your machine. If not, you can download it from [here](https://go.dev/).

### Installation

1. Clone the repository to your local machine.
   ```bash
   git clone https://github.com/saurabhy27/redis-database.git
2. Change to the project directory.
    ```bash
    cd redis-like-server
3. Run the server.
    ```bash
    go run main.go
4. Start the client by running the following command in Mac/Linux system
    ```bash
    nc <localhostip> 80
5. You can interact with the server by entering Redis-like commands in the client. For example
    ```bash
    redis> SET test test123
    OK
    redis> SET care care123
    OK
    redis> GET test
    test123
    redis> KEYS
    ERR wrong number of arguments for given command
    redis> KEYS *
    test
    care
    redis> EXPIRE test 10
    1
    redis> TTL test
    7
    redis> KEYS *
    care
    redis> ZADD saurabh 10 saurabh10
    1
    redis> ZADD saurabh 12 saurabh12
    1
    redis> ZRANGE saurabh 0 1
    saurabh12  12.000000
    saurabh10  10.000000

