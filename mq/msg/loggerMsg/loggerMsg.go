package loggerMsg

import "reflect"

type UserRegisterRecord struct {
	UserId  int64  `json:"user_id"`
	Date    string `json:"date"`
	LoginIp string `json:"login_ip"`
}

type UserLoginRecord struct {
	UserId      int64  `json:"user_id"`
	LoginStatus int    `json:"login_status"`
	Date        string `json:"date"`
}

type Bjpk10Record struct {
	CategoryId     int    `json:"category_id"`
	GameId         int    `json:"game_id"`
	Round          int64  `json:"round"`
	Result         []int  `json:"result"`
	DrawResultDate string `json:"draw_result_date"`
	Pattern        Bjpk10 `json:"pattern"`
	DrawStatus     int    `json:"draw_status"`
}

type LsRecord struct {
	CategoryId     int    `json:"category_id"`
	GameId         int    `json:"game_id"`
	Round          int64  `json:"round"`
	Result         []int  `json:"result"`
	DrawResultDate string `json:"draw_result_date"`
	Pattern        Ls     `json:"pattern"`
	DrawStatus     int    `json:"draw_status"`
}

type Spk10Record struct {
	CategoryId     int    `json:"category_id"`
	GameId         int    `json:"game_id"`
	Round          int64  `json:"round"`
	Result         []int  `json:"result"`
	DrawResultDate string `json:"draw_result_date"`
	Pattern        Spk10  `json:"pattern"`
	DrawStatus     int    `json:"draw_status"`
}

type CqsscRecord struct {
	CategoryId     int    `json:"category_id"`
	GameId         int    `json:"game_id"`
	Round          int64  `json:"round"`
	Result         []int  `json:"result"`
	DrawResultDate string `json:"draw_result_date"`
	Pattern        Cqssc  `json:"pattern"`
	DrawStatus     int    `json:"draw_status"`
}

type Gdkl10Record struct {
	CategoryId     int    `json:"category_id"`
	GameId         int    `json:"game_id"`
	Round          int64  `json:"round"`
	Result         []int  `json:"result"`
	DrawResultDate string `json:"draw_result_date"`
	Pattern        Gdkl10 `json:"pattern"`
	DrawStatus     int    `json:"draw_status"`
}

type Jsk3Record struct {
	CategoryId     int    `json:"category_id"`
	GameId         int    `json:"game_id"`
	Round          int64  `json:"round"`
	Result         []int  `json:"result"`
	DrawResultDate string `json:"draw_result_date"`
	Pattern        Jsk3   `json:"pattern"`
	DrawStatus     int    `json:"draw_status"`
}

type Bjpk10 struct {
	First          map[string]int `json:"first"`
	Second         map[string]int `json:"second"`
	Third          map[string]int `json:"third"`
	Fourth         map[string]int `json:"fourth"`
	Fifth          map[string]int `json:"fifth"`
	Sixth          map[string]int `json:"sixth"`
	Seventh        map[string]int `json:"seventh"`
	Eighth         map[string]int `json:"eighth"`
	Ninth          map[string]int `json:"ninth"`
	Tenth          map[string]int `json:"tenth"`
	DragonTiger    map[string]int `json:"dragon_tiger"`
	FirstSecondSum map[string]int `json:"first_second_sum"`
	FirstSecondCom string         `json:"first_second_com"`
}

type Ls struct {
	First          map[string]int `json:"first"`
	Second         map[string]int `json:"second"`
	Third          map[string]int `json:"third"`
	Fourth         map[string]int `json:"fourth"`
	Fifth          map[string]int `json:"fifth"`
	Sixth          map[string]int `json:"sixth"`
	Seventh        map[string]int `json:"seventh"`
	Eighth         map[string]int `json:"eighth"`
	Ninth          map[string]int `json:"ninth"`
	Tenth          map[string]int `json:"tenth"`
	DragonTiger    map[string]int `json:"dragon_tiger"`
	FirstSecondSum map[string]int `json:"first_second_sum"`
	FirstSecondCom string         `json:"first_second_com"`
}

