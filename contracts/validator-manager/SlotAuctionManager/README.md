# Slot Auction Validator Manager (SAVM)

The SAVM is a new form of proof for becoming a validator on an avalanche chain. it involves the owner of a chain to start an auction where N validator slots are able to be bid on, each with equal weight W. Once the auction is completed, the winners will automatically have their Node `initialized` by the `ValidatorManager`. 

# Type of auction

The SAVM is a slightly modified [`Second Price Auction`] (https://en.wikipedia.org/wiki/Generalized_second-price_auction) Where the significant differences are the ability to bid multiple times with the same NodeID, and the inability to place a bid of equal value. 

# No duplicate bids

The reason for no duplicate bids comes from 2 separate issues
1. The `Heap` library from openzeppelin only holds `uint256` values, meaning to efficiently store and sift through bidding info we need to have a map that links BidAmount => BidderInformation, meaning no duplicate keys
2. In the event of a tie where there is one `ValidatorSlot` being auctioned, the question of who gains the slot is not defined and is usually kept to a value which of course we cannot do in solidity, therefore whoever bids first will get to keep their bid and any other bid with the same amount will have to bid a different amount

# Allowing NodeIDs to bid multiple times

In a normal second price auction, it is only required to bid once, however in the case of the SAVM, if we only allowed users to bid once, it could lead to malicious users lowballing validator slots using other users nodeIDs, therefore it is easier to allow any NodeID to bid as much as they want, the same NodeID will be allowed to bid assuming they are not already in a winning position, this makes it so a bidder doesn't increase their bid for no reason, and also helps with issues of having to update that bid value in the heap, since there is no guarantee it will be at the head.

# SAVM auction implementation
The SAVM is first initalized by the owner, passing in arguments for the amount of validator slots being auctioned off, the validator weights, the validators life time, and the minimum length of the auction, The auction uses a `Heap` library from openZeppelin that holds `uint256` values. 

At any given time the heap will be at or less than size N where N is the number of validators being auctioned. Until the `Heap` reaches it's max capacity, all bids will be considered a winner and be placed into the `Heap`. At capacity, the SAVM will only consider bids that are more than the smallest winning bid, in that case, the smallest qualifying bid will be removed and replaced with the new bid. Whenever a bid is removed from the heap, it becomes the second price for the smallest qualifying bid

Once the end time has been reached, the owner of the auction will complete it, in which the `Heap` will be emptied, automatically `initializing` each of them as a validator

# Edge cases

# No bidders
If no one bids then no nodes will become validator

# Not enough bidders
If there are not enough bidders to give the lowest qualifying bid a second price, then the lowest bidder will pay the full price they bid

# Validator is unable to be initialized
This is in the case where one of the N nodes who have won the auction cannot be initialized and it causes the `completeAuction` to fail

# Validator churn
Two options I can think of
1. If the auction owner were to choose a combination of weights and slots that would cause churn, then don't allow it.
2. A separate function that allows Validators to be initialized in smaller batches when the auction is complete, assuming it would cause validator churn

# Reward system/tokenomics 
This is up to discussion, but from my understanding the goal of the SAVM is to make it so the bidder needs to come up with an estimated value of how much they believe they will gain from being a validator through "Micro market manipulation". 

The tokens are not held to return later or else its just proof of stake with some spice. What happens to the tokens once they have been used to buy the Validator slot is up to discussion, the only thing I highly suggest is that the tokens gained from the auction should be used as gas to complete the end of auction due to it most likely being a gas heavy operation depending on the number of slots, especially if we end up burning the tokens. 