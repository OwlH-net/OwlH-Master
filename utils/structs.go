package utils

type NodeEnroll struct {
    Node     NodeData     `json:"node"`
    Group    GroupData    `json:"group"`
    Suricata SuricataData `json:"suricata"`
    Stap     StapData     `json:"stap"`
}

type NodeData struct {
    IP           string `json:"ip"`
    Name         string `json:"name"`
    Port         int    `json:"port"`
    NodeUser     string `json:"nodeuser"`
    NodePassword string `json:"nodepassword"`
    Force        bool   `json:"force"`
}

type GroupData struct {
    UUID string `json:"uuid"`
}

type SuricataData struct {
    Interface  string `json:"interface"`
    Bpf        string `json:"bpf"`
    BpfFile    string `json:"bpffile"`
    ConfigFile string `json:"configfile"`
    Ruleset    string `json:"ruleset"`
    Name       string `json:"name"`
    Status     string `json:"status"`
}

type StapData struct {
    Interface string `json:"interface"`
    Cert      string `json:"cert"`
    Port      string `json:"port"`
    Type      string `json:"type"`
    Name      string `json:"name"`
    Status    string `json:"status"`
}

/////////////////////////////////////////////
type EnrollNewNodeStruct struct {
    Node     NewNodeData  `json:"node"`
    Group    []string     `json:"group"`
    Suricata SuricataData `json:"suricata"`
    Tags     []string     `json:"tags"`
    Orgs     []string     `json:"orgs"`
}
type NewNodeData struct {
    UUID     string `json:"uuid"`
    IP       string `json:"ip"`
    Name     string `json:"name"`
    Port     string `json:"port"`
    NodeUser string `json:"nodeuser"`
    NodePass string `json:"nodepass"`
}

// type GroupArray struct {
//     UUID    []string `json:"uuid"`
// }

type WazuhDetails struct {
    Path    bool   `json:"path"`
    Bin     bool   `json:"bin"`
    Running bool   `json:"running"`
    Name    string `json:"name"`
    ID      string `json:"id"`
    Ip      string `json:"ip"`
}
