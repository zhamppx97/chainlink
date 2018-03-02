// Package adapters contain the core adapters used by the Chainlink node.
//
// HTTPGet
//
// The HTTPGet adapter is used to grab the JSON data from the given URL.
//  { "type": "HTTPGet", "url": "https://some-api-example.net/api" }
//
// HTTPPost
//
// Sends a POST request to the specified URL and will return the response.
//  { "type": "HTTPPost", "url": "https://weiwatchers.com/api" }
//
// JSONParse
//
// The JSONParse adapter will obtain the value(s) for the given field(s).
//  { "type": "JSONParse", "path": ["someField"] }
//
// EthBytes32
//
// The EthBytes32 adapter will take the given values and format them for
// the Ethereum blockhain.
//  { "type": "EthBytes32" }
//
// EthTx
//
// The EthTx adapter will write the data to the given address and functionSelector.
//   {
//     "type": "EthTx",
//     "address": "0x0000000000000000000000000000000000000000",
//     "functionSelector": "0xffffffff"
//   }
//
package adapters
