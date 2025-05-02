package utils

import "github.com/docker/docker/pkg/namesgenerator"

func RandomName() string {
	return namesgenerator.GetRandomName(5)
}
