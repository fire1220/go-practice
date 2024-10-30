package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"unsafe"
)

// 非对称加密的JWT
func main() {
	m := map[string]any{
		"foo":  "bar",
		"name": "jock",
	}
	sig, err := RSASign("./sample_key", m)
	// eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYW1lIjoiam9jayJ9.uB6HfQEgk8yo-vmUmzpDEOqRf3SCrmlb3F9bodoKu2YExD2CtA2WfDK8qv-D2FMED2MUeRarHriCIpYK4-WrHNwPSG-7nlhwstlmVyCk2K2JRXFttAnXZyryF-2dWdwHwY8l7aKJ1nU7O51jYLaEdPIiC3RvRoFh0VSiCKiFDH196X5jA0ot72tTCIRRm-VxSVHdfAJZRHaXLPqKPLQCMaAYDT-3yzyjudGkDPSU6pjK-5qeWwE92-U-daoEdwLcPeYQrEBjJLRV-7QP3JSmqgiRx2fx2vnh_kOU2nBSegDmuocj4DaGFcl7xLFhlg3nlWnwcNQ_nm_xzRlTv6YwiA
	fmt.Println("非对称加密RSA的JWT")
	fmt.Println("生成JWT:")
	fmt.Println(sig, err)
	fmt.Println("验证JWT:")
	fmt.Println(RSAVerify("./sample_key.pub", sig))
}

// RSASign 生成RSA的JWT
func RSASign(privateKeyPath string, m map[string]any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(m))
	// 验证RAS的JWT
	keyData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("read private key file err: %v", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return "", fmt.Errorf("decode private key err: %v", err)
	}
	sig, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}
	return sig, nil
}

// RSAVerify 验证RSA的JWT
func RSAVerify(publicKeyPath, sig string) (string, error) {
	parts := strings.Split(sig, ".")
	jwtHeadByte, err := jwt.NewParser().DecodeSegment(parts[0]) // 解码jwt头信息{"typ":"JWT","alg":"RS256"}
	if err != nil {
		return "", fmt.Errorf("could not decode head: %v", err)
	}
	var jwtHead struct {
		Typ string `json:"typ"`
		Alg string `json:"alg"`
	}
	err = json.Unmarshal(jwtHeadByte, &jwtHead)
	if err != nil {
		return "", fmt.Errorf("head unmarshal err: %v", err)
	}
	if jwtHead.Typ != "JWT" {
		return "", fmt.Errorf("not a complete JWT current is %v", jwtHead.Typ)
	}
	if jwtHead.Alg != "RS256" {
		return "", fmt.Errorf("not JWT of RSA current is %v", jwtHead.Alg)
	}
	jwtBodyByte, err := jwt.NewParser().DecodeSegment(parts[1]) // 解码jwt内容
	if err != nil {
		return "", fmt.Errorf("could not decode body: %v", err)
	}
	jwtBody := *(*string)(unsafe.Pointer(&jwtBodyByte))
	jwtSig, err := jwt.NewParser().DecodeSegment(parts[2]) // 解码签名
	if err != nil {
		return jwtBody, fmt.Errorf("could not decode segment: %v", err)
	}
	// 验证签名
	keyData, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return jwtBody, fmt.Errorf("read public key file err: %v", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return jwtBody, fmt.Errorf("decode public key err: %v", err)
	}
	method := jwt.GetSigningMethod("RS256")
	err = method.Verify(strings.Join(parts[0:2], "."), jwtSig, key)
	return jwtBody, err
}
