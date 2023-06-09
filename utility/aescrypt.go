package utility

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

// Encrypt 使用给定密钥加密字符串。
func Encrypt(plaintext string) (string, error) {

	loading, err := ReadConfigFile([]string{"aes.key"} /*, "manifest/config/"*/)
	if err != nil {
		return "", err
	}

	key := []byte(loading["aes.key"])
	paddingLength := aes.BlockSize - len(plaintext)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	plaintext = plaintext + string(padding)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:i+aes.BlockSize], []byte(plaintext[i:i+aes.BlockSize]))
	}

	result := strings.TrimRight(base64.StdEncoding.EncodeToString(ciphertext), "=")

	return result, err
}

// AesDecrypt 使用给定的密钥解密字符串。
func AesDecrypt(cipher string) (string, error) {
	loading, err := ReadConfigFile([]string{"aes.key"} /*, "manifest/config/"*/)
	if err != nil {
		return "", errors.New("AesDecrypt:读取配置文件失败")
	}
	key := []byte(loading["aes.key"])

	// The ciphertext to be decrypted
	ciphertext, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return "", errors.New("AesDecrypt:密文有误解密失败")
	}
	// Create a new AES decryption block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.New("AesDecrypt:创建AES解密块失败")
	}

	// Create a decryptor in ECB mode
	mode := newECBDecrypter(block)

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove the filled part
	plaintext = pKCS7Unpad(plaintext, block.BlockSize())

	return string(plaintext), err
}

// NewECBDecrypter 创建一个新的 ECB 模式解密器
func newECBDecrypter(block cipher.Block) cipher.BlockMode {
	return &ecbDecrypter{block}
}

// DecryptStructure 解密后的结构体
func DecryptStructure(req string, structure interface{}) error {
	aesDecrypt, err := AesDecrypt(req)
	if err != nil {
		return err
	}

	// 将解密后的数据转换为结构体数据
	err = json.Unmarshal([]byte(aesDecrypt), &structure)
	if err != nil {
		return errors.New("密文转换为结构体失败")
	}
	return nil
}

// ReadConfigFile 解读配置文件的数据
func ReadConfigFile(configNames []string) (map[string]string, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("manifest/config/")
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}
	configs := make(map[string]string, len(configNames))
	for _, name := range configNames {
		configs[name] = v.GetString(name)
	}
	return configs, nil
}

// PKCS7Unpad 删除 PKCS 7 填充
func pKCS7Unpad(data []byte, blockSize int) []byte {
	if len(data) == 0 {
		return []byte{}
	}
	padding := int(data[len(data)-1])
	if padding < 1 || padding > blockSize {
		return data
	}
	for i := len(data) - 1; i > len(data)-padding-1; i-- {
		if int(data[i]) != padding {
			return data
		}
	}
	return data[:len(data)-padding]
}

// RandomNumber 随机生成19位随机数
func RandomNumber() *big.Int {
	for true {
		rand.Seed(time.Now().UnixNano() + time.Now().UnixMicro())
		max := big.NewInt(0)
		max.SetString("99999999999999999999", 10)
		randomNumber := rand.Intn(int(max.Int64()))
		newNumber := big.NewInt(int64(randomNumber))

		if len(newNumber.Text(10)) == 19 {
			return big.NewInt(int64(randomNumber))
		}
	}
	return nil
}

type ecbDecrypter struct {
	b cipher.Block
}

func (x *ecbDecrypter) BlockSize() int { return x.b.BlockSize() }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.b.BlockSize() != 0 {
		panic("input not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst[:x.b.BlockSize()], src[:x.b.BlockSize()])
		src = src[x.b.BlockSize():]
		dst = dst[x.b.BlockSize():]
	}
}
