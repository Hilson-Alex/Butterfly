package runtime

type BF__MessageContent map[string]interface{}

type BF__MessageInfo struct {
	Content      BF__MessageContent
	SenderName   string
	ReceiverName string
	EventName    string
}

func BF__MessageCreate(sender string, content BF__MessageContent) BF__MessageInfo {
	return BF__MessageInfo{
		Content:    content,
		SenderName: sender,
	}
}
