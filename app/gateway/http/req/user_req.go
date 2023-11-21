package req

type RegisterReq struct {
	UserName string `form:"user_name" binding:"required"` // 必填
	Pwd      string `form:"password" binding:"required"`
	Phone    string `form:"phone" binding:"required,e164"` //手机号预定义格式 例如+8612345678910
}

type LoginReq struct {
	UserName string `form:"user_name" ` // 必填
	Pwd      string `form:"password" binding:"required"`
	Phone    string `form:"phone" ` //手机号预定义格式 例如+8612345678910
}

type AddBackPack struct {
	UserId  int64 `form:"user_id" binding:"required"` // 必填
	GoodsId int64 `form:"goods_id" binding:"required"` // 必填
	Amount  int64 `form:"amount"`
	LastPrice int64 `form:"last_price"` // 必填
}
