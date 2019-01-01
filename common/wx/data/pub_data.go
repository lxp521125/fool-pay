package data

type ChargePub struct {
	// 基本数据
	Appid     string `json:"appid" form:"appid"  xml:"appid,omitempty,cdata"`
	MchId     string `json:"mch_id" form:"mch_id"  xml:"mch_id,omitempty,cdata"`
	NonceStr  string `json:"nonce_str" form:"nonce_str"  xml:"nonce_str,omitempty,cdata"`
	SignType  string `json:"sign_type" form:"sign_type"  xml:"sign_type,omitempty,cdata"`
	FeeType   string `json:"fee_type" form:"fee_type"  xml:"fee_type,omitempty,cdata"`
	NotifyUrl string `json:"notify_url" form:"notify_url"  xml:"notify_url,omitempty,cdata"`
	TradeType string `json:"trade_type" form:"trade_type"  xml:"trade_type,omitempty,cdata"` //设置APP支付
	LimitPay  string `json:"limit_pay" form:"limit_pay"  xml:"limit_pay,omitempty,cdata"`    // 指定不使用信用卡
	// 业务数据
	DeviceInfo string `json:"device_info" form:"device_info"  xml:"device_info,omitempty,cdata"`
	Body       string `json:"body" form:"body"  xml:"body,omitempty,cdata"`
	//detail string `json:"detail" form:"detail"  xml:"detail,omitempty,cdata"` JSON_UNESCAPED_UNICODE);
	Attach         string `json:"attach" form:"attach"  xml:"attach,omitempty,cdata"`
	OutTradeNo     string `json:"out_trade_no" form:"out_trade_no"  xml:"out_trade_no,omitempty,cdata"`
	TotalFee       string `json:"total_fee" form:"total_fee"  xml:"total_fee,omitempty,cdata"`
	SpbillCreateIp string `json:"spbill_create_ip" form:"spbill_create_ip"  xml:"spbill_create_ip,omitempty,cdata"`
	TimeStart      string `json:"time_start" form:"time_start"  xml:"time_start,omitempty,cdata"`
	TimeExpire     string `json:"time_expire" form:"time_expire"  xml:"time_expire,omitempty,cdata"`
	Openid         string `json:"openid" form:"openid"  xml:"openid,omitempty,cdata"`
}
