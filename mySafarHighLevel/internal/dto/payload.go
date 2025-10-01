// internal/dto/payload.go
package dto

type Segment struct {
    Date string `json:"date"`
    From string `json:"from"`
    To   string `json:"to"`
}

type Payload struct {
    Token        string   `json:"token"`
    IsBaggage    bool     `json:"is_baggage"`
    Lang         string   `json:"lang"`
    FilterAirlines []string `json:"filter_airlines"`
    IsDirectOnly int      `json:"is_direct_only"`
    Src          int      `json:"src"`
    Yth          int      `json:"yth"`
    Inf          int      `json:"inf"`
    Ins          int      `json:"ins"`
    Segments     []Segment `json:"segments"`
    IsCharter    bool     `json:"is_charter"`
    PriceOrder   int      `json:"price_order"`
    ArrOrder     int      `json:"arr_order"`
    DepOrder     int      `json:"dep_order"`
    DurationOrder int     `json:"duration_order"`
    Adt          int      `json:"adt"`
    Chd          int      `json:"chd"`
    GdsWhiteList []string `json:"gds_white_list"`
    GdsBlackList []string `json:"gds_black_list"`
    Count        int      `json:"count"`
    Class_       string   `json:"class_"`
}
