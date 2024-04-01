package test

import (
	"os"
	"path"

	"github.com/northes/gox"
)

func GenerateTestContent() []byte {
	return []byte(gox.FakeAnimalName())
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
