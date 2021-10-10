package rates

import (
	"fmt"
	"log"
	"testing"
)

func TestNewRate(t *testing.T) {
	df, err := NewRate(log.Default())
	if err != nil {
		t.Errorf("NewRate() error = %v", err)
		return
	}
	if df == nil {
		t.Fatalf("NewRate() error = %v", err)
		return
	}

	fmt.Printf("NewRate() = %v\n", df.ratings)
}
