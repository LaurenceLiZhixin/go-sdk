package client

import (
	"context"
	"strconv"
	"testing"
	"time"

	"go.uber.org/atomic"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigurationItem(t *testing.T) {
	ctx := context.Background()

	t.Run("get configuration item", func(t *testing.T) {
		resp, err := testClient.GetConfigurationItem(ctx, "example-config", "mykey")
		assert.Nil(t, err)
		assert.Equal(t, "mykey_value", resp.Value)
	})

	t.Run("get configuration item with invalid storeName", func(t *testing.T) {
		_, err := testClient.GetConfigurationItem(ctx, "", "mykey")
		assert.NotNil(t, err)
	})
}

func TestGetConfigurationItems(t *testing.T) {
	ctx := context.Background()

	t.Run("Test get configuration items", func(t *testing.T) {
		resp, err := testClient.GetConfigurationItems(ctx, "example-config", []string{"mykey1", "mykey2", "mykey3"})
		assert.Nil(t, err)
		for i, v := range resp {
			assert.Equal(t, "mykey"+strconv.Itoa(i+1)+"_value", v.Value)
		}
	})
}

func TestSubscribeConfigurationItems(t *testing.T) {
	ctx := context.Background()

	counter := 0
	totalCounter := 0
	t.Run("Test subscribe configuration items", func(t *testing.T) {
		err := testClient.SubscribeConfigurationItems(ctx, "example-config",
			[]string{"mykey", "mykey2", "mykey3"}, func(s string, items []*ConfigurationItem) {
				counter++
				for _, v := range items {
					assert.Equal(t, v.Value, v.Key+"_"+strconv.Itoa(counter-1))
					totalCounter++
				}
			})
		assert.Nil(t, err)
	})
	time.Sleep(time.Second*5 + time.Millisecond*500)
	assert.Equal(t, 5, counter)
	assert.Equal(t, 15, totalCounter)
}

func TestUnSubscribeConfigurationItems(t *testing.T) {
	ctx := context.Background()

	counter := atomic.Int32{}
	totalCounter := atomic.Int32{}
	t.Run("Test unsubscribe configuration items", func(t *testing.T) {
		subscribeID := ""
		subscribeIDChan := make(chan string)
		go func() {
			err := testClient.SubscribeConfigurationItems(ctx, "example-config",
				[]string{"mykey", "mykey2", "mykey3"}, func(id string, items []*ConfigurationItem) {
					counter.Inc()
					for _, v := range items {
						assert.Equal(t, v.Value, v.Key+"_"+strconv.Itoa(int(counter.Load()-1)))
						totalCounter.Inc()
					}
					select {
					case subscribeIDChan <- id:
					default:
					}
				})
			assert.Nil(t, err)
		}()
		subscribeID = <-subscribeIDChan
		time.Sleep(time.Second * 2)
		time.Sleep(time.Millisecond * 500)
		err := testClient.UnsubscribeConfigurationItems(ctx, "example-config", subscribeID)
		assert.Nil(t, err)
	})
	time.Sleep(time.Second * 5)
	assert.Equal(t, 3, int(counter.Load()))
	assert.Equal(t, 9, int(totalCounter.Load()))
}
