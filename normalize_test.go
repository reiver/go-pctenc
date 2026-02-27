package pctenc_test

import (
	"testing"

	"github.com/reiver/go-pctenc"
)

func TestNormalize(t *testing.T) {
	tests := []struct{
		PercentEncoded string
		Expected       string
	}{
		{
			PercentEncoded: "",
			Expected:       "",
		},



		{
			PercentEncoded: "apple",
			Expected:       "apple",
		},
		{
			PercentEncoded: "BANANA",
			Expected:       "BANANA",
		},
		{
			PercentEncoded: "Cherry",
			Expected:       "Cherry",
		},



		{
			PercentEncoded: "%00",
			Expected:       "%00",
		},
		{
			PercentEncoded: "%01",
			Expected:       "%01",
		},
		{
			PercentEncoded: "%02",
			Expected:       "%02",
		},
		{
			PercentEncoded: "%03",
			Expected:       "%03",
		},
		{
			PercentEncoded: "%04",
			Expected:       "%04",
		},
		{
			PercentEncoded: "%05",
			Expected:       "%05",
		},
		{
			PercentEncoded: "%06",
			Expected:       "%06",
		},
		{
			PercentEncoded: "%07",
			Expected:       "%07",
		},
		{
			PercentEncoded: "%08",
			Expected:       "%08",
		},
		{
			PercentEncoded: "%09",
			Expected:       "%09",
		},
		{
			PercentEncoded: "%0A",
			Expected:       "%0A",
		},
		{
			PercentEncoded: "%0a",
			Expected:       "%0A",
		},
		{
			PercentEncoded: "%0B",
			Expected:       "%0B",
		},
		{
			PercentEncoded: "%0b",
			Expected:       "%0B",
		},
		{
			PercentEncoded: "%0C",
			Expected:       "%0C",
		},
		{
			PercentEncoded: "%0c",
			Expected:       "%0C",
		},
		{
			PercentEncoded: "%0D",
			Expected:       "%0D",
		},
		{
			PercentEncoded: "%0d",
			Expected:       "%0D",
		},
		{
			PercentEncoded: "%0E",
			Expected:       "%0E",
		},
		{
			PercentEncoded: "%0e",
			Expected:       "%0E",
		},
		{
			PercentEncoded: "%0F",
			Expected:       "%0F",
		},
		{
			PercentEncoded: "%0f",
			Expected:       "%0F",
		},



		{
			PercentEncoded: "%10",
			Expected:       "%10",
		},
		{
			PercentEncoded: "%11",
			Expected:       "%11",
		},
		{
			PercentEncoded: "%12",
			Expected:       "%12",
		},
		{
			PercentEncoded: "%13",
			Expected:       "%13",
		},
		{
			PercentEncoded: "%14",
			Expected:       "%14",
		},
		{
			PercentEncoded: "%15",
			Expected:       "%15",
		},
		{
			PercentEncoded: "%16",
			Expected:       "%16",
		},
		{
			PercentEncoded: "%17",
			Expected:       "%17",
		},
		{
			PercentEncoded: "%18",
			Expected:       "%18",
		},
		{
			PercentEncoded: "%19",
			Expected:       "%19",
		},
		{
			PercentEncoded: "%1A",
			Expected:       "%1A",
		},
		{
			PercentEncoded: "%1a",
			Expected:       "%1A",
		},
		{
			PercentEncoded: "%1B",
			Expected:       "%1B",
		},
		{
			PercentEncoded: "%1b",
			Expected:       "%1B",
		},
		{
			PercentEncoded: "%1C",
			Expected:       "%1C",
		},
		{
			PercentEncoded: "%1c",
			Expected:       "%1C",
		},
		{
			PercentEncoded: "%1D",
			Expected:       "%1D",
		},
		{
			PercentEncoded: "%1d",
			Expected:       "%1D",
		},
		{
			PercentEncoded: "%1E",
			Expected:       "%1E",
		},
		{
			PercentEncoded: "%1e",
			Expected:       "%1E",
		},
		{
			PercentEncoded: "%1F",
			Expected:       "%1F",
		},
		{
			PercentEncoded: "%1f",
			Expected:       "%1F",
		},



		{
			PercentEncoded: "%20",
			Expected:       "%20",
		},
		{
			PercentEncoded: "%21",
			Expected:       "%21",
		},
		{
			PercentEncoded: "%22",
			Expected:       "%22",
		},
		{
			PercentEncoded: "%23",
			Expected:       "%23",
		},
		{
			PercentEncoded: "%24",
			Expected:       "%24",
		},
		{
			PercentEncoded: "%25",
			Expected:       "%25",
		},
		{
			PercentEncoded: "%26",
			Expected:       "%26",
		},
		{
			PercentEncoded: "%27",
			Expected:       "%27",
		},
		{
			PercentEncoded: "%28",
			Expected:       "%28",
		},
		{
			PercentEncoded: "%29",
			Expected:       "%29",
		},
		{
			PercentEncoded: "%2A",
			Expected:       "%2A",
		},
		{
			PercentEncoded: "%2a",
			Expected:       "%2A",
		},
		{
			PercentEncoded: "%2B",
			Expected:       "%2B",
		},
		{
			PercentEncoded: "%2b",
			Expected:       "%2B",
		},
		{
			PercentEncoded: "%2C",
			Expected:       "%2C",
		},
		{
			PercentEncoded: "%2c",
			Expected:       "%2C",
		},
		{
			PercentEncoded: "%2D",
			Expected:       "%2D",
		},
		{
			PercentEncoded: "%2d",
			Expected:       "%2D",
		},
		{
			PercentEncoded: "%2E",
			Expected:       "%2E",
		},
		{
			PercentEncoded: "%2e",
			Expected:       "%2E",
		},
		{
			PercentEncoded: "%2F",
			Expected:       "%2F",
		},
		{
			PercentEncoded: "%2f",
			Expected:       "%2F",
		},



		{
			PercentEncoded: "%30",
			Expected:       "%30",
		},
		{
			PercentEncoded: "%31",
			Expected:       "%31",
		},
		{
			PercentEncoded: "%32",
			Expected:       "%32",
		},
		{
			PercentEncoded: "%33",
			Expected:       "%33",
		},
		{
			PercentEncoded: "%34",
			Expected:       "%34",
		},
		{
			PercentEncoded: "%35",
			Expected:       "%35",
		},
		{
			PercentEncoded: "%36",
			Expected:       "%36",
		},
		{
			PercentEncoded: "%37",
			Expected:       "%37",
		},
		{
			PercentEncoded: "%38",
			Expected:       "%38",
		},
		{
			PercentEncoded: "%39",
			Expected:       "%39",
		},
		{
			PercentEncoded: "%3A",
			Expected:       "%3A",
		},
		{
			PercentEncoded: "%3a",
			Expected:       "%3A",
		},
		{
			PercentEncoded: "%3B",
			Expected:       "%3B",
		},
		{
			PercentEncoded: "%3b",
			Expected:       "%3B",
		},
		{
			PercentEncoded: "%3C",
			Expected:       "%3C",
		},
		{
			PercentEncoded: "%3c",
			Expected:       "%3C",
		},
		{
			PercentEncoded: "%3D",
			Expected:       "%3D",
		},
		{
			PercentEncoded: "%3d",
			Expected:       "%3D",
		},
		{
			PercentEncoded: "%3E",
			Expected:       "%3E",
		},
		{
			PercentEncoded: "%3e",
			Expected:       "%3E",
		},
		{
			PercentEncoded: "%3F",
			Expected:       "%3F",
		},
		{
			PercentEncoded: "%3f",
			Expected:       "%3F",
		},



		{
			PercentEncoded: "%40",
			Expected:       "%40",
		},
		{
			PercentEncoded: "%41",
			Expected:       "%41",
		},
		{
			PercentEncoded: "%42",
			Expected:       "%42",
		},
		{
			PercentEncoded: "%43",
			Expected:       "%43",
		},
		{
			PercentEncoded: "%44",
			Expected:       "%44",
		},
		{
			PercentEncoded: "%45",
			Expected:       "%45",
		},
		{
			PercentEncoded: "%46",
			Expected:       "%46",
		},
		{
			PercentEncoded: "%47",
			Expected:       "%47",
		},
		{
			PercentEncoded: "%48",
			Expected:       "%48",
		},
		{
			PercentEncoded: "%49",
			Expected:       "%49",
		},
		{
			PercentEncoded: "%4A",
			Expected:       "%4A",
		},
		{
			PercentEncoded: "%4a",
			Expected:       "%4A",
		},
		{
			PercentEncoded: "%4B",
			Expected:       "%4B",
		},
		{
			PercentEncoded: "%4b",
			Expected:       "%4B",
		},
		{
			PercentEncoded: "%4C",
			Expected:       "%4C",
		},
		{
			PercentEncoded: "%4c",
			Expected:       "%4C",
		},
		{
			PercentEncoded: "%4D",
			Expected:       "%4D",
		},
		{
			PercentEncoded: "%4d",
			Expected:       "%4D",
		},
		{
			PercentEncoded: "%4E",
			Expected:       "%4E",
		},
		{
			PercentEncoded: "%4e",
			Expected:       "%4E",
		},
		{
			PercentEncoded: "%4F",
			Expected:       "%4F",
		},
		{
			PercentEncoded: "%4f",
			Expected:       "%4F",
		},



		{
			PercentEncoded: "%50",
			Expected:       "%50",
		},
		{
			PercentEncoded: "%51",
			Expected:       "%51",
		},
		{
			PercentEncoded: "%52",
			Expected:       "%52",
		},
		{
			PercentEncoded: "%53",
			Expected:       "%53",
		},
		{
			PercentEncoded: "%54",
			Expected:       "%54",
		},
		{
			PercentEncoded: "%55",
			Expected:       "%55",
		},
		{
			PercentEncoded: "%56",
			Expected:       "%56",
		},
		{
			PercentEncoded: "%57",
			Expected:       "%57",
		},
		{
			PercentEncoded: "%58",
			Expected:       "%58",
		},
		{
			PercentEncoded: "%59",
			Expected:       "%59",
		},
		{
			PercentEncoded: "%5A",
			Expected:       "%5A",
		},
		{
			PercentEncoded: "%5a",
			Expected:       "%5A",
		},
		{
			PercentEncoded: "%5B",
			Expected:       "%5B",
		},
		{
			PercentEncoded: "%5b",
			Expected:       "%5B",
		},
		{
			PercentEncoded: "%5C",
			Expected:       "%5C",
		},
		{
			PercentEncoded: "%5c",
			Expected:       "%5C",
		},
		{
			PercentEncoded: "%5D",
			Expected:       "%5D",
		},
		{
			PercentEncoded: "%5d",
			Expected:       "%5D",
		},
		{
			PercentEncoded: "%5E",
			Expected:       "%5E",
		},
		{
			PercentEncoded: "%5e",
			Expected:       "%5E",
		},
		{
			PercentEncoded: "%5F",
			Expected:       "%5F",
		},
		{
			PercentEncoded: "%5f",
			Expected:       "%5F",
		},



		{
			PercentEncoded: "%60",
			Expected:       "%60",
		},
		{
			PercentEncoded: "%61",
			Expected:       "%61",
		},
		{
			PercentEncoded: "%62",
			Expected:       "%62",
		},
		{
			PercentEncoded: "%63",
			Expected:       "%63",
		},
		{
			PercentEncoded: "%64",
			Expected:       "%64",
		},
		{
			PercentEncoded: "%65",
			Expected:       "%65",
		},
		{
			PercentEncoded: "%66",
			Expected:       "%66",
		},
		{
			PercentEncoded: "%67",
			Expected:       "%67",
		},
		{
			PercentEncoded: "%68",
			Expected:       "%68",
		},
		{
			PercentEncoded: "%69",
			Expected:       "%69",
		},
		{
			PercentEncoded: "%6A",
			Expected:       "%6A",
		},
		{
			PercentEncoded: "%6a",
			Expected:       "%6A",
		},
		{
			PercentEncoded: "%6B",
			Expected:       "%6B",
		},
		{
			PercentEncoded: "%6b",
			Expected:       "%6B",
		},
		{
			PercentEncoded: "%6C",
			Expected:       "%6C",
		},
		{
			PercentEncoded: "%6c",
			Expected:       "%6C",
		},
		{
			PercentEncoded: "%6D",
			Expected:       "%6D",
		},
		{
			PercentEncoded: "%6d",
			Expected:       "%6D",
		},
		{
			PercentEncoded: "%6E",
			Expected:       "%6E",
		},
		{
			PercentEncoded: "%6e",
			Expected:       "%6E",
		},
		{
			PercentEncoded: "%6F",
			Expected:       "%6F",
		},
		{
			PercentEncoded: "%6f",
			Expected:       "%6F",
		},



		{
			PercentEncoded: "%70",
			Expected:       "%70",
		},
		{
			PercentEncoded: "%71",
			Expected:       "%71",
		},
		{
			PercentEncoded: "%72",
			Expected:       "%72",
		},
		{
			PercentEncoded: "%73",
			Expected:       "%73",
		},
		{
			PercentEncoded: "%74",
			Expected:       "%74",
		},
		{
			PercentEncoded: "%75",
			Expected:       "%75",
		},
		{
			PercentEncoded: "%76",
			Expected:       "%76",
		},
		{
			PercentEncoded: "%77",
			Expected:       "%77",
		},
		{
			PercentEncoded: "%78",
			Expected:       "%78",
		},
		{
			PercentEncoded: "%79",
			Expected:       "%79",
		},
		{
			PercentEncoded: "%7A",
			Expected:       "%7A",
		},
		{
			PercentEncoded: "%7a",
			Expected:       "%7A",
		},
		{
			PercentEncoded: "%7B",
			Expected:       "%7B",
		},
		{
			PercentEncoded: "%7b",
			Expected:       "%7B",
		},
		{
			PercentEncoded: "%7C",
			Expected:       "%7C",
		},
		{
			PercentEncoded: "%7c",
			Expected:       "%7C",
		},
		{
			PercentEncoded: "%7D",
			Expected:       "%7D",
		},
		{
			PercentEncoded: "%7d",
			Expected:       "%7D",
		},
		{
			PercentEncoded: "%7E",
			Expected:       "%7E",
		},
		{
			PercentEncoded: "%7e",
			Expected:       "%7E",
		},
		{
			PercentEncoded: "%7F",
			Expected:       "%7F",
		},
		{
			PercentEncoded: "%7f",
			Expected:       "%7F",
		},



		{
			PercentEncoded: "%80",
			Expected:       "%80",
		},
		{
			PercentEncoded: "%81",
			Expected:       "%81",
		},
		{
			PercentEncoded: "%82",
			Expected:       "%82",
		},
		{
			PercentEncoded: "%83",
			Expected:       "%83",
		},
		{
			PercentEncoded: "%84",
			Expected:       "%84",
		},
		{
			PercentEncoded: "%85",
			Expected:       "%85",
		},
		{
			PercentEncoded: "%86",
			Expected:       "%86",
		},
		{
			PercentEncoded: "%87",
			Expected:       "%87",
		},
		{
			PercentEncoded: "%88",
			Expected:       "%88",
		},
		{
			PercentEncoded: "%89",
			Expected:       "%89",
		},
		{
			PercentEncoded: "%8A",
			Expected:       "%8A",
		},
		{
			PercentEncoded: "%8a",
			Expected:       "%8A",
		},
		{
			PercentEncoded: "%8B",
			Expected:       "%8B",
		},
		{
			PercentEncoded: "%8b",
			Expected:       "%8B",
		},
		{
			PercentEncoded: "%8C",
			Expected:       "%8C",
		},
		{
			PercentEncoded: "%8c",
			Expected:       "%8C",
		},
		{
			PercentEncoded: "%8D",
			Expected:       "%8D",
		},
		{
			PercentEncoded: "%8d",
			Expected:       "%8D",
		},
		{
			PercentEncoded: "%8E",
			Expected:       "%8E",
		},
		{
			PercentEncoded: "%8e",
			Expected:       "%8E",
		},
		{
			PercentEncoded: "%8F",
			Expected:       "%8F",
		},
		{
			PercentEncoded: "%8f",
			Expected:       "%8F",
		},



		{
			PercentEncoded: "%90",
			Expected:       "%90",
		},
		{
			PercentEncoded: "%91",
			Expected:       "%91",
		},
		{
			PercentEncoded: "%92",
			Expected:       "%92",
		},
		{
			PercentEncoded: "%93",
			Expected:       "%93",
		},
		{
			PercentEncoded: "%94",
			Expected:       "%94",
		},
		{
			PercentEncoded: "%95",
			Expected:       "%95",
		},
		{
			PercentEncoded: "%96",
			Expected:       "%96",
		},
		{
			PercentEncoded: "%97",
			Expected:       "%97",
		},
		{
			PercentEncoded: "%98",
			Expected:       "%98",
		},
		{
			PercentEncoded: "%99",
			Expected:       "%99",
		},
		{
			PercentEncoded: "%9A",
			Expected:       "%9A",
		},
		{
			PercentEncoded: "%9a",
			Expected:       "%9A",
		},
		{
			PercentEncoded: "%9B",
			Expected:       "%9B",
		},
		{
			PercentEncoded: "%9b",
			Expected:       "%9B",
		},
		{
			PercentEncoded: "%9C",
			Expected:       "%9C",
		},
		{
			PercentEncoded: "%9c",
			Expected:       "%9C",
		},
		{
			PercentEncoded: "%9D",
			Expected:       "%9D",
		},
		{
			PercentEncoded: "%9d",
			Expected:       "%9D",
		},
		{
			PercentEncoded: "%9E",
			Expected:       "%9E",
		},
		{
			PercentEncoded: "%9e",
			Expected:       "%9E",
		},
		{
			PercentEncoded: "%9F",
			Expected:       "%9F",
		},
		{
			PercentEncoded: "%9f",
			Expected:       "%9F",
		},



		{
			PercentEncoded: "%A0",
			Expected:       "%A0",
		},
		{
			PercentEncoded: "%A1",
			Expected:       "%A1",
		},
		{
			PercentEncoded: "%A2",
			Expected:       "%A2",
		},
		{
			PercentEncoded: "%A3",
			Expected:       "%A3",
		},
		{
			PercentEncoded: "%A4",
			Expected:       "%A4",
		},
		{
			PercentEncoded: "%A5",
			Expected:       "%A5",
		},
		{
			PercentEncoded: "%A6",
			Expected:       "%A6",
		},
		{
			PercentEncoded: "%A7",
			Expected:       "%A7",
		},
		{
			PercentEncoded: "%A8",
			Expected:       "%A8",
		},
		{
			PercentEncoded: "%A9",
			Expected:       "%A9",
		},
		{
			PercentEncoded: "%AA",
			Expected:       "%AA",
		},
		{
			PercentEncoded: "%Aa",
			Expected:       "%AA",
		},
		{
			PercentEncoded: "%AB",
			Expected:       "%AB",
		},
		{
			PercentEncoded: "%Ab",
			Expected:       "%AB",
		},
		{
			PercentEncoded: "%AC",
			Expected:       "%AC",
		},
		{
			PercentEncoded: "%Ac",
			Expected:       "%AC",
		},
		{
			PercentEncoded: "%AD",
			Expected:       "%AD",
		},
		{
			PercentEncoded: "%Ad",
			Expected:       "%AD",
		},
		{
			PercentEncoded: "%AE",
			Expected:       "%AE",
		},
		{
			PercentEncoded: "%Ae",
			Expected:       "%AE",
		},
		{
			PercentEncoded: "%AF",
			Expected:       "%AF",
		},
		{
			PercentEncoded: "%Af",
			Expected:       "%AF",
		},



		{
			PercentEncoded: "%a0",
			Expected:       "%A0",
		},
		{
			PercentEncoded: "%a1",
			Expected:       "%A1",
		},
		{
			PercentEncoded: "%a2",
			Expected:       "%A2",
		},
		{
			PercentEncoded: "%a3",
			Expected:       "%A3",
		},
		{
			PercentEncoded: "%a4",
			Expected:       "%A4",
		},
		{
			PercentEncoded: "%a5",
			Expected:       "%A5",
		},
		{
			PercentEncoded: "%a6",
			Expected:       "%A6",
		},
		{
			PercentEncoded: "%a7",
			Expected:       "%A7",
		},
		{
			PercentEncoded: "%a8",
			Expected:       "%A8",
		},
		{
			PercentEncoded: "%a9",
			Expected:       "%A9",
		},
		{
			PercentEncoded: "%aA",
			Expected:       "%AA",
		},
		{
			PercentEncoded: "%aa",
			Expected:       "%AA",
		},
		{
			PercentEncoded: "%aB",
			Expected:       "%AB",
		},
		{
			PercentEncoded: "%ab",
			Expected:       "%AB",
		},
		{
			PercentEncoded: "%aC",
			Expected:       "%AC",
		},
		{
			PercentEncoded: "%ac",
			Expected:       "%AC",
		},
		{
			PercentEncoded: "%aD",
			Expected:       "%AD",
		},
		{
			PercentEncoded: "%ad",
			Expected:       "%AD",
		},
		{
			PercentEncoded: "%aE",
			Expected:       "%AE",
		},
		{
			PercentEncoded: "%ae",
			Expected:       "%AE",
		},
		{
			PercentEncoded: "%aF",
			Expected:       "%AF",
		},
		{
			PercentEncoded: "%af",
			Expected:       "%AF",
		},



		{
			PercentEncoded: "%B0",
			Expected:       "%B0",
		},
		{
			PercentEncoded: "%B1",
			Expected:       "%B1",
		},
		{
			PercentEncoded: "%B2",
			Expected:       "%B2",
		},
		{
			PercentEncoded: "%B3",
			Expected:       "%B3",
		},
		{
			PercentEncoded: "%B4",
			Expected:       "%B4",
		},
		{
			PercentEncoded: "%B5",
			Expected:       "%B5",
		},
		{
			PercentEncoded: "%B6",
			Expected:       "%B6",
		},
		{
			PercentEncoded: "%B7",
			Expected:       "%B7",
		},
		{
			PercentEncoded: "%B8",
			Expected:       "%B8",
		},
		{
			PercentEncoded: "%B9",
			Expected:       "%B9",
		},
		{
			PercentEncoded: "%BA",
			Expected:       "%BA",
		},
		{
			PercentEncoded: "%Ba",
			Expected:       "%BA",
		},
		{
			PercentEncoded: "%BB",
			Expected:       "%BB",
		},
		{
			PercentEncoded: "%Bb",
			Expected:       "%BB",
		},
		{
			PercentEncoded: "%BC",
			Expected:       "%BC",
		},
		{
			PercentEncoded: "%Bc",
			Expected:       "%BC",
		},
		{
			PercentEncoded: "%BD",
			Expected:       "%BD",
		},
		{
			PercentEncoded: "%Bd",
			Expected:       "%BD",
		},
		{
			PercentEncoded: "%BE",
			Expected:       "%BE",
		},
		{
			PercentEncoded: "%Be",
			Expected:       "%BE",
		},
		{
			PercentEncoded: "%BF",
			Expected:       "%BF",
		},
		{
			PercentEncoded: "%Bf",
			Expected:       "%BF",
		},



		{
			PercentEncoded: "%b0",
			Expected:       "%B0",
		},
		{
			PercentEncoded: "%b1",
			Expected:       "%B1",
		},
		{
			PercentEncoded: "%b2",
			Expected:       "%B2",
		},
		{
			PercentEncoded: "%b3",
			Expected:       "%B3",
		},
		{
			PercentEncoded: "%b4",
			Expected:       "%B4",
		},
		{
			PercentEncoded: "%b5",
			Expected:       "%B5",
		},
		{
			PercentEncoded: "%b6",
			Expected:       "%B6",
		},
		{
			PercentEncoded: "%b7",
			Expected:       "%B7",
		},
		{
			PercentEncoded: "%b8",
			Expected:       "%B8",
		},
		{
			PercentEncoded: "%b9",
			Expected:       "%B9",
		},
		{
			PercentEncoded: "%bA",
			Expected:       "%BA",
		},
		{
			PercentEncoded: "%ba",
			Expected:       "%BA",
		},
		{
			PercentEncoded: "%bB",
			Expected:       "%BB",
		},
		{
			PercentEncoded: "%bb",
			Expected:       "%BB",
		},
		{
			PercentEncoded: "%bC",
			Expected:       "%BC",
		},
		{
			PercentEncoded: "%bc",
			Expected:       "%BC",
		},
		{
			PercentEncoded: "%bD",
			Expected:       "%BD",
		},
		{
			PercentEncoded: "%bd",
			Expected:       "%BD",
		},
		{
			PercentEncoded: "%bE",
			Expected:       "%BE",
		},
		{
			PercentEncoded: "%be",
			Expected:       "%BE",
		},
		{
			PercentEncoded: "%bF",
			Expected:       "%BF",
		},
		{
			PercentEncoded: "%bf",
			Expected:       "%BF",
		},



		{
			PercentEncoded: "%C0",
			Expected:       "%C0",
		},
		{
			PercentEncoded: "%C1",
			Expected:       "%C1",
		},
		{
			PercentEncoded: "%C2",
			Expected:       "%C2",
		},
		{
			PercentEncoded: "%C3",
			Expected:       "%C3",
		},
		{
			PercentEncoded: "%C4",
			Expected:       "%C4",
		},
		{
			PercentEncoded: "%C5",
			Expected:       "%C5",
		},
		{
			PercentEncoded: "%C6",
			Expected:       "%C6",
		},
		{
			PercentEncoded: "%C7",
			Expected:       "%C7",
		},
		{
			PercentEncoded: "%C8",
			Expected:       "%C8",
		},
		{
			PercentEncoded: "%C9",
			Expected:       "%C9",
		},
		{
			PercentEncoded: "%CA",
			Expected:       "%CA",
		},
		{
			PercentEncoded: "%Ca",
			Expected:       "%CA",
		},
		{
			PercentEncoded: "%CB",
			Expected:       "%CB",
		},
		{
			PercentEncoded: "%Cb",
			Expected:       "%CB",
		},
		{
			PercentEncoded: "%CC",
			Expected:       "%CC",
		},
		{
			PercentEncoded: "%Cc",
			Expected:       "%CC",
		},
		{
			PercentEncoded: "%CD",
			Expected:       "%CD",
		},
		{
			PercentEncoded: "%Cd",
			Expected:       "%CD",
		},
		{
			PercentEncoded: "%CE",
			Expected:       "%CE",
		},
		{
			PercentEncoded: "%Ce",
			Expected:       "%CE",
		},
		{
			PercentEncoded: "%CF",
			Expected:       "%CF",
		},
		{
			PercentEncoded: "%Cf",
			Expected:       "%CF",
		},



		{
			PercentEncoded: "%c0",
			Expected:       "%C0",
		},
		{
			PercentEncoded: "%c1",
			Expected:       "%C1",
		},
		{
			PercentEncoded: "%c2",
			Expected:       "%C2",
		},
		{
			PercentEncoded: "%c3",
			Expected:       "%C3",
		},
		{
			PercentEncoded: "%c4",
			Expected:       "%C4",
		},
		{
			PercentEncoded: "%c5",
			Expected:       "%C5",
		},
		{
			PercentEncoded: "%c6",
			Expected:       "%C6",
		},
		{
			PercentEncoded: "%c7",
			Expected:       "%C7",
		},
		{
			PercentEncoded: "%c8",
			Expected:       "%C8",
		},
		{
			PercentEncoded: "%c9",
			Expected:       "%C9",
		},
		{
			PercentEncoded: "%cA",
			Expected:       "%CA",
		},
		{
			PercentEncoded: "%ca",
			Expected:       "%CA",
		},
		{
			PercentEncoded: "%cB",
			Expected:       "%CB",
		},
		{
			PercentEncoded: "%cb",
			Expected:       "%CB",
		},
		{
			PercentEncoded: "%cC",
			Expected:       "%CC",
		},
		{
			PercentEncoded: "%cc",
			Expected:       "%CC",
		},
		{
			PercentEncoded: "%cD",
			Expected:       "%CD",
		},
		{
			PercentEncoded: "%cd",
			Expected:       "%CD",
		},
		{
			PercentEncoded: "%cE",
			Expected:       "%CE",
		},
		{
			PercentEncoded: "%ce",
			Expected:       "%CE",
		},
		{
			PercentEncoded: "%cF",
			Expected:       "%CF",
		},
		{
			PercentEncoded: "%cf",
			Expected:       "%CF",
		},



		{
			PercentEncoded: "%D0",
			Expected:       "%D0",
		},
		{
			PercentEncoded: "%D1",
			Expected:       "%D1",
		},
		{
			PercentEncoded: "%D2",
			Expected:       "%D2",
		},
		{
			PercentEncoded: "%D3",
			Expected:       "%D3",
		},
		{
			PercentEncoded: "%D4",
			Expected:       "%D4",
		},
		{
			PercentEncoded: "%D5",
			Expected:       "%D5",
		},
		{
			PercentEncoded: "%D6",
			Expected:       "%D6",
		},
		{
			PercentEncoded: "%D7",
			Expected:       "%D7",
		},
		{
			PercentEncoded: "%D8",
			Expected:       "%D8",
		},
		{
			PercentEncoded: "%D9",
			Expected:       "%D9",
		},
		{
			PercentEncoded: "%DA",
			Expected:       "%DA",
		},
		{
			PercentEncoded: "%Da",
			Expected:       "%DA",
		},
		{
			PercentEncoded: "%DB",
			Expected:       "%DB",
		},
		{
			PercentEncoded: "%Db",
			Expected:       "%DB",
		},
		{
			PercentEncoded: "%DC",
			Expected:       "%DC",
		},
		{
			PercentEncoded: "%Dc",
			Expected:       "%DC",
		},
		{
			PercentEncoded: "%DD",
			Expected:       "%DD",
		},
		{
			PercentEncoded: "%Dd",
			Expected:       "%DD",
		},
		{
			PercentEncoded: "%DE",
			Expected:       "%DE",
		},
		{
			PercentEncoded: "%De",
			Expected:       "%DE",
		},
		{
			PercentEncoded: "%DF",
			Expected:       "%DF",
		},
		{
			PercentEncoded: "%Df",
			Expected:       "%DF",
		},



		{
			PercentEncoded: "%d0",
			Expected:       "%D0",
		},
		{
			PercentEncoded: "%d1",
			Expected:       "%D1",
		},
		{
			PercentEncoded: "%d2",
			Expected:       "%D2",
		},
		{
			PercentEncoded: "%d3",
			Expected:       "%D3",
		},
		{
			PercentEncoded: "%d4",
			Expected:       "%D4",
		},
		{
			PercentEncoded: "%d5",
			Expected:       "%D5",
		},
		{
			PercentEncoded: "%d6",
			Expected:       "%D6",
		},
		{
			PercentEncoded: "%d7",
			Expected:       "%D7",
		},
		{
			PercentEncoded: "%d8",
			Expected:       "%D8",
		},
		{
			PercentEncoded: "%d9",
			Expected:       "%D9",
		},
		{
			PercentEncoded: "%dA",
			Expected:       "%DA",
		},
		{
			PercentEncoded: "%da",
			Expected:       "%DA",
		},
		{
			PercentEncoded: "%dB",
			Expected:       "%DB",
		},
		{
			PercentEncoded: "%db",
			Expected:       "%DB",
		},
		{
			PercentEncoded: "%dC",
			Expected:       "%DC",
		},
		{
			PercentEncoded: "%dc",
			Expected:       "%DC",
		},
		{
			PercentEncoded: "%dD",
			Expected:       "%DD",
		},
		{
			PercentEncoded: "%dd",
			Expected:       "%DD",
		},
		{
			PercentEncoded: "%dE",
			Expected:       "%DE",
		},
		{
			PercentEncoded: "%de",
			Expected:       "%DE",
		},
		{
			PercentEncoded: "%dF",
			Expected:       "%DF",
		},
		{
			PercentEncoded: "%df",
			Expected:       "%DF",
		},



		{
			PercentEncoded: "%E0",
			Expected:       "%E0",
		},
		{
			PercentEncoded: "%E1",
			Expected:       "%E1",
		},
		{
			PercentEncoded: "%E2",
			Expected:       "%E2",
		},
		{
			PercentEncoded: "%E3",
			Expected:       "%E3",
		},
		{
			PercentEncoded: "%E4",
			Expected:       "%E4",
		},
		{
			PercentEncoded: "%E5",
			Expected:       "%E5",
		},
		{
			PercentEncoded: "%E6",
			Expected:       "%E6",
		},
		{
			PercentEncoded: "%E7",
			Expected:       "%E7",
		},
		{
			PercentEncoded: "%E8",
			Expected:       "%E8",
		},
		{
			PercentEncoded: "%E9",
			Expected:       "%E9",
		},
		{
			PercentEncoded: "%EA",
			Expected:       "%EA",
		},
		{
			PercentEncoded: "%Ea",
			Expected:       "%EA",
		},
		{
			PercentEncoded: "%EB",
			Expected:       "%EB",
		},
		{
			PercentEncoded: "%Eb",
			Expected:       "%EB",
		},
		{
			PercentEncoded: "%EC",
			Expected:       "%EC",
		},
		{
			PercentEncoded: "%Ec",
			Expected:       "%EC",
		},
		{
			PercentEncoded: "%ED",
			Expected:       "%ED",
		},
		{
			PercentEncoded: "%Ed",
			Expected:       "%ED",
		},
		{
			PercentEncoded: "%EE",
			Expected:       "%EE",
		},
		{
			PercentEncoded: "%Ee",
			Expected:       "%EE",
		},
		{
			PercentEncoded: "%EF",
			Expected:       "%EF",
		},
		{
			PercentEncoded: "%Ef",
			Expected:       "%EF",
		},



		{
			PercentEncoded: "%e0",
			Expected:       "%E0",
		},
		{
			PercentEncoded: "%e1",
			Expected:       "%E1",
		},
		{
			PercentEncoded: "%e2",
			Expected:       "%E2",
		},
		{
			PercentEncoded: "%e3",
			Expected:       "%E3",
		},
		{
			PercentEncoded: "%e4",
			Expected:       "%E4",
		},
		{
			PercentEncoded: "%e5",
			Expected:       "%E5",
		},
		{
			PercentEncoded: "%e6",
			Expected:       "%E6",
		},
		{
			PercentEncoded: "%e7",
			Expected:       "%E7",
		},
		{
			PercentEncoded: "%e8",
			Expected:       "%E8",
		},
		{
			PercentEncoded: "%e9",
			Expected:       "%E9",
		},
		{
			PercentEncoded: "%eA",
			Expected:       "%EA",
		},
		{
			PercentEncoded: "%ea",
			Expected:       "%EA",
		},
		{
			PercentEncoded: "%eB",
			Expected:       "%EB",
		},
		{
			PercentEncoded: "%eb",
			Expected:       "%EB",
		},
		{
			PercentEncoded: "%eC",
			Expected:       "%EC",
		},
		{
			PercentEncoded: "%ec",
			Expected:       "%EC",
		},
		{
			PercentEncoded: "%eD",
			Expected:       "%ED",
		},
		{
			PercentEncoded: "%ed",
			Expected:       "%ED",
		},
		{
			PercentEncoded: "%eE",
			Expected:       "%EE",
		},
		{
			PercentEncoded: "%ee",
			Expected:       "%EE",
		},
		{
			PercentEncoded: "%eF",
			Expected:       "%EF",
		},
		{
			PercentEncoded: "%ef",
			Expected:       "%EF",
		},



		{
			PercentEncoded: "%F0",
			Expected:       "%F0",
		},
		{
			PercentEncoded: "%F1",
			Expected:       "%F1",
		},
		{
			PercentEncoded: "%F2",
			Expected:       "%F2",
		},
		{
			PercentEncoded: "%F3",
			Expected:       "%F3",
		},
		{
			PercentEncoded: "%F4",
			Expected:       "%F4",
		},
		{
			PercentEncoded: "%F5",
			Expected:       "%F5",
		},
		{
			PercentEncoded: "%F6",
			Expected:       "%F6",
		},
		{
			PercentEncoded: "%F7",
			Expected:       "%F7",
		},
		{
			PercentEncoded: "%F8",
			Expected:       "%F8",
		},
		{
			PercentEncoded: "%F9",
			Expected:       "%F9",
		},
		{
			PercentEncoded: "%FA",
			Expected:       "%FA",
		},
		{
			PercentEncoded: "%Fa",
			Expected:       "%FA",
		},
		{
			PercentEncoded: "%FB",
			Expected:       "%FB",
		},
		{
			PercentEncoded: "%Fb",
			Expected:       "%FB",
		},
		{
			PercentEncoded: "%FC",
			Expected:       "%FC",
		},
		{
			PercentEncoded: "%Fc",
			Expected:       "%FC",
		},
		{
			PercentEncoded: "%FD",
			Expected:       "%FD",
		},
		{
			PercentEncoded: "%Fd",
			Expected:       "%FD",
		},
		{
			PercentEncoded: "%FE",
			Expected:       "%FE",
		},
		{
			PercentEncoded: "%Fe",
			Expected:       "%FE",
		},
		{
			PercentEncoded: "%FF",
			Expected:       "%FF",
		},
		{
			PercentEncoded: "%Ff",
			Expected:       "%FF",
		},



		{
			PercentEncoded: "%f0",
			Expected:       "%F0",
		},
		{
			PercentEncoded: "%f1",
			Expected:       "%F1",
		},
		{
			PercentEncoded: "%f2",
			Expected:       "%F2",
		},
		{
			PercentEncoded: "%f3",
			Expected:       "%F3",
		},
		{
			PercentEncoded: "%f4",
			Expected:       "%F4",
		},
		{
			PercentEncoded: "%f5",
			Expected:       "%F5",
		},
		{
			PercentEncoded: "%f6",
			Expected:       "%F6",
		},
		{
			PercentEncoded: "%f7",
			Expected:       "%F7",
		},
		{
			PercentEncoded: "%f8",
			Expected:       "%F8",
		},
		{
			PercentEncoded: "%f9",
			Expected:       "%F9",
		},
		{
			PercentEncoded: "%fA",
			Expected:       "%FA",
		},
		{
			PercentEncoded: "%fa",
			Expected:       "%FA",
		},
		{
			PercentEncoded: "%fB",
			Expected:       "%FB",
		},
		{
			PercentEncoded: "%fb",
			Expected:       "%FB",
		},
		{
			PercentEncoded: "%fC",
			Expected:       "%FC",
		},
		{
			PercentEncoded: "%fc",
			Expected:       "%FC",
		},
		{
			PercentEncoded: "%fD",
			Expected:       "%FD",
		},
		{
			PercentEncoded: "%fd",
			Expected:       "%FD",
		},
		{
			PercentEncoded: "%fE",
			Expected:       "%FE",
		},
		{
			PercentEncoded: "%fe",
			Expected:       "%FE",
		},
		{
			PercentEncoded: "%fF",
			Expected:       "%FF",
		},
		{
			PercentEncoded: "%ff",
			Expected:       "%FF",
		},



		{
			PercentEncoded: "Hi, please %fe%ed%fa%ce%f0%0d! 😈",
			Expected:       "Hi, please %FE%ED%FA%CE%F0%0D! 😈",
		},
		{
			PercentEncoded: "Hi, please %Fe%ed my %Fa%ce %F0%0d! 😈",
			Expected:       "Hi, please %FE%ED my %FA%CE %F0%0D! 😈",
		},
	}

	for testNumber, test := range tests {
		actual := pctenc.Normalize(test.PercentEncoded)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual normalized percent-encoded is not what was expected.", testNumber)
			t.Logf("EXPECTED:        %q", expected)
			t.Logf("ACTUAL:          %q", actual)
			t.Logf("PERCENT-ENCODED: %q", test.PercentEncoded)
			continue
		}
	}
}
