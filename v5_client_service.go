package bybit

// V5ServiceI :
type V5ServiceI interface {
	Market() V5MarketServiceI
	Order() V5OrderServiceI
	Position() V5PositionServiceI
	Execution() V5ExecutionServiceI
	Account() V5AccountServiceI
	SpotLeverageToken() V5SpotLeverageTokenServiceI
	SpotMarginTrade() V5SpotMarginTradeServiceI
	Asset() V5AssetServiceI
	User() V5UserServiceI
	Announcements() V5AnnouncementsServiceI
}

// V5Service :
type V5Service struct {
	client *Client
}

func (s *V5Service) Announcements() V5AnnouncementsServiceI {
	return &V5AnnouncementsService{client: s.client}
}

// Market :
func (s *V5Service) Market() V5MarketServiceI {
	return &V5MarketService{client: s.client}
}

// Order :
func (s *V5Service) Order() V5OrderServiceI {
	return &V5OrderService{client: s.client}
}

// Position :
func (s *V5Service) Position() V5PositionServiceI {
	return &V5PositionService{client: s.client}
}

// Execution :
func (s *V5Service) Execution() V5ExecutionServiceI {
	return &V5ExecutionService{client: s.client}
}

// Account :
func (s *V5Service) Account() V5AccountServiceI {
	return &V5AccountService{client: s.client}
}

// SpotLeverageToken :
func (s *V5Service) SpotLeverageToken() V5SpotLeverageTokenServiceI {
	return &V5SpotLeverageTokenService{client: s.client}
}

// SpotMarginTrade :
func (s *V5Service) SpotMarginTrade() V5SpotMarginTradeServiceI {
	return &V5SpotMarginTradeService{client: s.client}
}

// Asset :
func (s *V5Service) Asset() V5AssetServiceI {
	return &V5AssetService{client: s.client}
}

// User :
func (s *V5Service) User() V5UserServiceI {
	return &V5UserService{client: s.client}
}

// V5 :
func (c *Client) V5() V5ServiceI {
	return &V5Service{c.withCheckResponseBody(checkV5ResponseBody)}
}
