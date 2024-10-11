package main

import (
	"cpiaa/structs"
	"testing"
)

// Testa a função Ping
func TestPing(t *testing.T) {
	cmdr := structs.NewCommander()
	result, err := cmdr.Ping("google.com")

	if err != nil {
		t.Fatalf("Erro inesperado em Ping: %v", err)
	}

	if !result.Successful {
		t.Error("Esperava sucesso no ping, mas falhou.")
	}

	if result.Time <= 0 {
		t.Error("Esperava tempo de resposta maior que zero.")
	}
}

// Testa a função GetSystemInfo
func TestGetSystemInfo(t *testing.T) {
	cmdr := structs.NewCommander()
	info, err := cmdr.GetSystemInfo()

	if err != nil {
		t.Fatalf("Esperava sem erro, mas encontrou: %v", err)
	}

	if info.Hostname == "" {
		t.Error("Esperava que o hostname não fosse vazio.")
	}

	if info.IPAddress == "" {
		t.Error("Esperava que o endereço IP não fosse vazio.")
	}
}

// Teste simulado para IP local
func TestLocalIPAddress(t *testing.T) {
	ip, err := structs.GetLocalIPAddress()
	if err != nil {
		t.Fatalf("Erro inesperado ao obter IP local: %v", err)
	}

	if ip == "" {
		t.Error("Esperava um endereço IP não vazio.")
	}
}
