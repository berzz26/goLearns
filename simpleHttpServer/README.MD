This is my implementation of a very basic http server without using any (fancy)external lib in go using the net/http package.
It basically just has a main func and a handler that handles a get and a post request.
The data storage and retrieval is done in a handler/users.json file (yeah, learnt file io asw with this) 