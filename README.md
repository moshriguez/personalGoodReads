# personalGoodReads
An app for keeping track of my books

The app will consist of a Backend rest or grpc api (tbd).
Users of the app will log in and have full CRUD access to the data, but this will only be applied the the cache. Only a superuser will be able to make actual changes the the DB.

The app will track the books in my personal library. It will keep track of the books I've read. It will also track the books that were lost in the move. 

TODO:
- update update endpoint so partial updates do not clear existing data
- compare grpc and rest api
    - we will implement both for practice
    - grpc is probably overkill - but I'd like to do this since I work with this at work and this will demonstrate that I know how to use GRPC
- add memcache
- implement endpoints:
    - AddBook - add a book to DB; will satisfy book ingestion
    - GetBook - the basic retreival of a single book
    - GetBooks - the basic get all for the books collection
    - UpdateBook - basic update endpoint. To be used for data correction
    - ReadBook - a special update to mark a book as read and record the timestamp
    - DeleteBook - soft deleted a book. Uses update method. entry remains in the DB
        - currently implemented as a DELETE endpoint. need to create/update to soft delete
        - may want to keep for data correction
- have a faktory worker or some chron job to periodically refresh the cache to help with initial load time
- update endpoints so logged in user only persist data to the DB; unlogged in users will only persist to the cache

Bugs:
- does not launch with docker
- update endpoint updates only the fields passed in, but will not allow bool fields be set back to false once set to true.
    - could hold these values as strings in db and then convert? would need some validation before hitting db. - seems hacky
