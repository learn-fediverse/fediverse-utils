package rsahelpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

const (
	RSAPrivateKeyLabel = "RSA PRIVATE KEY"
	PublicKeyLabel     = "PUBLIC KEY"
)

func GenerateRSPrivateKey(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func PrivateKeyToPKCS1PEMBlock(key *rsa.PrivateKey) *pem.Block {
	return &pem.Block{
		Type:  RSAPrivateKeyLabel,
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
}

func PrivateKeyToPKCS1PEMString(key *rsa.PrivateKey) string {
	return string(pem.EncodeToMemory(PrivateKeyToPKCS1PEMBlock(key)))
}

func PublicKeyToPKIXBlock(key *rsa.PublicKey) (*pem.Block, error) {
	publicKeyPEMBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return nil, err
	}
	return &pem.Block{
		Type:  PublicKeyLabel,
		Bytes: publicKeyPEMBytes,
	}, nil
}

func PublicKeyToPKIXString(key *rsa.PublicKey) (string, error) {
	publicKeyPEMBlock, err := PublicKeyToPKIXBlock(key)
	if err != nil {
		return "", err
	}
	return string(pem.EncodeToMemory(publicKeyPEMBlock)), nil
}

// So this library does not provide a way to parse neither the PKCS1 PEM private
// key, nor the PKIX PEM public key. This is because parsing has to places where
// it can fail: the PEM decoding, and the PKCS1 or PKIX parsing. Probably a good
// idea to force the client code to handle that.
//
// To parse the private key:
//
//     block, _ := pem.Decode([]byte(content))
//     if block == nil || block.Type != "RSA PRIVATE KEY" {
//         // Handle error
//     }
//     privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
//     if err != nil {
//       // Handle error
//     }
//
// To parse the public key:
//
//     block, _ := pem.Decode([]byte(content))
//     if block == nil || block.Type != "PUBLIC KEY" {
//         // Handle error
//     }
//     publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
//     if err != nil {
//       // Handle error
//     }