# API Restful - Simple CRUD

**Go Language - Gorilla library - No database (handled in memory as a map)**

This is a simple CRUD of "Messages". There are four endpoints available:

(GET)    /messages                 -> fetches all messages in memory

(POST)   /messages/create          -> creates a message

(PUT)    /messages/update/{id}     -> updates the message with certain id

(DELETE) /messages/delete/{id}     -> deletes the message with certain id

*Note: you should first create messages in order to perform any other action. Data is handled in memory during runtime*
