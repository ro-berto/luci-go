// Copyright 2021 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package secrets

import (
	"context"
	"encoding/base64"
)

// URLSafeEncrypt encrypts plaintext to an encrypted, URL safe string.
func URLSafeEncrypt(ctx context.Context, plaintext, additionalData []byte) (string, error) {
	encryptedBytes, err := Encrypt(ctx, plaintext, additionalData)
	if err != nil {
		return "", err
	}

	encryptedString := base64.RawURLEncoding.EncodeToString(encryptedBytes)
	return encryptedString, nil
}

// URLSafeDecrypt decrypts the plaintext from the string generated by
// URLSafeEncrypt with the same additional data.
func URLSafeDecrypt(ctx context.Context, encryptedString string, additionalData []byte) ([]byte, error) {
	encryptedBytes, err := base64.RawURLEncoding.DecodeString(encryptedString)
	if err != nil {
		return nil, err
	}

	plaintext, err := Decrypt(ctx, encryptedBytes, additionalData)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
