# Slot Auction Validator Manager (SAVM)

The SAVM is a new form of proof for becoming a validator on an avalanche chain. it involves the owner of a chain to start an auction where N validator slots are able to be bid on, each with equal weight W. Once the auction is completed, the winners will automatically have their Node `initialized` by the `ValidatorManager`. 

# Type of auction

The SAVM is a slightly modified [`Second Price Auction`] (https://en.wikipedia.org/wiki/Generalized_second-price_auction) Where the significant differences are the ability to bid multiple times with the same NodeID, and the inability to place a bid of equal value. 

# SAVM auction implementation
The SAVM has a decentralized approach where any user is able to start and end an auction given the right conditions are met. Starting an auction takes the `auctionSettings` determined by the admin (`validatorSlots`, `weight`, `auctionDuration`, `validatorDuration`, `minimumBid`) and automatically starts an auction with as many open slots as possible with validator churn restrictions.

At any given time the heap will be at or less than size N where N is the number of validators being auctioned. Until the `Heap` reaches it's max capacity, all bids will be considered a winner and be placed into the `Heap`. At capacity, the SAVM will only consider bids that are greater than the smallest winning bid to remove the ability of a user increasing the second price , in that case, the smallest qualifying bid will be removed and replaced with the new bid. Whenever a bid is removed from the heap, it becomes the second price for the smallest qualifying bid.

The auction can then be ended by any user, given that `auctionEndtime` has been reached, which is the timestamp of the auction beginning added to `auctionDuration` which is determined by the admin. The heap will be emptied, sending back any funds 

# No duplicate bids

The reason for no duplicate bids comes from 2 separate issues
1. The `Heap` library from openzeppelin only holds `uint256` values, meaning to efficiently store and sift through bidding info we need to have a map that links BidAmount => BidderInformation, meaning no duplicate keys
2. In the event of a tie where there is one `ValidatorSlot` being auctioned, the question of who gains the slot is not defined and is usually kept to a value which of course we cannot do in solidity, therefore whoever bids first will get to keep their bid and any other bid with the same amount will have to bid a different amount

# Allowing NodeIDs to bid multiple times

In a normal second price auction, it is only required to bid once, however in the case of the SAVM, if we only allowed users to bid once, it could lead to malicious users lowballing validator slots using other users nodeIDs, therefore it is easier to allow any NodeID to bid as much as they want, the same NodeID will be allowed to bid assuming they are not already in a winning position, this makes it so a bidder doesn't increase their bid for no reason, and also helps with issues of having to update that bid value in the heap, since there is no guarantee it will be at the head.

# Second price differences
In a normal second price auction where N items of equal value are being auctioned off, the winning bidders usually only pay the most expensive losing bid, however in the SAVM, even though each slot is worth the same amount, each winning bidder will pay the amount of the next winning bidder. The reason behind this is because the MEV is not equal for each validator, the more liquid funds a node owner has, the more money they can extract from a block. Thus making the value of the slot much more than someone with half the funds. With this implementation it will lead to users with bigger cash flows to end up paying more, while also still making a profit larger than other winners who have smaller liquidity. 

# Edge cases

# No bidders
If no one bids then no nodes will become validator

# Not enough bidders
If there are not enough bidders to give the lowest qualifying bid a second price, then the lowest bidder will pay the minimumBid

# Validator is unable to be initialized
This is in the case where one of the N nodes who have won the auction cannot be initialized and it causes the `completeAuction` to fail, In the case that this happens, there will be a safeEndSlotAuction which doesn't register anyone as a validator, but instead adds them to a mapping of nodeIDs that are able to become validators.

# Validator churn
Validator churn should not come up, as when an auction starts, the smart contract will automatically calculate how many validators it can auction off without going over the maximum churn percentage. On top of that, when the auction ends validatorChurn for the current period will be calculated, if churn would be triggered then a safe auction end will be used.

# Reward system/tokenomics 
No rewards will be given for winning an auction, the only way validators should be able to get any currency is by extracting the MEV from the blocks they mine. A higher weight correlates to a higher return with MEV since it means they will mine more. This leads to a judgement call by the bidder to try and predict how much MEV they can gain with the weight of the validators being auctioned off 

# Important notes
The SAVM could be the only validator manager on an L1, but it is highly recommended against it. The possibility of an L1 takeover, especially on smaller ones, is too high, since it is not possible to detect who owns what nodeIDs. Therefore it is recommended to run the SAVM alongside a separate validator manager, specifically staking which will allow plenty of other validators to validate the L1.

# Things to fix/Implement
Currently the SAVM does not support pipelining auctions, this is where the auction will end before the slots are open, in preparation for the old slots to be removed. The current implementation of the SAVM has that in mind.

There is currently no incentive for a validator to remove themselves, while it is possible for someone else to remove the validator after their time period is over, it is a better idea to make the owners remove themselves if possible.