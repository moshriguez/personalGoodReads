# personalGoodReads
An app for keeping track of my books

The app will consist of a Backend rest or grpc api (tbd)
Users of the app will log in and have full CRUD access to the DB, but this will on be applied the the cache. Only a superuser will be able to make actual changes the the DB.

The app will track the books in my personal library. It will keep track of the books I've read. It will also track the books that were lost in the move. 

TODO:
- update update endpoint so partial updates do not clear existing data
- compare grpc and rest api
    - maybe implement both for practice
    - grpc is probably overkill
- add memcache
- implement endpoints:
    - AddBook - add a book to DB; will satisfy book ingestion
    - GetBook - the basic retreival of a single book
    - GetBooks - the basic get all for the books collection
    - UpdateBook - basic update endpoint. To be used for data correction
    - ReadBook - a special update to mark a book as read and record the timestamp
    - DeleteBook - soft deleted a book. Uses update method. entry remains in the DB
- have a faktory worker or some chron job to periodically refresh the cache to help with initial load time
- update endpoints so logged in user only persist data to the DB; unlogged in users will only persist to the cache

Bugs:
- 