type Spk10 struct {
	First          map[string]int `json:"first"`
	Second         map[string]int `json:"second"`
	Third          map[string]int `json:"third"`
	Fourth         map[string]int `json:"fourth"`
	Fifth          map[string]int `json:"fifth"`
	Sixth          map[string]int `json:"sixth"`
	Seventh        map[string]int `json:"seventh"`
	Eighth         map[string]int `json:"eighth"`
	Ninth          map[string]int `json:"ninth"`
	Tenth          map[string]int `json:"tenth"`
	DragonTiger    map[string]int `json:"dragon_tiger"`
	FirstSecondSum map[string]int `json:"first_second_sum"`
	FirstSecondCom string         `json:"first_second_com"`
}

type Cqssc struct {
	BallOne         map[string]int `json:"ball_1"`
	BallTwo         map[string]int `json:"ball_2"`
	BallThree       map[string]int `json:"ball_3"`
	BallFour        map[string]int `json:"ball_4"`
	BallFive        map[string]int `json:"ball_5"`
	Sum             map[string]int `json:"sum"`
	DragonTiger     map[string]int `json:"dragon_tiger"`
	FrontMiddleBack []int          `json:"front_middle_back"`
}

type Gdkl10 struct {
	BallOne     map[string]int `json:"ball_1"`
	BallTwo     map[string]int `json:"ball_2"`
	BallThree   map[string]int `json:"ball_3"`
	BallFour    map[string]int `json:"ball_4"`
	BallFive    map[string]int `json:"ball_5"`
	BallSix     map[string]int `json:"ball_6"`
	BallSeven   map[string]int `json:"ball_7"`
	BallEight   map[string]int `json:"ball_8"`
	Sum         map[string]int `json:"sum"`
	DragonTiger map[string]int `json:"dragon_tiger"`
}

type Jsk3 struct {
	BigSmall     int   `json:"big_small"`
	SingleDouble int   `json:"single_double"`
	Sum          int   `json:"sum"`
	SameTwo      int   `json:"same_two"`
	LinkTwo      []int `json:"link_two"`
	SameThree    int   `json:"same_three"`
	LinkThree    int   `json:"link_three"`
}

type PlaceOrderRecord struct {
	OrderId      int64   `json:"order_id"`
	UserId       int64   `json:"user_id"`
	GameId       int     `json:"game_id"`
	Round        int64   `json:"round"`
	Target       string  `json:"target"`
	Value        string  `json:"value"`
	Odds         float32 `json:"odds"`
	Limit        float32 `json:"limit"`
	Refund       float32 `json:"refund"`
	OrderCredit  float32 `json:"order_credit"`
	EffectCredit float32 `json:"effect_credit"`
	PayoutResult bool    `json:"payout_result"`
	PayoutCredit float32 `json:"payout_credit"`
	Date         string  `json:"date"`
}

type DrawResultOrderRecord struct {
	OrderId      int64   `json:"order_id"`
	PayoutResult bool    `json:"payout_result"`
	PayoutCredit float32 `json:"payout_credit"`
}

type UpdateWalletRecord struct {
	TransferId         int64   `json:"transfer_id"`
	UserId             int64   `json:"user_id"`
	Account            string  `json:"account"`
	AgentId            int64     `json:"agent_id"`
	AgentAccount       string  `json:"agent_account"`
	MasterAgentId      int64     `json:"master_agent_id"`
	MasterAgentAccount string  `json:"master_agent_account"`
	Reason             int     `json:"reason"`
	ReasonMessage      string  `json:"reason_message"`
	ChangeCredit       float32 `json:"change_credit"`
	BalanceCredit      float32 `json:"balance_credit"`
	GameId             int     `json:"game_id"`
	OrderId            int64   `json:"order_id"`
	AgentTransferId    string  `json:"agent_transfer_id"`
	Date               string  `json:"date"`
}

