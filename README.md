# Find most fertile matron


## Challenge
This project is part of a challenge proposed by MakersPlace.
The challenge is the following: in a determined range of blocks, a defined range of blocks, count the times the 'Birth' event was emitted and find the most fertile matron, that means, the matronId who gave birth to the most Kitties in the contract.

## Solution

Using the geth package (the one used to sync ethereum nodes), we set a connection with an endpoint (Infura), created a smart contract binding for the go application to interact with and ask for logs in batches since we cannot fetch more than 9999 registers from Infura.
In this specific case I made use of go routines to work on the API calls in an asynchronous manner.
That dind't end up working as expected because at some point there may be bottle necks.


## Results

Most fertile matrons: [1083637,  1100747], both with 22 appearances in total between the specified block range.

Number of 'Birth' events that happened 
    between blocks 6607985 and 7028323 is: 186652

Kitty 1083637 has genes[467882257905024579446667743955452078189962217938379332103542434935342116], 
            generation[0], 
            and birthTime[1539134711]

Kitty 1100747 has genes[678562013145546313391979201724185257126131816581653569215158509597308012], 
            generation[1], 
            and birthTime[1539793671]
