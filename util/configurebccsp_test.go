//go:build pkcs11
// +build pkcs11

/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package util

import (
	"testing"

	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/bccsp/pkcs11"
	"github.com/stretchr/testify/assert"
)

func TestConfigureBCCSP(t *testing.T) {
	mspDir := t.TempDir()

	lib, pin, label := pkcs11.FindPKCS11Lib()
	opts := &factory.FactoryOpts{
		Default: "PKCS11",
		PKCS11: &pkcs11.PKCS11Opts{
			Security: 256,
			Hash:     "SHA2",
			Library:  lib,
			Label:    label,
			Pin:      pin,
		},
	}

	err := ConfigureBCCSP(&opts, mspDir, "")
	assert.NoError(t, err, "bccsp initialization failed")
}

func TestSanitizePKCS11Opts(t *testing.T) {
	lib, pin, label := pkcs11.FindPKCS11Lib()
	opts := pkcs11.PKCS11Opts{
		Security: 256,
		Hash:     "SHA2",
		Library:  lib,
		Label:    label,
		Pin:      pin,
	}
	p11opts := sanitizePKCS11Opts(opts)
	assert.Equal(t, p11opts.Label, "******")
	assert.Equal(t, p11opts.Pin, "******")
}