type QueryLotteryRequest struct {
	GameId    int    `json:"game_id"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`

	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Sort    string `json:"sort"`
	SortDir int    `json:"dir"`
}

func (req *QueryLotteryRequest) NotEmpty() bool {

	if req.GameId != reflect.Zero(reflect.TypeOf(req.GameId)).Interface() &&
		req.BeginDate != reflect.Zero(reflect.TypeOf(req.BeginDate)).Interface() &&
		req.EndDate != reflect.Zero(reflect.TypeOf(req.EndDate)).Interface() {
		return true
	} else {
		return false
	}
}

type QueryWalletChange struct {
	Scope     int    `json:"scope"`
	QuerierId int64    `json:"querier_id"`
	UserId    int64  `json:"user_id"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`

	Limit                 int    `json:"limit"`
	Page                  int    `json:"page"`
	Sort                  string `json:"sort"`
	SortDir               int    `json:"dir"`
	FilterAgentAccount    string `json:"filter_agent_account"`
	FilterAccount         string `json:"filter_account"`
	FilterReason          int    `json:"filter_reason"`
	FilterOrderId         int64  `json:"filter_order_id"`
	FilterAgentTransferId string `json:"filter_agent_transfer_id"`
}

func (logMsg *QueryWalletChange) NotEmpty() bool {

	if logMsg.QuerierId != reflect.Zero(reflect.TypeOf(logMsg.QuerierId)).Interface() &&
		logMsg.UserId != reflect.Zero(reflect.TypeOf(logMsg.UserId)).Interface() &&
		logMsg.BeginDate != reflect.Zero(reflect.TypeOf(logMsg.BeginDate)).Interface() &&
		logMsg.EndDate != reflect.Zero(reflect.TypeOf(logMsg.EndDate)).Interface() {
		return true
	}
	return false
}
func (logMsg *QueryWalletChange) NotEmpty_transfer() bool {

	if logMsg.QuerierId != reflect.Zero(reflect.TypeOf(logMsg.QuerierId)).Interface() {

		return true
	}
	return false
}

type QueryGameStatistics struct {
	Scope     int    `json:"scope"`
	Type      int    `json:"type"`
	QuerierId int64    `json:"querier_id"`
	GameId    int    `json:"game_id"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`

	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	FilterAgentId int64    `json:"filter_agent_id"`
	Sort          string `json:"sort"`
	SortDir       int    `json:"dir"`
}

func (logMsg *QueryGameStatistics) NotEmpty() bool {

	if logMsg.GameId != reflect.Zero(reflect.TypeOf(logMsg.GameId)).Interface() &&
		logMsg.BeginDate != reflect.Zero(reflect.TypeOf(logMsg.BeginDate)).Interface() &&
		logMsg.EndDate != reflect.Zero(reflect.TypeOf(logMsg.EndDate)).Interface() {
		return true
	}
	return false
}

type GameStatisticsResponseMessage struct {
	Id                 int64   `json:"id"`
	AgentId            int64     `json:"agent_id"`
	AgentAccount       string  `json:"agent_account"`
	MasterAgentId      int64     `json:"master_agent_id"`
	MasterAgentAccount string  `json:"master_agent_account"`
	GameId             int     `json:"game_id"`
	PlaceOrder         int     `json:"place_order"`
	OrderCredit        float32 `json:"order_credit"`
	EffectCredit       float32 `json:"effect_credit"`
	AverageOrderCredit float32 `json:"average_credit"`
	Payout             int     `json:"payout"`
	PayoutCredit       float32 `json:"payout_credit"`
	Refund             float32 `json:"refund"`
	ProfitAndLoss      float32 `json:"profit_and_loss"`
	Date               string  `json:"date"`
}

