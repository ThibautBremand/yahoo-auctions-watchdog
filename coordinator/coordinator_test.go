package coordinator

import (
	"reflect"
	"testing"
	"yahoo-auctions-watchdog/cache"
)

func TestBuildCache(t *testing.T) {
	t.Run("Test with empty cache", func(t *testing.T) {
		lastItemIDs := map[string][]string{
			"url1": {"ID_1", "ID_2"},
		}

		cached := map[string]cache.CachedIDs{}

		got := buildCache(lastItemIDs, cached)
		exp := map[string]cache.CachedIDs{
			"url1": {"url1", "ID_1,ID_2"},
		}

		if !reflect.DeepEqual(exp, got) {
			t.Errorf("expected %+v but got %+v", exp, got)
		}
	})

	t.Run("Test with filled cache but a large number of scraped product IDs", func(t *testing.T) {
		lastItemIDs := map[string][]string{
			"url1": {
				"ID_1",
				"ID_2",
				"ID_3",
				"ID_4",
				"ID_5",
				"ID_6",
				"ID_7",
				"ID_8",
				"ID_9",
				"ID_10",
				"ID_11",
				"ID_12",
				"ID_13",
				"ID_14",
			},
		}

		cached := map[string]cache.CachedIDs{
			"url1": {
				URL:     "url1",
				LastIDs: "ID_A,ID_B,ID_C",
			},
		}

		got := buildCache(lastItemIDs, cached)

		// Only keep the last 10 product IDs since cache.LastIDsSize equals 10
		exp := map[string]cache.CachedIDs{
			"url1": {"url1", "ID_1,ID_2,ID_3,ID_4,ID_5,ID_6,ID_7,ID_8,ID_9,ID_10"},
		}

		if !reflect.DeepEqual(exp, got) {
			t.Errorf("expected %+v but got %+v", exp, got)
		}
	})

	t.Run("Test with small filled cache and a small number of scraped product IDs", func(t *testing.T) {
		lastItemIDs := map[string][]string{
			"url1": {
				"ID_1",
				"ID_2",
			},
		}

		cached := map[string]cache.CachedIDs{
			"url1": {
				URL:     "url1",
				LastIDs: "ID_A,ID_B,ID_C",
			},
		}

		got := buildCache(lastItemIDs, cached)

		// Only keep the last 10 product IDs since cache.LastIDsSize equals 10
		exp := map[string]cache.CachedIDs{
			"url1": {"url1", "ID_1,ID_2,ID_A,ID_B,ID_C"},
		}

		if !reflect.DeepEqual(exp, got) {
			t.Errorf("expected %+v but got %+v", exp, got)
		}
	})

	t.Run("Test with large filled cache and a small number of scraped product IDs", func(t *testing.T) {
		lastItemIDs := map[string][]string{
			"url1": {
				"ID_1",
				"ID_2",
			},
			"url2": {
				"ID_3",
			},
		}

		cached := map[string]cache.CachedIDs{
			"url1": {
				URL:     "url1",
				LastIDs: "ID_A,ID_B,ID_C,ID_D,ID_E,ID_F,ID_G,ID_H,ID_I,ID_J,ID_K,ID_L,ID_M,ID_N,ID_O",
			},
			"url2": {
				URL:     "url2",
				LastIDs: "ID_Z",
			},
		}

		got := buildCache(lastItemIDs, cached)

		// Only keep the last 10 product IDs since cache.LastIDsSize equals 10
		exp := map[string]cache.CachedIDs{
			"url1": {"url1", "ID_1,ID_2,ID_A,ID_B,ID_C,ID_D,ID_E,ID_F,ID_G,ID_H"},
			"url2": {"url2", "ID_3,ID_Z"},
		}

		if !reflect.DeepEqual(exp, got) {
			t.Errorf("expected %+v but got %+v", exp, got)
		}
	})
}
