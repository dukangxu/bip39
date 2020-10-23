package bip39

import (
	"testing"
	"fmt"
	"log"
	"encoding/hex"
)

func TestNewEntropy(t *testing.T) {
	var entropy []byte
	var err error

	fmt.Println("================")
	entropy, err = NewEntropy(128)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(len(fmt.Sprintf("%b", entropy)), fmt.Sprintf("%b", entropy))
	}

	fmt.Println("================")
	entropy, err = NewEntropy(129)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(len(fmt.Sprintf("%b", entropy)), fmt.Sprintf("%b", entropy))
	}
	fmt.Println("================")
}

func TestNewMnemonic(t *testing.T) {
	var entropy []byte
	var err error
	entropy, err = NewEntropy(128)
	if err != nil {
		log.Println(err)
		return
	}
	mnemonic, err := NewMnemonic(entropy)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(mnemonic)
}

func TestNewSeed(t *testing.T) {
	var entropy []byte
	var err error
	entropy, err = NewEntropy(128)
	if err != nil {
		log.Println(err)
		return
	}
	mnemonic, err := NewMnemonic(entropy)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(mnemonic)
	seed := NewSeed(mnemonic, "123456")
	fmt.Println(hex.EncodeToString(seed))
}

func TestCheckMnemonic(t *testing.T) {
	mnemonic := "brand year live clerk seven memory glance rural elevator organ kiwi tree" //memory <-> tree
	if err := CheckMnemonic(mnemonic); err != nil {
		log.Println(err)
		return
	}
}
