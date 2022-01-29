package server

type TopicListResp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data TopicListData `json:"data"`
}

type TopicList struct {
	ID         int    `json:"id"` // topic id 重点
	UID        int    `json:"uid,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	Avatar     string `json:"avatar"`
	Title      string `json:"title"` // 重点
	Summary    string `json:"summary"`
	Ctime      int    `json:"ctime"`
	CntView    int    `json:"cntView"`
	Cate2ID    int    `json:"cate2Id"`
	Cate2Title string `json:"cate2Title"`
	CntLike    int    `json:"cntLike,omitempty"`
	CntCollect int    `json:"cntCollect,omitempty"`
	CntReply   int    `json:"cntReply,omitempty"`
	IsOldData  bool   `json:"isOldData,omitempty"`
}

type Pagination struct {
	Total       int    `json:"total"`
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"pageSize"`
	Sort        string `json:"sort"`
}

type TopicListData struct {
	List       []TopicList `json:"list"`
	Pagination Pagination  `json:"pagination"`
}

type TopicInfoResp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data TopicInfoData `json:"data"`
}

type TopicInfoData struct {
	ID          int    `json:"id"`
	UID         int    `json:"uid"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Title       string `json:"title"`
	Content     string `json:"content"` // 取出 content 即可
	ContentHTML string `json:"contentHtml"`
	Ctime       int    `json:"ctime"`
	CntView     int    `json:"cntView"`
	Cate2ID     int    `json:"cate2Id"`
	Cate2Title  string `json:"cate2Title"`
}
