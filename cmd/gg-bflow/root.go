package main

import (
	"github.com/spf13/cobra"
)

const longDesc = `
GG-BFlow will behave Hub-like/Messaging-like Protocol/Streaming-like Protocol for streaming your buffer data to your client.
GG-BFlow utilizing some of the technologies, such as:
1. [gRPC](https://grpc.io/)
2. [memcached](https://memcached.org/)
3. [mongodb](https://www.mongodb.com/)
`

var rootCmd = &cobra.Command{
	Use:   "gg-bflow",
	Short: "GG BFlow Stream like-Protocol",
	Long:  longDesc,
}
