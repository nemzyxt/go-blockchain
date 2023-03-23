package blockchain

/*
 * Take data from the block
 * Create a counter (nonce) which starts from 0
 * Generate a hash using the data and the nonce
 * Check whether the hash meets certain requirements
 * Requirements :
	- The first bytes of the hash must contain 0s
 */