type QueryUserStatistics struct {
	Scope     int    `json:"scope"`
	QuerierId int64    `json:"querier_id"`
	Type      int    `json:"type"`
	UserId    int64  `json:"user_id"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`

	Limit              int    `json:"limit"`
	Page               int    `json:"page"`
	FilterAgentAccount string `json:"filter_agent_account"`
	Sort               string `json:"sort"`
	SortDir            int    `json:"dir"`
}

func (logMsg QueryUserStatistics) NotEmpty() bool {

	if logMsg.UserId != reflect.Zero(reflect.TypeOf(logMsg.UserId)).Interface() &&
		logMsg.BeginDate != reflect.Zero(reflect.TypeOf(logMsg.BeginDate)).Interface() &&
		logMsg.EndDate != reflect.Zero(reflect.TypeOf(logMsg.EndDate)).Interface() {
		return true
	}
	return false
}

type UserStatisticsResponseMessage struct {
	Id                 int64   `json:"id"`
	AgentId            int64     `json:"agent_id"`
	AgentAccount       string  `json:"agent_account"`
	MasterAgentId      int64     `json:"master_agent_id"`
	MasterAgentAccount string  `json:"master_agent_account"`
	UserId             int64   `json:"user_id"`
	Account            string  `json:"account"`
	Name               string  `json:"name"`
	PlaceOrder         int     `json:"place_order"`
	OrderCredit        float32 `json:"order_credit"`
	EffectCredit       float32 `json:"effect_credit"`
	AverageOrderCredit float32 `json:"average_credit"`
	Payout             int     `json:"payout"`
	PayoutCredit       float32 `json:"payout_credit"`
	Refund             float32 `json:"refund"`
	ProfitAndLoss      float32 `json:"profit_and_loss"`
	Date               string  `json:"date"`
}

type QueryAgentStatistics struct {
	Scope     int    `json:"scope"`
	QuerierId int64    `json:"querier_id"`
	Type      int    `json:"type"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`

	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	FilterAgentId int64    `json:"filter_agent_id"`
	Sort          string `json:"sort"`
	SortDir       int    `json:"dir"`
}

func (logMsg QueryAgentStatistics) NotEmpty() bool {
	if logMsg.QuerierId != reflect.Zero(reflect.TypeOf(logMsg.QuerierId)).Interface() && logMsg.BeginDate != reflect.Zero(reflect.TypeOf(logMsg.BeginDate)).Interface() &&
		logMsg.EndDate != reflect.Zero(reflect.TypeOf(logMsg.EndDate)).Interface() {
		return true
	}
	return false
}

type AgentStatisticsResponseMessage struct {
	Id                 int64   `json:"id"`
	AgentId            int64     `json:"agent_id"`
	AgentAccount       string  `json:"agent_account"`
	MasterAgentId      int64     `json:"master_agent_id"`
	MasterAgentAccount string  `json:"master_agent_account"`
	TransferInCredit   float32 `json:"transfer_in_credit"`
	TransferOutCredit  float32 `json:"transfer_out_credit"`
	RegisterNumbers    int     `json:"register_numbers"`
	PlaceOrder         int     `json:"place_order"`
	OrderCredit        float32 `json:"order_credit"`
	EffectCredit       float32 `json:"effect_credit"`
	AverageOrderCredit float32 `json:"average_credit"`
	Payout             int     `json:"payout"`
	PayoutCredit       float32 `json:"payout_credit"`
	Refund             float32 `json:"refund"`
	ProfitAndLoss      float32 `json:"profit_and_loss"`
	Date               string  `json:"date"`
}

