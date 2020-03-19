package main

type shard struct {
	ShardID  int
	Reserved bool
	Active   bool
	Server   bool
}

type gatewayresponse struct {
	URL               string
	Shards            int
	SessionStartLimit struct {
		Total      int
		Remaining  int
		ResetAfter int
	}
}
