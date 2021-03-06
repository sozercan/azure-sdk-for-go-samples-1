// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package helpers

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Azure/go-autorest/autorest/utils"
)

// PrintAndLog writes to stdout and to a logger.
func PrintAndLog(message string) {
	log.Println(message)
	fmt.Println(message)
}

// GetRandomLetterSequence returns a sequence of English characters of length n.
func GetRandomLetterSequence(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func contains(array []string, element string) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}
	return false
}

// UserAgent return the string to be appended to user agent header
func UserAgent() string {
	return "samples " + utils.GetCommit()
}
