package dto

type RecommendationResponse struct {
    RequestID       string `json:"request_id"`
    Recommendations []struct {
        SegmentIndex int `json:"segment_index"`
        Flights      []struct {
            Airline   string  `json:"airline"`
            Price     float64 `json:"price"`
            Currency  string  `json:"currency"`
            Departure string  `json:"departure_time"`
            Arrival   string  `json:"arrival_time"`
        } `json:"flights"`
    } `json:"recommendations"`
}
