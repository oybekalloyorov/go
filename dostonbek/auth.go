package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Args defines the parameters for RPC calls.
type Args struct {
	A, B int
}

// Arith defines a simple RPC type.
type Arith int

// Multiply multiplies two integers and returns the result.
func (t *Arith) Multiply(args *Args, reply *int) error {
	if args.A < 0 || args.B < 0 {
		return fmt.Errorf("inputs must be non-negative")
	}
	*reply = args.A * args.B
	return nil
}

var jwtSecret = []byte("supersecretkey")

func validateJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token: %v", err)
	}
	return nil
}

func main() {
	// Register RPC
	arith := new(Arith)
	if err := rpc.Register(arith); err != nil {
		log.Fatalf("rpc.Register error: %v", err)
	}

	// /token endpoint (test only) — JWT beradi
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": "admin",
			"exp":  time.Now().Add(time.Hour).Unix(),
		})
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "failed to sign token", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	})

	// /rpc-json — JWT tekshiriladi, JSON => ichki RPC chaqiriladi
	http.HandleFunc("/rpc-json", func(w http.ResponseWriter, r *http.Request) {
		// JWT tekshirish
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: missing Bearer token", http.StatusUnauthorized)
			return
		}
		reqToken := strings.TrimPrefix(authHeader, "Bearer ")
		if err := validateJWT(reqToken); err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		// JSON body o'qish
		var args Args
		if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
			http.Error(w, "bad request: invalid json", http.StatusBadRequest)
			return
		}
		// Input validation (misol)
		if args.A < 0 || args.B < 0 {
			http.Error(w, "inputs must be non-negative", http.StatusBadRequest)
			return
		}

		// Ichki RPC chaqiruv (to'g'ridan-to'g'ri funksiya chaqirish ham mumkin)
		var result int
		// since arith is local, call method directly instead of rpc client
		if err := arith.Multiply(&args, &result); err != nil {
			http.Error(w, "rpc error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Javobni JSON formatida yuborish
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"result": result,
		})
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
