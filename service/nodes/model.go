package nodes

type Node struct {
	Node            string `json:"node"`
	Connections     int64  `json:"connections"`
	LiveConnections int64  `json:"live_connections"`

	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`

	MaxFDs      int64  `json:"max_fds"`
	MemoryTotal string `json:"memory_total"`
	MemoryUsed  string `json:"memory_used"`

	NodeStatus       string `json:"node_status"`
	OTPRelease       string `json:"otp_release"`
	ProcessAvailable int64  `json:"process_available"`
	ProcessUsed      int64  `json:"process_used"`

	Uptime  int64  `json:"uptime"`
	Version string `json:"version"`
	Edition string `json:"edition"`

	SysPath string `json:"sys_path"`
	LogPath string `json:"log_path"`

	Role string `json:"role"`
}

type NodeMetrics struct {
	AuthorizationDeny          int64 `json:"authorization.deny"`
	AuthenticationFailure      int64 `json:"authentication.failure"`
	DeliveryDroppedTooLarge    int64 `json:"delivery.dropped.too_large"`
	MessagesSent               int64 `json:"messages.sent"`
	PacketsDisconnectSent      int64 `json:"packets.disconnect.sent"`
	BytesSent                  int64 `json:"bytes.sent"`
	PacketsConnackAuthError    int64 `json:"packets.connack.auth_error"`
	DeliveryDroppedNoLocal     int64 `json:"delivery.dropped.no_local"`
	OverloadProtectionDelayOK  int64 `json:"overload_protection.delay.ok"`
	PacketsPubcompInuse        int64 `json:"packets.pubcomp.inuse"`
	ClientDisconnected         int64 `json:"client.disconnected"`
	PacketsSubackSent          int64 `json:"packets.suback.sent"`
	PacketsPubackSent          int64 `json:"packets.puback.sent"`
	AuthorizationAllow         int64 `json:"authorization.allow"`
	PacketsSubscribeReceived   int64 `json:"packets.subscribe.received"`
	OverloadProtectionGC       int64 `json:"overload_protection.gc"`
	PacketsPublishError        int64 `json:"packets.publish.error"`
	ClientSubscribe            int64 `json:"client.subscribe"`
	DeliveryDropped            int64 `json:"delivery.dropped"`
	MessagesQos2Received       int64 `json:"messages.qos2.received"`
	MessagesQos1Received       int64 `json:"messages.qos1.received"`
	MessagesQos0Received       int64 `json:"messages.qos0.received"`
	AuthenticationSuccess      int64 `json:"authentication.success"`
	PacketsConnackError        int64 `json:"packets.connack.error"`
	OverloadProtectionHiber    int64 `json:"overload_protection.hibernation"`
	PacketsDisconnectReceived  int64 `json:"packets.disconnect.received"`
	PacketsPingrespSent        int64 `json:"packets.pingresp.sent"`
	PacketsPingreqReceived     int64 `json:"packets.pingreq.received"`
	PacketsUnsubscribeReceived int64 `json:"packets.unsubscribe.received"`
	PacketsPubcompMissed       int64 `json:"packets.pubcomp.missed"`
	PacketsPubackMissed        int64 `json:"packets.puback.missed"`
	DeliveryDroppedExpired     int64 `json:"delivery.dropped.expired"`
	MessagesReceived           int64 `json:"messages.received"`
	PacketsConnectReceived     int64 `json:"packets.connect.received"`
	AuthenticationSuccessAnon  int64 `json:"authentication.success.anonymous"`
	MessagesDroppedAwaitPubrel int64 `json:"messages.dropped.await_pubrel_timeout"`
	PacketsPubrelReceived      int64 `json:"packets.pubrel.received"`
	PacketsPubrecReceived      int64 `json:"packets.pubrec.received"`
	PacketsPubackReceived      int64 `json:"packets.puback.received"`
	PacketsSent                int64 `json:"packets.sent"`
	PacketsReceived            int64 `json:"packets.received"`
	BytesReceived              int64 `json:"bytes.received"`
	MessagesDroppedNoSubs      int64 `json:"messages.dropped.no_subscribers"`
	PacketsUnsubscribeError    int64 `json:"packets.unsubscribe.error"`
	PacketsAuthSent            int64 `json:"packets.auth.sent"`
	PacketsPublishAuthError    int64 `json:"packets.publish.auth_error"`
	ClientConnack              int64 `json:"client.connack"`
	PacketsPubackInuse         int64 `json:"packets.puback.inuse"`
	PacketsPubcompSent         int64 `json:"packets.pubcomp.sent"`
	PacketsPubcompReceived     int64 `json:"packets.pubcomp.received"`
	MessagesPublish            int64 `json:"messages.publish"`
	ClientUnsubscribe          int64 `json:"client.unsubscribe"`
	PacketsPubrecInuse         int64 `json:"packets.pubrec.inuse"`
	MessagesDelivered          int64 `json:"messages.delivered"`
	MessagesForward            int64 `json:"messages.forward"`
	PacketsPubrelMissed        int64 `json:"packets.pubrel.missed"`
	PacketsPublishReceived     int64 `json:"packets.publish.received"`
	PacketsConnackSent         int64 `json:"packets.connack.sent"`
	DeliveryDroppedQos0        int64 `json:"delivery.dropped.qos0_msg"`
	MessagesDelayed            int64 `json:"messages.delayed"`
	MessagesDropped            int64 `json:"messages.dropped"`
	PacketsPublishDropped      int64 `json:"packets.publish.dropped"`
	AuthorizationSuperuser     int64 `json:"authorization.superuser"`
	AuthorizationMatchedAllow  int64 `json:"authorization.matched.allow"`
	ClientConnect              int64 `json:"client.connect"`
	MessagesAcked              int64 `json:"messages.acked"`
	MessagesQos2Sent           int64 `json:"messages.qos2.sent"`
	MessagesQos1Sent           int64 `json:"messages.qos1.sent"`
	MessagesQos0Sent           int64 `json:"messages.qos0.sent"`
	PacketsSubscribeError      int64 `json:"packets.subscribe.error"`
	ClientAuthorize            int64 `json:"client.authorize"`
	SessionCreated             int64 `json:"session.created"`
	PacketsPubrecMissed        int64 `json:"packets.pubrec.missed"`
	ClientAuthAnonymous        int64 `json:"client.auth.anonymous"`
	PacketsPublishInuse        int64 `json:"packets.publish.inuse"`
	AuthorizationCacheHit      int64 `json:"authorization.cache_hit"`
	SessionTerminated          int64 `json:"session.terminated"`
	SessionResumed             int64 `json:"session.resumed"`
	SessionTakenover           int64 `json:"session.takenover"`
	AuthorizationMatchedDeny   int64 `json:"authorization.matched.deny"`
	ClientAuthenticate         int64 `json:"client.authenticate"`
	SessionDiscarded           int64 `json:"session.discarded"`
	PacketsPubrelSent          int64 `json:"packets.pubrel.sent"`
	PacketsPubrecSent          int64 `json:"packets.pubrec.sent"`
	PacketsPublishSent         int64 `json:"packets.publish.sent"`
	OverloadProtectionNewConn  int64 `json:"overload_protection.new_conn"`
	DeliveryDroppedQueueFull   int64 `json:"delivery.dropped.queue_full"`
	AuthorizationNomatch       int64 `json:"authorization.nomatch"`
	OverloadProtectionDelayTO  int64 `json:"overload_protection.delay.timeout"`
	ClientConnected            int64 `json:"client.connected"`
	PacketsAuthReceived        int64 `json:"packets.auth.received"`
	PacketsUnsubackSent        int64 `json:"packets.unsuback.sent"`
	PacketsSubscribeAuthError  int64 `json:"packets.subscribe.auth_error"`
}

type NodeStats struct {
	SubscriptionsSharedMax   int64 `json:"subscriptions.shared.max"`
	SubscriptionsMax         int64 `json:"subscriptions.max"`
	SubscribersMax           int64 `json:"subscribers.max"`
	TopicsCount              int64 `json:"topics.count"`
	ChannelsCount            int64 `json:"channels.count"`
	SubscriptionsCount       int64 `json:"subscriptions.count"`
	SuboptionsMax            int64 `json:"suboptions.max"`
	TopicsMax                int64 `json:"topics.max"`
	ConnectionsMax           int64 `json:"connections.max"`
	RetainedCount            int64 `json:"retained.count"`
	LiveConnectionsMax       int64 `json:"live_connections.max"`
	SubscriptionsSharedCount int64 `json:"subscriptions.shared.count"`
	SuboptionsCount          int64 `json:"suboptions.count"`
	SessionsCount            int64 `json:"sessions.count"`
	ChannelsMax              int64 `json:"channels.max"`
	RetainedMax              int64 `json:"retained.max"`
	SessionsMax              int64 `json:"sessions.max"`
	DelayedMax               int64 `json:"delayed.max"`
	DelayedCount             int64 `json:"delayed.count"`
	LiveConnectionsCount     int64 `json:"live_connections.count"`
	SubscribersCount         int64 `json:"subscribers.count"`
	ConnectionsCount         int64 `json:"connections.count"`
}
