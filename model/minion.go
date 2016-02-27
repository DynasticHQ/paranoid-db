package model

type Minion struct {
	AgentId     string        `json: "agent_id,omitempty" gorethink:"agent_id,omitempty"`
	Name        string        `json: "name" gorethink:"name"`
	NeedsReboot bool          `json: "needs_reboot" gorethink: "needs_reboot"`
	Rebooted    bool          `json: "rebooted" gorethink: "rebooted"`
	Secrets     MinionSecrets `json: "secrets" gorethink: "secrets"`
	SystemType  string        `json: "system_type" gorethink: "system_type"`
	Version     string        `json: "version" gorethink: "version"`
}

type MinionSecrets struct {
	PrivateKey string `json: "private_key" gorethink: "private_key"`
	PublicKey  string `json: "public_key" gorethink: "public_key"`
}
