package main

import (
	"testing"
)

func TestGetMostCommonWords_BasicFunctionality(t *testing.T) {
	words := []string{"hello", "world", "hello", "go", "world", "hello"}
	K := 2
	result, err := GetMostCommonWords(words, K)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 words, got %d", len(result))
	}
	if len(result) > 0 && result[0] != "hello" {
		t.Errorf("Expected 'hello' as first word, got %v", result)
	}
}

func TestGetMostCommonWords_EmptyInput(t *testing.T) {
	words := []string{}
	K := 3
	result, err := GetMostCommonWords(words, K)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected result to be empty, got %v", result)
	}
}

func TestGetMostCommonWords_SingleWord(t *testing.T) {
	words := []string{"test", "test", "test"}
	K := 1
	result, err := GetMostCommonWords(words, K)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != 1 || result[0] != "test" {
		t.Errorf("Expected ['test'], got %v", result)
	}
}

func TestGetMostCommonWords_SameFrequency(t *testing.T) {
	words := []string{"a", "a", "b", "b", "c", "c"}
	K := 2
	result, err := GetMostCommonWords(words, K)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 words, got %d", len(result))
	}
	if len(result) >= 2 && result[0] > result[1] {
		t.Errorf("Expected lexicographical order, got %v", result)
	}
}

func TestGetMostCommonWords_InvalidK(t *testing.T) {
	words := []string{"hello", "world"}
	K := 0
	_, err := GetMostCommonWords(words, K)
	if err == nil {
		t.Error("Expected error for K <= 0, but got none")
	}
}

func TestGetMostCommonWords_KGreaterThanUniqueWords(t *testing.T) {
	words := []string{"a", "a", "b", "b", "c"}
	K := 5
	result, err := GetMostCommonWords(words, K)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("Expected 3 words, got %d", len(result))
	}
}

func TestGetMostCommonWordsLexicographicalOrder(t *testing.T) {
	words := []string{"apple", "banana", "apple", "banana", "cherry"}
	K := 3
	result, err := GetMostCommonWords(words, K)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("Expected 3 words, got %d", len(result))
	}
	if len(result) >= 2 && result[0] == "banana" && result[1] == "apple" {
		t.Errorf("Expected lexicographical order, got: %v", result)
	}
}
