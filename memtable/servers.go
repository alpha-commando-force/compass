package memtable

type Ops interface {
	// RegisterServerNode server node first register to compass.
	RegisterServerNode(serverName string, nodeIP string, value []byte, timeOut int64) error
	// DeleteServerNode delete offline node.
	DeleteServerNode(serverName string, nodeIP string) error
	// GetServerNode get server's node.
	GetServerNode(serverName string, nodeIP string) ([]byte, error)
	// ListServerNodes list all nodes in server.
	ListServerNodes(serverName string) (map[string][]byte, error)
	// RegisterServer server first register to compass.
	RegisterServer(serverName interface{}) error
	// HeartBeat server's node report heart beat periodically.
	HeartBeat(serverName string, nodeIP string) error
}
