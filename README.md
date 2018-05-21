# anagram-search

### Instruction for running this program:
    1. Install Docker(skip if you already have it installed)
    2. unzip and cd into the root directory(`cd /app`)
    3. Start the server: `docker-compose up`
        - This will pull down the images needed to run the program
        NOTE: the server says it's listening on 8080, but the container exposes ports 3000 and 8080, so you can run the ruby tests w/o changing the port
### Documentation for running the optional endpoints
    1. Endpoint that takes a set of words and returns whether or not       they are all anagrams of each other 
        - POST /words/check
        - ex: curl -i -X POST -d '{ "words": ["read", "dear", "dare"] }' http://localhost:8080/words/check
    2. Endpoint to delete a word *and all of its anagrams*
        - DELETE /anagrams/:word
    3. 
### Thoughts 

- I chose Redis as my datastore for its fast read/write speed as well as its support for complex types. 
- I also thought Redis was the right choice for this application as the ultimate goal was to optimize for fast lookups. 
- This lies a tradeoff where the data model is optimized for fast reads, but not as flexible as a relational model, for example. 
- If I had a lot of time to work on this project, I would like to add some persistent storage to back up redis, and when doing ingests, I would like to have some event driven system to sync the databases.      
    

    