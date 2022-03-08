package pushplus

type SendMsgReq struct {
	Token       string `json:"token"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content"`
	Template    string `json:"template,omitempty"` // 消息模板 html,json,cloudMonitor 默认 html
	Topic       string `json:"topic,omitempty"`    // 群组编码，不填仅发送给自己；channel为webhook时无效
	Channel     string `json:"channel,omitempty"`  // 发送渠道, wechat,webhook,cp 默认 wechat
	Webhook     string `json:"webhook,omitempty"`
	CallbackUrl string `json:"callback_url,omitempty"` // 发送结果回调地址
	Timestamp   string `json:"timestamp,omitempty"`    // 毫秒时间戳。格式如：1632993318000。服务器时间戳大于此时间戳，则消息不会发送
}

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type SendMsgCallbackResp struct {
	ShortCode  string `json:"shortCode"`  // 消息流水号
	SendStatus int    `json:"sendStatus"` // 发送状态；0未发送，1发送中，2发送成功，3发送失败
	Message    string `json:"message"`    // 推送错误内容（如有）
}

type AuthReq struct {
	Token     string `json:"token,omitempty"`     // 用户令牌，同发送消息接口令牌
	SecretKey string `json:"secretKey,omitempty"` // 用户密钥
}

type AuthResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessKey string `json:"accessKey"` // 要获取的 accessKey
		ExpiresIn int64  `json:"expiresIn"` // accessKey 的有效期(秒为单位，7200秒之内的值)；
	} `json:"data"`
}

type PageReq struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}
type Pager struct {
	PageNum  int `json:"pageNum"`  // 当前页码
	PageSize int `json:"pageSize"` // 分页大小
	Total    int `json:"total"`    // 总行数
	Pages    int `json:"pages"`    // 总页数
}

type MsgListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []struct {
			TopicName   string `json:"topicName"`   // 群组名称，一对多消息才有值
			MessageType int    `json:"messageType"` // 消息类型;1-一对一消息,2-一对多消息
			Title       string `json:"title"`       // 消息标题
			ShortCode   string `json:"shortCode"`   // 消息短链码;可用于查询消息发送结果
			Channel     string `json:"channel"`     // 消息发送渠道； wechat-微信公众号,mail-邮件,cp-企业微信应用,webhook-第三方webhook
			UpdateTime  string `json:"updateTime"`  // 更新日期
		} `json:"list"`
	} `json:"data"`
}

type SendMsgResultResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Status       int    `json:"status"`
		ErrorMessage string `json:"errorMessage"`
		UpdateTime   string `json:"updateTime"`
	} `json:"data"`
}

type UserInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OpenId      string  `json:"openId"`      // 用户微信的openId
		UnionId     string  `json:"unionId"`     // 用户微信的unionId
		NickName    string  `json:"nickName"`    // 昵称
		HeadImgUrl  string  `json:"headImgUrl"`  // 头像
		UserSex     int     `json:"userSex"`     // 性别: 0-未设置，1-男，2-女
		Token       string  `json:"token"`       // 用户令牌
		PhoneNumber string  `json:"phoneNumber"` // 手机号
		Email       string  `json:"email"`       // 邮箱
		EmailStatus int     `json:"emailStatus"` // 邮箱验证状态；0-未验证，1-待验证，2-已验证
		Birthday    string  `json:"birthday"`    // 生日
		Points      float64 `json:"points"`      // 用户积分
	} `json:"data"`
}

type UserLimitTimeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		SendLimit     int    `json:"sendLimit"`     // 发送限制状态;1-无限制，2-短期限制，3-永久限制
		UserLimitTime string `json:"userLimitTime"` // 解封时间
	} `json:"data"`
}

type TopicListReq struct {
	PageReq
	Params struct {
		TopicType int `json:"topicType"` // 群组类型;0-我创建的，1-我加入的
	} `json:"params"`
}

type TopicListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []struct {
			TopicId    int    `json:"topicId"`    // 群组编号
			TopicCode  string `json:"topicCode"`  // 群组编码
			TopicName  string `json:"topicName"`  // 群组名称
			CreateTime string `json:"createTime"` // 创建时间
		} `json:"list"`
	} `json:"data"`
}

type TopicDetailResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TopicId        int    `json:"topicId"`
		TopicName      string `json:"topicName"`
		TopicCode      string `json:"topicCode"`
		QrCodeImgUrl   string `json:"qrCodeImgUrl"`
		Contact        string `json:"contact"`
		Introduction   string `json:"introduction"`
		ReceiptMessage string `json:"receiptMessage"`
		CreateTime     string `json:"createTime"`
		TopicUserCount int    `json:"topicUserCount"`
	} `json:"data"`
}
type WebhookListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		PageNum  int `json:"pageNum"`
		PageSize int `json:"pageSize"`
		Total    int `json:"total"`
		Pages    int `json:"pages"`
		List     []struct {
			Id          int    `json:"id"`
			WebhookCode string `json:"webhookCode"`
			WebhookName string `json:"webhookName"`
			WebhookType int    `json:"webhookType"`
			WebhookUrl  string `json:"webhookUrl"`
			CreateTime  string `json:"createTime"`
		} `json:"list"`
	} `json:"data"`
}
