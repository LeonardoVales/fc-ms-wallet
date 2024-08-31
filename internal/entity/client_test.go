package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "eu@gmail.com")
	if err != nil {
		t.Error("Error creating new client", err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "eu@gmail.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "eu@gmail.com")
	err := client.Update("John Doe Update", "nos@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "nos@gmail.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "eu@gmail.com")
	err := client.Update("", "")
	assert.Error(t, err, "name is required")

}
