package retcode

const (
	//0表示成功
	SUCCESS = 0
	//-1~-9预留
	ERR_PARAM = -1
	//-10~-99表示系统级错误
	ERR_WRONG_REDIS_OPERATE      = -10 //redis操作错误
	ERR_WRONG_MYSQL_OPERATE      = -11 //mysql操作错误
	ERR_WRONG_CACHE_OPERATE      = -12 //memcache操作错误
	ERR_WRONG_SYSTEM_OPERATE     = -13 //系统错误
	ERR_WRONG_FORMAT_JSON        = -14 //json格式化错误
	ERR_WRONG_HTTP_GET_REQUEST   = -15 //http get请求错误
	ERR_WRONG_HTTP_POST_REQUEST  = -16 //http post请求错误
	ERR_WRONG_WECHAT_API_REQUEST = -17 //微信请求接口错误
	ERR_WRONG_PINGPP_API_REQUEST = -18 //ping++支付申请失败
	PARAM_ERROR                  = -19 //参数错误
	//-100~-999表示业务逻辑中的错误，各个项目自定

	//-1000~……表示公共业务逻辑错误
	ERR_NO_LOGIN             = -1001 //用户未登录
	ERR_ACCOUNT              = -1002 //帐号或密码错误
	ERR_ACCESS_DENIED        = -1003 // 没有权限
	ERR_NO_WECHAT_LOGIN      = -1004 //微信用户未登录
	ERR_WECHAT_NO_BIND_PHONE = -1005 //未绑定手机号
	ERR_CAPTCHA_INVALID      = -1006 //验证码不正确
	UPLOAD_ERROR             = -1007 //上传错误
	IMAGE_FORMAT_ERROR       = -1008 //图片格式错误
	FILE_IS_EMPTY            = -1009 //没有上传文件
	NOT_ZIP_FILE             = -1010 //上传的非zip文件
	ZIP_ERROR                = -1011 //解压失败
	NOT_INDEX_HTML_FILE      = -2012 //不存在index.html文件
	WECHAT_LOGIN_ERR         = -2013 //微信登录失败

	ERR_OBJECT_NOT_FOUND = -2000 //数据不存在
	ERR_LOST_ID_ARGUMENT = -2001 //缺少ID
)
