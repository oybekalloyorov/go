package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

// ... (Payload, Segment, RecommendationResponse same as yours) ...

// paste your Payload, Segment, RecommendationResponse structs here (same as before)

type Payload struct {
	Token          string    `json:"token"`
	IsBaggage      bool      `json:"is_baggage"`
	Lang           string    `json:"lang"`
	FilterAirlines []string  `json:"filter_airlines"`
	IsDirectOnly   int       `json:"is_direct_only"`
	Src            int       `json:"src"`
	Yth            int       `json:"yth"`
	Inf            int       `json:"inf"`
	Ins            int       `json:"ins"`
	Segments       []Segment `json:"segments"`
	IsCharter      bool      `json:"is_charter"`
	PriceOrder     int       `json:"price_order"`
	ArrOrder       int       `json:"arr_order"`
	DepOrder       int       `json:"dep_order"`
	DurationOrder  int       `json:"duration_order"`
	Adt            int       `json:"adt"`
	Chd            int       `json:"chd"`
	GdsWhiteList   []string  `json:"gds_white_list"`
	GdsBlackList   []string  `json:"gds_black_list"`
	Count          int       `json:"count"`
	Class          string    `json:"class_"`
}

type Segment struct {
	Date string `json:"date"`
	From string `json:"from"`
	To   string `json:"to"`
}

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

