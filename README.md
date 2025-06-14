# Short-URL

URL Shortener made using GO and Redis

## About
This is a simple URL shortner app made using go. Server is made using gin. Redis is used as the DB here. 

Basically the incoming URL is mapped to a hashed random 6 letter string and stored to the redis DB and later the original url is retrieved based on request from client from the shortenised URL.

#### TO DO

+ Track usage: Log visit count
+ add expiry to the short URL
+ url input validation