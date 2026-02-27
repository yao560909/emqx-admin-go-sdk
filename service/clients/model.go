package clients

type Client struct {
	ProtoName                        string  `json:"proto_name"`
	Reductions                       int64   `json:"reductions"`
	MqueueLen                        int64   `json:"mqueue_len"`
	SendMsgDroppedQueueFull          int64   `json:"send_msg.dropped.queue_full"`
	ExpiryInterval                   int64   `json:"expiry_interval"`
	RecvCnt                          int64   `json:"recv_cnt"`
	SendCnt                          int64   `json:"send_cnt"`
	MqueueMax                        int64   `json:"mqueue_max"`
	Keepalive                        int64   `json:"keepalive"`
	RecvOct                          int64   `json:"recv_oct"`
	SendMsgDroppedExpired            int64   `json:"send_msg.dropped.expired"`
	HeapSize                         int64   `json:"heap_size"`
	SendPkt                          int64   `json:"send_pkt"`
	IsBridge                         bool    `json:"is_bridge"`
	SendMsgQos1                      int64   `json:"send_msg.qos1"`
	AwaitingRelMax                   int64   `json:"awaiting_rel_max"`
	RecvMsgDroppedAwaitPubrelTimeout int64   `json:"recv_msg.dropped.await_pubrel_timeout"`
	SubscriptionsMax                 string  `json:"subscriptions_max"`
	PeerPort                         int64   `json:"peerport"`
	EnableAuthn                      bool    `json:"enable_authn"`
	RecvMsgQos0                      int64   `json:"recv_msg.qos0"`
	CreatedAt                        string  `json:"created_at"`
	SendMsgDroppedTooLarge           int64   `json:"send_msg.dropped.too_large"`
	SendMsgDropped                   int64   `json:"send_msg.dropped"`
	Node                             string  `json:"node"`
	SendMsgQos2                      int64   `json:"send_msg.qos2"`
	MqueueDropped                    int64   `json:"mqueue_dropped"`
	SendMsgQos0                      int64   `json:"send_msg.qos0"`
	Connected                        bool    `json:"connected"`
	ConnectedAt                      string  `json:"connected_at"`
	InflightCnt                      int64   `json:"inflight_cnt"`
	Port                             int64   `json:"port"`
	SubscriptionsCnt                 int64   `json:"subscriptions_cnt"`
	AwaitingRelCnt                   int64   `json:"awaiting_rel_cnt"`
	SendMsg                          int64   `json:"send_msg"`
	InflightMax                      int64   `json:"inflight_max"`
	RecvMsgQos1                      int64   `json:"recv_msg.qos1"`
	Username                         string  `json:"username"`
	ClientID                         string  `json:"clientid"`
	Listener                         string  `json:"listener"`
	ProtoVer                         int64   `json:"proto_ver"`
	RecvPkt                          int64   `json:"recv_pkt"`
	IPAddress                        string  `json:"ip_address"`
	Mountpoint                       *string `json:"mountpoint"`
	RecvMsg                          int64   `json:"recv_msg"`
	RecvMsgQos2                      int64   `json:"recv_msg.qos2"`
	SendOct                          int64   `json:"send_oct"`
	RecvMsgDropped                   int64   `json:"recv_msg.dropped"`
	IsPersistent                     bool    `json:"is_persistent"`
	MailboxLen                       int64   `json:"mailbox_len"`
	CleanStart                       bool    `json:"clean_start"`
}

type Meta struct {
	Count   int  `json:"count"`
	HasNext bool `json:"hasnext"`
	Limit   int  `json:"limit"`
	Page    int  `json:"page"`
}

type Subscription struct {
	Clientid string `json:"clientid"`
	Nl       int    `json:"nl"`
	Node     string `json:"node"`
	Qos      int    `json:"qos"`
	Rap      int    `json:"rap"`
	Rh       int    `json:"rh"`
	Topic    string `json:"topic"`
}

type AuthzCache struct {
	Access      *Access `json:"access"`
	Result      string  `json:"result"`
	Topic       string  `json:"topic"`
	UpdatedTime int64   `json:"updated_time"`
}

type Access struct {
	ActionType string `json:"action_type"`
	Qos        int    `json:"qos"`
}