type QueryTransferStatistics struct {
	Scope     int    `json:"scope"`
	QuerierId int64    `json:"querier_id"`
	Type      int    `json:"type"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`

	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (logMsg QueryTransferStatistics) NotEmpty() bool {

	if logMsg.QuerierId != reflect.Zero(reflect.TypeOf(logMsg.QuerierId)).Interface() &&
		logMsg.BeginDate != reflect.Zero(reflect.TypeOf(logMsg.BeginDate)).Interface() &&
		logMsg.EndDate != reflect.Zero(reflect.TypeOf(logMsg.EndDate)).Interface() {
		return true
	}
	return false
}

type TransferStatisticsResponseMessage struct {
	Id                 int64   `json:"id"`
	AgentId            int64     `json:"agent_id"`
	AgentAccount       string  `json:"agent_account"`
	MasterAgentId      int64     `json:"master_agent_id"`
	MasterAgentAccount string  `json:"master_agent_account"`
	TransferIn         int     `json:"transfer_in"`
	TransferInCredit   float32 `json:"transfer_in_credit"`
	TransferOut        int     `json:"transfer_out"`
	TransferOutCredit  float32 `json:"transfer_out_credit"`
	Date               string  `json:"date"`
}

type QueryOperationStatistics struct {
	Scope     int    `json:"scope"`
	QuerierId int64    `json:"querier_id"`
	Type      int    `json:"type"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`

	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (logMsg QueryOperationStatistics) NotEmpty() bool {

	if logMsg.QuerierId != reflect.Zero(reflect.TypeOf(logMsg.QuerierId)).Interface() &&
		logMsg.BeginDate != reflect.Zero(reflect.TypeOf(logMsg.BeginDate)).Interface() &&
		logMsg.EndDate != reflect.Zero(reflect.TypeOf(logMsg.EndDate)).Interface() {
		return true
	}
	return false
}

type QueryAuthRecord struct {
	Scope     int    `json:"scope"`
	QuerierId int64    `json:"querier_id"`
	UserId    int64  `json:"user_id"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`

	Limit         int    `json:"limit"`
	Page          int    `json:"page"`
	Sort          string `json:"sort"`
	SortDir       int    `json:"dir"`
	FilterAgentId int64    `json:"filter_agent_id"`
}

func (logMsg QueryAuthRecord) NotEmpty() bool {
	if logMsg.QuerierId != reflect.Zero(reflect.TypeOf(logMsg.QuerierId)).Interface() &&
		logMsg.BeginDate != reflect.Zero(reflect.TypeOf(logMsg.BeginDate)).Interface() &&
		logMsg.EndDate != reflect.Zero(reflect.TypeOf(logMsg.EndDate)).Interface() {
		return true
	}
	return false
}

type QueryOrderRecord struct {
	Scope     int    `json:"scope"`
	QuerierId int64   `json:"querier_id"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`

	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Sort    string `json:"sort"`
	SortDir int    `json:"dir"`

	FilterGameId  int   `json:"filter_game_id"`
	FilterRound   int   `json:"filter_round"`
	FilterAgentId int64   `json:"filter_agent_id"`
	FilterUserId  int64 `json:"filter_user_id"`
}

func (logMsg QueryOrderRecord) NotEmpty() bool {
	if //logMsg.QuerierId != reflect.Zero(reflect.TypeOf(logMsg.QuerierId)).Interface() &&
	logMsg.BeginDate != reflect.Zero(reflect.TypeOf(logMsg.BeginDate)).Interface() &&
		logMsg.EndDate != reflect.Zero(reflect.TypeOf(logMsg.EndDate)).Interface() { //&&
		//logMsg.Scope != reflect.Zero(reflect.TypeOf(logMsg.Scope)).Interface()

		return true
	}
	return false
}

type OperationStatisticsResponseMessage struct {
	Id                 int64  `json:"id"`
	AgentId            int64    `json:"agent_id"`
	AgentAccount       string `json:"agent_account"`
	MasterAgentId      int64    `json:"master_agent_id"`
	MasterAgentAccount string `json:"master_agent_account"`
	LoginNumbers       int    `json:"login_numbers"`
	RegisterNumbers    int    `json:"register_numbers"`
	Date               string `json:"date"`
}

type UserLoginResponseMessage struct {
	AgentId            int64    `json:"agent_id"`
	AgentAccount       string `json:"agent_account"`
	MasterAgentId      int64    `json:"master_agent_id"`
	MasterAgentAccount string `json:"master_agent_account"`
	Counts             int    `json:"counts"`
	Date               string `json:"date"`
}

type QueryLotteryStatistics struct {
	GameId int `json:"game_id"`
}

type QueryLongDragonStatistics struct {
	GameId int `json:"game_id"`
}

type Bjpk10StatisticsResponseMessage struct {
	First   []int `json:"first"`
	Second  []int `json:"second"`
	Third   []int `json:"third"`
	Fourth  []int `json:"fourth"`
	Fifth   []int `json:"fifth"`
	Sixth   []int `json:"sixth"`
	Seventh []int `json:"seventh"`
	Eighth  []int `json:"eighth"`
	Ninth   []int `json:"ninth"`
	Tenth   []int `json:"tenth"`
}

type Bjpk10LongDragonResponseMessage struct {
	First          map[string]int `json:"first"`
	Second         map[string]int `json:"second"`
	Third          map[string]int `json:"third"`
	Fourth         map[string]int `json:"fourth"`
	Fifth          map[string]int `json:"fifth"`
	Sixth          map[string]int `json:"sixth"`
	Seventh        map[string]int `json:"seventh"`
	Eighth         map[string]int `json:"eighth"`
	Ninth          map[string]int `json:"ninth"`
	Tenth          map[string]int `json:"tenth"`
	FirstSecondSum map[string]int `json:"first_second_sum"`
}

type LsStatisticsResponseMessage struct {
	First   []int `json:"first"`
	Second  []int `json:"second"`
	Third   []int `json:"third"`
	Fourth  []int `json:"fourth"`
	Fifth   []int `json:"fifth"`
	Sixth   []int `json:"sixth"`
	Seventh []int `json:"seventh"`
	Eighth  []int `json:"eighth"`
	Ninth   []int `json:"ninth"`
	Tenth   []int `json:"tenth"`
}

type LsLongDragonResponseMessage struct {
	First          map[string]int `json:"first"`
	Second         map[string]int `json:"second"`
	Third          map[string]int `json:"third"`
	Fourth         map[string]int `json:"fourth"`
	Fifth          map[string]int `json:"fifth"`
	Sixth          map[string]int `json:"sixth"`
	Seventh        map[string]int `json:"seventh"`
	Eighth         map[string]int `json:"eighth"`
	Ninth          map[string]int `json:"ninth"`
	Tenth          map[string]int `json:"tenth"`
	FirstSecondSum map[string]int `json:"first_second_sum"`
}

type Spk10StatisticsResponseMessage struct {
	First   []int `json:"first"`
	Second  []int `json:"second"`
	Third   []int `json:"third"`
	Fourth  []int `json:"fourth"`
	Fifth   []int `json:"fifth"`
	Sixth   []int `json:"sixth"`
	Seventh []int `json:"seventh"`
	Eighth  []int `json:"eighth"`
	Ninth   []int `json:"ninth"`
	Tenth   []int `json:"tenth"`
}

type Spk10LongDragonResponseMessage struct {
	First          map[string]int `json:"first"`
	Second         map[string]int `json:"second"`
	Third          map[string]int `json:"third"`
	Fourth         map[string]int `json:"fourth"`
	Fifth          map[string]int `json:"fifth"`
	Sixth          map[string]int `json:"sixth"`
	Seventh        map[string]int `json:"seventh"`
	Eighth         map[string]int `json:"eighth"`
	Ninth          map[string]int `json:"ninth"`
	Tenth          map[string]int `json:"tenth"`
	FirstSecondSum map[string]int `json:"first_second_sum"`
}

type CqsscStatisticsResponseMessage struct {
	BallOne   []int `json:"ball_1"`
	BallTwo   []int `json:"ball_2"`
	BallThree []int `json:"ball_3"`
	BallFour  []int `json:"ball_4"`
	BallFive  []int `json:"ball_5"`
}

type CqsscLongDragonResponseMessage struct {
	BallOne     map[string]int `json:"ball_1"`
	BallTwo     map[string]int `json:"ball_2"`
	BallThree   map[string]int `json:"ball_3"`
	BallFour    map[string]int `json:"ball_4"`
	BallFive    map[string]int `json:"ball_5"`
	Sum         map[string]int `json:"sum"`
	DragonTiger map[string]int `json:"dragon_tiger"`
}

type Gdkl10StatisticsResponseMessage struct {
	BallOne   []int `json:"ball_1"`
	BallTwo   []int `json:"ball_2"`
	BallThree []int `json:"ball_3"`
	BallFour  []int `json:"ball_4"`
	BallFive  []int `json:"ball_5"`
	BallSix   []int `json:"ball_6"`
	BallSeven []int `json:"ball_7"`
	BallEight []int `json:"ball_8"`
}

type Gdkl10LongDragonResponseMessage struct {
	BallOne   map[string]int `json:"ball_1"`
	BallTwo   map[string]int `json:"ball_2"`
	BallThree map[string]int `json:"ball_3"`
	BallFour  map[string]int `json:"ball_4"`
	BallFive  map[string]int `json:"ball_5"`
	BallSix   map[string]int `json:"ball_6"`
	BallSeven map[string]int `json:"ball_7"`
	BallEight map[string]int `json:"ball_8"`
	Sum       map[string]int `json:"sum"`
}

type QueryLotteryPlaceOrderSum struct {
	AgentId int64   `json:"agent_id"`
	GameId  int   `json:"game_id"`
	Round   int64 `json:"round"`

	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Bjpk10PlaceOrderSumResponseMessage struct {
	AgentId int64                 `json:"agent_id"`
	GameId  int                 `json:"game_id"`
	Round   int64               `json:"round"`
	Sum     Bjpk10PlaceOrderSum `json:"sum"`
}

type Bjpk10PlaceOrderSum struct {
	Total             int   `json:"total"`
	Position          []int `json:"position"`
	PositionBS        []int `json:"position_bs"`
	PositionSD        []int `json:"position_sd"`
	FirstSecondSumBS  int   `json:"first_second_sum_bs"`
	FirstSecondSumSD  int   `json:"first_second_sum_sd"`
	FirstSecondSumSum []int `json:"first_second_sum_sum"`
	FirstSecondCom    []int `json:"first_second_com"`
	DragonTiger       []int `json:"dragon_tiger"`
}

type LsPlaceOrderSumResponseMessage struct {
	AgentId int64             `json:"agent_id"`
	GameId  int             `json:"game_id"`
	Round   int64           `json:"round"`
	Sum     LsPlaceOrderSum `json:"sum"`
}

type LsPlaceOrderSum struct {
	Total             int   `json:"total"`
	Position          []int `json:"position"`
	PositionBS        []int `json:"position_bs"`
	PositionSD        []int `json:"position_sd"`
	FirstSecondSumBS  int   `json:"first_second_sum_bs"`
	FirstSecondSumSD  int   `json:"first_second_sum_sd"`
	FirstSecondSumSum []int `json:"first_second_sum_sum"`
	FirstSecondCom    []int `json:"first_second_com"`
	DragonTiger       []int `json:"dragon_tiger"`
}

type Spk10PlaceOrderSumResponseMessage struct {
	AgentId int64                `json:"agent_id"`
	GameId  int                `json:"game_id"`
	Round   int64              `json:"round"`
	Sum     Spk10PlaceOrderSum `json:"sum"`
}

type Spk10PlaceOrderSum struct {
	Total             int   `json:"total"`
	Position          []int `json:"position"`
	PositionBS        []int `json:"position_bs"`
	PositionSD        []int `json:"position_sd"`
	FirstSecondSumBS  int   `json:"first_second_sum_bs"`
	FirstSecondSumSD  int   `json:"first_second_sum_sd"`
	FirstSecondSumSum []int `json:"first_second_sum_sum"`
	FirstSecondCom    []int `json:"first_second_com"`
	DragonTiger       []int `json:"dragon_tiger"`
}

type CqsscPlaceOrderSumResponseMessage struct {
	AgentId int64                `json:"agent_id"`
	GameId  int                `json:"game_id"`
	Round   int64              `json:"round"`
	Sum     CqsscPlaceOrderSum `json:"sum"`
}

type CqsscPlaceOrderSum struct {
	Total       int   `json:"total"`
	Ball15      []int `json:"ball_1_5"`
	Ball15Bs    []int `json:"ball_1_5_bs"`
	Ball15Sd    []int `json:"ball_1_5_sd"`
	SumBs       int   `json:"sum_bs"`
	SumSd       int   `json:"sum_sd"`
	DragonTiger []int `json:"dragon_tiger"`
	FrontThree  []int `json:"front_three"`
	MiddleThree []int `json:"middle_three"`
	LastThree   []int `json:"last_three"`
}

type Gdkl10PlaceOrderSumResponseMessage struct {
	AgentId int64                 `json:"agent_id"`
	GameId  int                 `json:"game_id"`
	Round   int64               `json:"round"`
	Sum     Gdkl10PlaceOrderSum `json:"sum"`
}

type Gdkl10PlaceOrderSum struct {
	Total        int   `json:"total"`
	Ball18       []int `json:"ball_1_8"`
	Ball18Bs     []int `json:"ball_1_8_bs"`
	Ball18Sd     []int `json:"ball_1_8_sd"`
	Ball18TailBs []int `json:"ball_1_8_tail_bs"`
	Ball18SumSd  []int `json:"ball_1_8_sum_sd"`
	SumBs        int   `json:"sum_bs"`
	SumSd        int   `json:"sum_sd"`
	SumTailBs    int   `json:"sum_tail_bs"`
	DragonTiger  []int `json:"dragon_tiger"`
	Ball18Dnsb   []int `json:"ball_1_8_dnsb"`
	Ball18Zfb    []int `json:"ball_1_8_zfb"`
	FormalNumber int   `json:"formal_number"`
	AnyTwo       int   `json:"any_two"`
	SerialTwo    int   `json:"serial_two"`
	AnyThree     int   `json:"any_three"`
	SerialThree  int   `json:"serial_three"`
	AnyFour      int   `json:"any_four"`
	AnyFive      int   `json:"any_five"`
}

type Jsk3PlaceOrderSumResponseMessage struct {
	AgentId int64               `json:"agent_id"`
	GameId  int               `json:"game_id"`
	Round   int64             `json:"round"`
	Sum     Jsk3PlaceOrderSum `json:"sum"`
}

type Jsk3PlaceOrderSum struct {
	Total      int `json:"total"`
	SingleDice int `json:"single_dice"`
	SumBs      int `json:"sum_sd"`
	SumSd      int `json:"sum_bs"`
	Sum        int `json:"sum"`
	SameTwo    int `json:"same_two"`
	LinkTwo    int `json:"link_two"`
	SameThree  int `json:"same_three"`
	LinkThree  int `json:"link_three"`
}

type WithdrawResultRecord struct {
	GameId int   `json:"game_id"`
	Round  int64 `json:"round"`
}

type MissRoundRecord struct {
	CategoryId int   `json:"category_id"`
	GameId     int   `json:"game_id"`
	Round      int64 `json:"round"`
}

type QueryMissRoundRecord struct {
	GameId int `json:"game_id"`

	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type MissRoundResponseMessage struct {
	GameId int   `json:"game_id"`
	Round  int64 `json:"round"`
}

