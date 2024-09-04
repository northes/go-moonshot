package test

import (
	"os"
	"path"
)

func GenerateTestContent() []byte {
	return []byte("夕阳无限好，麦当劳汉堡")
}

func GenerateTestFile(content []byte) (string, error) {
	dir := path.Join(os.TempDir(), "moonshot")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	filePath := path.Join(dir, "test.txt")

	_ = os.Remove(filePath)

	err = os.WriteFile(filePath, content, 0777)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