func main() {
	defaultURL := "https://api.mysafar.ru/v1/avia/get-recommendations"

	url := flag.String("url", defaultURL, "API endpoint")
	tokenFlag := flag.String("token", "", "API token (hash only) or set API_TOKEN env var")
	date := flag.String("date", "15.10.2025", "segment date dd.MM.yyyy")
	from := flag.String("from", "TAS", "origin code (IATA)")
	to := flag.String("to", "DXB", "destination code (IATA)")

	isBaggage := flag.Bool("is_baggage", true, "include baggage")
	adt := flag.Int("adt", 1, "adults count")
	chd := flag.Int("chd", 0, "children count")
	count := flag.Int("count", 20, "result count")
	class_ := flag.String("class", "a", "class_ value")
	lang := flag.String("lang", "ru", "lang (ru/uz/en)")
	gdsBlack := flag.String("gds_black_list", "", "comma-separated GDS codes to blacklist")
	filterAirlines := flag.String("filter_airlines", "", "comma-separated airline codes to filter out")
	isDirectOnly := flag.Int("is_direct_only", 0, "0 or 1")
	debug := flag.Bool("debug", false, "print debug request/headers")

	timeoutSec := flag.Int("timeout", 15, "request timeout seconds")
	// retries := flag.Int("retries", 3, "retry attempts (exponential backoff)")

	flag.Parse()

	token := *tokenFlag
	if token == "" {
		token = os.Getenv("API_TOKEN")
	}
	token = strings.TrimSpace(token)
	if strings.HasPrefix(strings.ToLower(token), "test ") {
		token = strings.TrimSpace(token[5:])
	}
	if token == "" {
		log.Fatal("ERROR: token is empty. Provide --token or API_TOKEN env var")
	}
	if strings.Contains(token, " ") {
		log.Fatal("ERROR: token contains spaces — pass only the token hash")
	}

	if _, err := validateDate(*date); err != nil {
		log.Fatalf("invalid date: %v", err)
	}

	if strings.EqualFold(*from, *to) {
		log.Fatalf("invalid input: from (%s) and to (%s) must be different", *from, *to)
	}

	// build slices safely
	var gdsBlackList []string
	if strings.TrimSpace(*gdsBlack) != "" {
		for _, s := range strings.Split(*gdsBlack, ",") {
			if t := strings.TrimSpace(s); t != "" {
				gdsBlackList = append(gdsBlackList, t)
			}
		}
	}
	var filterAirlinesList []string
	if strings.TrimSpace(*filterAirlines) != "" {
		for _, s := range strings.Split(*filterAirlines, ",") {
			if t := strings.TrimSpace(s); t != "" {
				filterAirlinesList = append(filterAirlinesList, t)
			}
		}
	}

	if filterAirlinesList == nil {
		filterAirlinesList = []string{}
	}
	if gdsBlackList == nil {
		gdsBlackList = []string{}
	}

	payload := Payload{
		Token:          token,
		Segments:       []Segment{{Date: *date, From: *from, To: *to}},
		IsBaggage:      *isBaggage,
		IsCharter:      false,
		PriceOrder:     1,
		ArrOrder:       1,
		DepOrder:       1,
		DurationOrder:  1,
		Adt:            *adt,
		Chd:            *chd,
		Inf:            0,
		Ins:            0,
		Src:            0,
		Yth:            0,
		Lang:           *lang,
		FilterAirlines: filterAirlinesList,
		IsDirectOnly:   *isDirectOnly,
		GdsWhiteList:   []string{},
		GdsBlackList:   gdsBlackList,
		Count:          *count,
		Class:          *class_,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("json marshal error: %v", err)
	}

	client := &http.Client{
		Timeout: time.Duration(*timeoutSec) * time.Second,
		Transport: &http.Transport{
			DialContext:         (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeoutSec+5)*time.Second)
	defer cancel()

	if *debug {
		masked := token
		if len(masked) > 8 {
			masked = masked[:4] + "..." + masked[len(masked)-4:]
		}
		fmt.Printf("DEBUG: URL = %s\n", *url)
		fmt.Printf("DEBUG: Authorization header = Token %s\n", masked)
		fmt.Printf("DEBUG: Body = %s\n", string(body))
	}

	respBody, err := sendWithDump(ctx, client, *url, body, token, *lang)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	// Print raw for easier debugging
	fmt.Println("RAW RESPONSE:", string(respBody))

	var typed RecommendationResponse
	if err := json.Unmarshal(respBody, &typed); err != nil {
		fmt.Fprintf(os.Stderr, "response unmarshal error: %v\n", err)
		return
	}

	// If recommendations nil -> print user-friendly message
	if typed.Recommendations == nil || len(typed.Recommendations) == 0 {
		fmt.Println("No flights found (recommendations is null or empty). Try changing date/class/filters.")
		return
	}

	pretty, _ := json.MarshalIndent(typed, "", "  ")
	fmt.Println(string(pretty))
}

func validateDate(d string) (time.Time, error) {
	if d == "" {
		return time.Time{}, errors.New("empty date")
	}
	t, err := time.Parse("02.01.2006", d)
	if err != nil {
		return time.Time{}, fmt.Errorf("wrong format, want dd.MM.yyyy: %w", err)
	}
	return t, nil
}

// sendWithDump - sets more headers (Accept, Accept-Language, User-Agent), dumps request & response
func sendWithDump(ctx context.Context, client *http.Client, url string, body []byte, token string, lang string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Important headers — match Postman as close as possible
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	if lang != "" {
		req.Header.Set("Accept-Language", lang)
	}
	// mimic Postman User-Agent (optional)
	req.Header.Set("User-Agent", "PostmanRuntime/7.29.0")
	if token != "" {
		req.Header.Set("Authorization", "Token "+token)
	}

	// Dump request (safe - DumpRequestOut rewinds body)
	reqDump, _ := httputil.DumpRequestOut(req, true)
	// Mask token in dump for printing
	out := string(reqDump)
	out = maskTokenInDump(out)
	fmt.Println("----- REQUEST DUMP -----")
	fmt.Println(out)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respDump, _ := httputil.DumpResponse(resp, true)
	fmt.Println("----- RESPONSE DUMP -----")
	fmt.Println(string(respDump))

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respBody, fmt.Errorf("non-2xx status: %d", resp.StatusCode)
	}
	return respBody, nil
}

func maskTokenInDump(s string) string {
	// best-effort mask: replace "Token <hash>" with "Token xxxx...xxxx"
	parts := strings.Split(s, "Authorization: ")
	if len(parts) < 2 {
		return s
	}
	rest := parts[1]
	idx := strings.Index(rest, "\r\n")
	if idx < 0 {
		return s
	}
	tokenLine := rest[:idx]
	if strings.HasPrefix(tokenLine, "Token ") {
		tok := strings.TrimSpace(strings.TrimPrefix(tokenLine, "Token "))
		if len(tok) > 8 {
			masked := tok[:4] + "..." + tok[len(tok)-4:]
			return strings.Replace(s, tokenLine, "Token "+masked, 1)
		}
	}
	return s
}
