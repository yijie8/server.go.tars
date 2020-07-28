package tbsdk

type DpXgList []struct {
	UserID     int    `json:"user_id"`
	ShopTitle  string `json:"shop_title"`
	SellerNick string `json:"seller_nick"`
	PictUrl    string `json:"pict_url"`
	// ShopType   string `json:"-"`
	// ShopUrl    string `json:"-"`
}

type CpXgList struct {
	NumIid       int64  `json:"num_iid"`
	Title        string `json:"title"`
	PictUrl      string `json:"pict_url"`
	ReservePrice string `json:"reserve_price"`  // 市场价
	ZkFinalPrice string `json:"zk_final_price"` // 实价
	Nick         string `json:"nick"`
	SellerId     int    `json:"seller_id"`
	Volume       int    `json:"volume"` // 销量
}

type CpOnly struct {
	CatName string `json:"cat_name"`
	NumIid  int64  `json:"num_iid"`
	Title   string `json:"title"`
	PictUrl string `json:"pict_url"`
	// SmallImages  struct{ String []struct{ url string } } `json:"small_images"`
	ReservePrice           string `json:"reserve_price"`
	ZkFinalPrice           string `json:"zk_final_price"`
	UserType               int    `json:"-"` // 没用
	Provcity               string `json:"-"` // 城市 没用
	ItemUrl                string `json:"-"`
	SellerId               int    `json:"seller_id"`
	Volume                 int    `json:"volume"`
	Nick                   string `json:"nick"`
	CatLeafName            string `json:"cat_leaf_name"`
	ShopDsr                string `json:"shop_dsr"`                  // 店铺dsr 评分
	Ratesum                int    `json:"ratesum"`                   // 卖家等级
	FreeShipment           bool   `json:"free_shipment"`             // 是否包邮
	PresaleDiscountFeeText string `json:"presale_discount_fee_text"` // 预售商品-商品优惠信息
	// "presale_tail_end_time":1937297392332,
	// "presale_tail_start_time":1937297392332,
	// "presale_end_time":1937297392332,
	// "presale_start_time":1937297392332,
	// "presale_deposit":"100",
	// "ju_play_end_time":1937297392332,
	// "ju_play_start_time":1937297392332,
	PlayInfo string `json:"play_info"` // 玩法 1聚划算满减：满N件减X元，满N件X折，满N件X元） 2天猫限时抢：前N分钟每件X元，前N分钟满N件每件X元，前N件每件X元）
	// "tmall_play_activity_end_time":1937297392332,
	// "tmall_play_activity_start_time":1937297392332
}

type TLJobj struct {
	Model   TLJobjModel `json:"model"`
	MsgCode string      `json:"msg_code"`
	MsgInfo string      `json:"msg_info"`
	Success bool        `json:"success"`
}

type TLJobjModel struct {
	RightsID string `json:"rights_id"`
	SendURL  string `json:"send_url"`
}

type GetCpListJson struct {
	SearchKey  string `json:"search_key"`
	RequestID  string `json:"-"`
	ResultList struct {
		MapData []struct {
			CategoryID   int    `json:"category_id"`   //分类ＩＤ
			CategoryName string `json:"category_name"` //分类名
			// CommissionRate       string `json:"commission_rate"`
			// CommissionType       string `json:"commission_type"`
			CouponAmount  string `json:"coupon_amount,omitempty"`
			CouponEndTime string `json:"coupon_end_time,omitempty"`
			CouponID      string `json:"coupon_id"`
			CouponInfo    string `json:"coupon_info"`
			// CouponRemainCount    int    `json:"coupon_remain_count"`
			CouponShareURL   string `json:"coupon_share_url,omitempty"`
			CouponStartFee   string `json:"coupon_start_fee,omitempty"`
			CouponStartTime  string `json:"coupon_start_time,omitempty"`
			CouponTotalCount int    `json:"coupon_total_count"`
			// IncludeDxjh          string `json:"include_dxjh"`
			// IncludeMkt           string `json:"include_mkt"`
			// InfoDxjh             string `json:"info_dxjh"`
			ItemDescription string `json:"item_description"` //简介
			// ItemID               int64  `json:"item_id"` //是什么ＩＤ呢
			// ItemURL              string `json:"item_url"`
			LevelOneCategoryID   int    `json:"level_one_category_id"`
			LevelOneCategoryName string `json:"level_one_category_name"`
			Nick                 string `json:"nick"`
			NumIid               int64  `json:"num_iid"`
			PictURL              string `json:"pict_url"`
			// PresaleDeposit       string `json:"presale_deposit"`
			// PresaleEndTime       int    `json:"presale_end_time"`
			// PresaleStartTime     int    `json:"presale_start_time"`
			// PresaleTailEndTime   int    `json:"presale_tail_end_time"`
			// PresaleTailStartTime int    `json:"presale_tail_start_time"`
			// Provcity             string `json:"provcity"`
			// RealPostFee          string `json:"real_post_fee"`
			ReservePrice string `json:"reserve_price"`
			SellerID     int    `json:"seller_id"`
			ShopDsr      int    `json:"shop_dsr"` // 店铺dsr 评分
			ShopTitle    string `json:"shop_title"`
			// ShortTitle           string `json:"short_title"`
			SmallImages struct {
				String []string `json:"string"`
			} `json:"small_images"`
			Title        string `json:"title"`
			TkTotalCommi string `json:"tk_total_commi"` //tk用尽
			TkTotalSales string `json:"tk_total_sales"` //tk累计量
			URL          string `json:"url"`
			// UserType     int    `json:"user_type"`
			Volume int `json:"volume"`
			// WhiteImage   string `json:"white_image"`
			ZkFinalPrice string `json:"zk_final_price"`
		} `json:"map_data"`
	} `json:"result_list"`
	TotalResults int `json:"total_results"` // 总数
}
type GetDPXgListKeyJson struct {
	SearchKey string `json:"search_key"`
	RequestID string `json:"-"`
	Results   struct {
		NTbkShop []struct {
			PictURL    string `json:"pict_url"`
			SellerNick string `json:"seller_nick"`
			ShopTitle  string `json:"shop_title"`
			UserID     int64  `json:"user_id"`
		} `json:"n_tbk_shop"`
	} `json:"results"`
	TotalResults int `json:"total_results"`
}

type GetContentJson struct {
	Contents struct {
		MapData []struct {
			AuthorAvatar      string `json:"author_avatar"`
			AuthorID          string `json:"author_id"`
			AuthorNick        string `json:"author_nick"`
			Clink             string `json:"clink"`
			ContentCategories string `json:"content_categories"`
			ContentID         int    `json:"content_id"`
			Images            struct {
				String []string `json:"string"`
			} `json:"images"`
			ImgCount int `json:"img_count"`
			ItemIds  struct {
				Number []int64 `json:"number"`
			} `json:"item_ids"`
			PublishTime string `json:"publish_time"`
			Score       int    `json:"score"`
			Summary     string `json:"summary"`
			Tags        string `json:"tags"`
			Title       string `json:"title"`
			Type        string `json:"type"`
			UIStyle     string `json:"ui_style"`
			UpdateTime  string `json:"update_time"`
		} `json:"map_data"`
	} `json:"contents"`
	Count         int    `json:"count"`
	LastTimestamp string `json:"last_timestamp"`
}
