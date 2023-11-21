package req

type Chat_msg struct {
	UserId     int64  `json:"userId"`
	TargetId   int64  `json:"targetId"`
	Type       int    `json:"type"`  //发送类型  1私聊  2群聊  3心跳
	Media      int    `json:"media"` //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string `json:"content"`
	CreateTime uint64
	ReadTime   uint64
}

type Create_community struct {
	OwnerId  int64  `json:"ownerId"`
	Name     string `json:"name"`
	LimitNum int64  `json:"limitNum"`
	Desc     string `json:"desc"`
}
