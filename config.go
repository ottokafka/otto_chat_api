package main

// initialize redis users

var errset = RedisClient.HSet("users", "init", "{user:init,pin:1234}").Err()
