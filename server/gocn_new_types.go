package server

type TopicListResp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data TopicListData `json:"data"`
}

type TopicList struct {
	GUID       string `json:"guid"`
	UID        int    `json:"uid"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Ctime      int    `json:"ctime"`
	CntView    int    `json:"cntView"`
	Cate2ID    int    `json:"cate2Id"`
	Cate2Title string `json:"cate2Title"`
	CntLike    int    `json:"cntLike,omitempty"`
	CntCollect int    `json:"cntCollect,omitempty"`
	CntReply   int    `json:"cntReply,omitempty"`
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

type Tdk struct {
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
}

type Topic struct {
	GUID        string `json:"guid"`
	UID         int    `json:"uid"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Title       string `json:"title"`
	ContentHTML string `json:"contentHtml"`
	Ctime       int    `json:"ctime"`
	CntView     int    `json:"cntView"`
	Cate2ID     int    `json:"cate2Id"`
	Cate2Title  string `json:"cate2Title"`
}

type TopicInfoData struct {
	Tdk   Tdk   `json:"tdk"`
	Topic Topic `json:"topic"`
}
