// Copyright The Wechat Pay Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sign

import (
	"crypto/rsa"
	"math/big"
	"testing"
)

func TestSignatureSHA256WithRSA(t *testing.T) {
	privateKey, err := LoadRSAPrivateKeyFromTxt(mockRSAPrivateKeyCert)
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		req    []byte
		expect string
	}{
		{
			[]byte("GET\n/v3/certificates\n1554208460\n593BEC0C930BF1AFEB40B4A08C8FB242\n\n"),
			"HUa+iOmADqtbJiGOAnzV/9InzUtWYhgYy6DeN5UFW80KueYqm4tLCff99S+FM0xq0zVefIs5/Cy53026Fnz29iNU9e9of97sIaB0j30FB4+YtJqhuk3UJdY/H072VJInMCN0o7GnvIFe9MNfIt77Ak8oJIoCsjRLATjfFxaHaMzTK0VWyG2pQddEg5nX3ezaEirCmmZjTPMOZLb7jsSU+q1C2vEkdbPUu2G64BWlzAXmooDZikWpD+3MnGHw+FCi1/isxtNHAmrbVeqNuxbF0io16ZlkaYhFgexICDJ+rvJrMs2Wpq/NYKWmNScc12o+cPrmwL8BW+Poj0yrKCfs1w==",
		},
		{
			[]byte("POST\n/v3/pay/transactions/native\n1611368330\nAF1404CC2980FB414C99C0B98883BD42\n" + `{"appid":"wx81be3101902f7cb2","mchid":"1601959334","description":"for testing","out_trade_no":"S20210124144305172434","time_expire":"2021-01-24T14:53:05+08:00","attach":"cipher code","notify_url":"https://luoji.live/notify","amount":{"total":1,"currency":"CNY"},"detail":{},"scene_info":{"payer_client_ip":"","store_info":{"id":""}}}` + "\n"),
			"ItuRCG6nAf6ZUi5C5LPa0beCGrG7+G4NdaCHLTmym+UzuZHFgFeqRZ4zKQ0n93qehchFWfQ7s00pgABYvXcOMsV1ld7AUjDTZBPucJK6yhFKz9jd20wtRdDG4LRCZcaTowD2f7LtlixFm8F3/YQaBavxiOe54tc3RX/22flYRzy4YFOpBt+bmjSPZIdSFi53323u7cohwvdHwX+avQCtLZKAUNFJIob66u05BbDEITzYuHjakjpb5btvWemjoZBPxkiETzmd4Oa1y2U+rfFCPZyWT4EV7UxHeEizBL8DkubEBD3KXeArqRX6yoMAU4ywmdFeWDbv1EF0Ndy9hiddZQ==",
		},
	}

	for _, c := range cases {
		signature, err := SignatureSHA256WithRSA(privateKey, c.req)
		if err != nil {
			t.Fatal(err)
		}

		if signature != c.expect {
			t.Fatalf("expect %s, got %s", c.expect, signature)
		}
	}
}

func TestSignatureSHA256WithRSAInvalidPrivateKey(t *testing.T) {
	var privateKey = &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: fromBase10("935393046677"),
			E: 65537,
		},
		D: fromBase10("72663984313281163"),
		Primes: []*big.Int{
			fromBase10("9892036654808464"),
			fromBase10("9456020830884701"),
		},
	}

	cases := []struct {
		req    []byte
		expect string
	}{
		{
			[]byte("GET\n/v3/certificates\n1554208460\n593BEC0C930BF1AFEB40B4A08C8FB242\n\n"),
			"HUa+iOmADqtbJiGOAnzV/9InzUtWYhgYy6DeN5UFW80KueYqm4tLCff99S+FM0xq0zVefIs5/Cy53026Fnz29iNU9e9of97sIaB0j30FB4+YtJqhuk3UJdY/H072VJInMCN0o7GnvIFe9MNfIt77Ak8oJIoCsjRLATjfFxaHaMzTK0VWyG2pQddEg5nX3ezaEirCmmZjTPMOZLb7jsSU+q1C2vEkdbPUu2G64BWlzAXmooDZikWpD+3MnGHw+FCi1/isxtNHAmrbVeqNuxbF0io16ZlkaYhFgexICDJ+rvJrMs2Wpq/NYKWmNScc12o+cPrmwL8BW+Poj0yrKCfs1w==",
		},
		{
			[]byte("POST\n/v3/pay/transactions/native\n1611368330\nAF1404CC2980FB414C99C0B98883BD42\n" + `{"appid":"wx81be3101902f7cb2","mchid":"1601959334","description":"for testing","out_trade_no":"S20210124144305172434","time_expire":"2021-01-24T14:53:05+08:00","attach":"cipher code","notify_url":"https://luoji.live/notify","amount":{"total":1,"currency":"CNY"},"detail":{},"scene_info":{"payer_client_ip":"","store_info":{"id":""}}}` + "\n"),
			"ItuRCG6nAf6ZUi5C5LPa0beCGrG7+G4NdaCHLTmym+UzuZHFgFeqRZ4zKQ0n93qehchFWfQ7s00pgABYvXcOMsV1ld7AUjDTZBPucJK6yhFKz9jd20wtRdDG4LRCZcaTowD2f7LtlixFm8F3/YQaBavxiOe54tc3RX/22flYRzy4YFOpBt+bmjSPZIdSFi53323u7cohwvdHwX+avQCtLZKAUNFJIob66u05BbDEITzYuHjakjpb5btvWemjoZBPxkiETzmd4Oa1y2U+rfFCPZyWT4EV7UxHeEizBL8DkubEBD3KXeArqRX6yoMAU4ywmdFeWDbv1EF0Ndy9hiddZQ==",
		},
	}

	for _, c := range cases {
		_, err := SignatureSHA256WithRSA(privateKey, c.req)
		if err == nil {
			t.Fatalf("should get an err")
		}
	}
}

func TestVerifySHA256WithRSA(t *testing.T) {
	publicKey, err := LoadRSAPublicKeyFromCert([]byte(mockRSAPublicKeyCert))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		signature string
		plain     []byte
		expect    bool
	}{
		{
			"bmdMjyk86N6+BoI8Sf6WEo0oEAgfbLqyQHop7asqdU8p8/RsnVXSoQzwQsyqSUl0mbOichQYpFHXl1Zk/jTNGclsJ49iLBN49pTlnc6bzFTR1qmMkdFMkZ4a0USLzrE9/m8UOSyEp5gT4X4oRtrYgFI0bMujUqGdIGNPgry8YRXvxAAUnE+9mwCY9LFNZxYk6rfbvmMdIjeQar321cmF63Iyq2Vo9Vb//j7wZB9LS/iGGAjOQ2hj9S79u7A9LIfuZDKG6ENIfCUbXabTpog/zFgksuwf821PH3Hy+/7oNepbDcOHrqJUQZ+lPx7h9jfK+yCTd2Rhf/U4w0z2hMBlzA==",
			[]byte("1611501424\n7c6ee840478cacdcf25b8fde1bc492c0\n" +
				`{"data":[{"effective_time":"2020-09-17T14:26:23+08:00","encrypt_certificate":{"algorithm":"AEAD_AES_256_GCM","associated_data":"certificate","ciphertext":"evjNpcxpdo0RxJ377B3SWapXayAVofHD6FF7Alzs01qcO2I8qej8qkiWgSIZWBx05InQJEzqCCKpJqWH7cCoV1Kf6lWa5oyQvAUZSxMbfWCSQ24maNz8mkGs41iwUfR36XpiaSAAhNUPuHhvd/VFZuOYUqEFk9C3m8SzbG0ne7zqLLP7oQi42beASVtz3UGIQu9Pcxm7cyJ/L5AUInvpx+Yq638TVq6A99Il3iDRJKL+C4gXMFplFdk0pVFCH3J6eiu0FbKgEO3fWinKxnbZ6sJHR2TkelCV+lsdb/kyctFOS0YIhlhrNyzDN/IeeiOVH9SD5ffuABv5PX7iA4HGCdR1BTBjeUEGWCTW1xWeo3jN9YAfbZxATQY8iL1LTv1Gkdw/510jn7PL89p+tuwyFlyyXosA/3o7o9W1SA4qZFrjFf3diMoEsEnHlxp3Atm81qvJLwbeqhtrtsLjqEM9o3l7j22dZquxahUfAQ8+7pRgX4tmc52OqliCT8bcEnCPjN8jWTu6KG8QT/rWDJk1tI3O+xXsOrxYMO3dStUs2Pv89JRmVFj0uBizT2lBrnFcvY4wshAcILPqt/lSFxaYlwIlOXf2M5NJ1zqjPTk7lvUyKrrmbTAVcp+PtktWMwz3sgslRs3qbLlPdiN9mUBKdVQpoQ9X5zZcBQwtEM0b324bPXi4jl6zRFPHybvPyJ4dUOe9GpGYNM0EXHsnxf7qdhhn8/TSm1yzlCU3Vw8ey4YdApYk03Dxu7497Rp1JVplKOlYx+XJOpcYlWSyNXgq3QlspBZBk5WwqCU4ENtX9VGGtm0FMtcUc0uEeo3WYSYUVkjuBwWGctzyszSi19R1YoG2wezMu48edCuUHM39FFQYDLDfWQ1nKgi7wNtN9EqA6skHOhEYwbe3A+jp3aUBuQH9cYbOq8MS42SzjGRuNmZiUpg+SxreEZ+f/TSdWQeuneHFrzbF+UY7ntNENNg+S6VGnaSlYP5Geg341QefBgfUeHbtcwtwO35J3AqhYEylyl/PqBq4+Nhme+qb66xiMiJ2i1pGfNCLdrh6qxefRGihuNDJyHhG+1V5y1cpUFubugD709C0EQm6Ebp4gQ93eo5ZeuQA80sQ7zVlVq3i2dA/VbdfznSMHytOBxe/5pPAnAXyQvrY6WYYtLie3UEuAAdQkR4SESpTrE0p9LJZUNkGhmYlYMftu1M2+do6LQJAXTO3xQ8ZW7uyKac9ETMRoH/n26Y6pkoCFH58DqxAbAcZc1VMObI+BKTeLr8iIgGP+6MidI8if0HB1n7cNNKPxB5gG7R0jmEPGJMcxL/gQdsXiRvczHLreOj5DWInxYLRvx/9xmwwKaZ0dSCFV+OLC5fyeQxDgZ3XNtC3pXf5ERmcjwmLANWzPj8EDiPIzYmva/qs1Wtrh0xM+fWuJSwRQt1jMxre7WhU3inRHtEvA2+OkTWgVOsZ0VBuc77Z3l53N4pq2ncNCz8ucs3QnU1ilWcNxE19PV1px+4O28EdQd0izFGOZY73/GIl9+KU9Q8OU9/H2IDsqDC1SV9oM7x1JknmhWu6Jc85XIorKA5nw7bfMwyW+GwOn0bkmynbwnDcb4gddmVuxy91bEZoDQeGbq1lU/Z+ydGaEDRmY0u6/1giQjGC5lWPqia8KN5sdPGNYFT2UEifiR3VofoNsXxohjCXxNRX5Sf94VN6i6/U1nLmPRnIwBGrRjINYlQUYuAHiKpwgU7hUnap4+6fWkjlJD5rH1beU4elJCOKrjDnAFJMtukUWTQaasy+TGU1lgjRAa3dy68a4SBoUm0N7VNO3GWod4YE0UALkoB0Cxo4YUdpO1+j3Toa4m+NsGQhyURAJ5ao7Cvf0gTRaFxIU0COUaME2IEwPQ==","nonce":"eabb3e044577"},"expire_time":"2025-09-16T14:26:23+08:00","serial_no":"477ED0046A54F0360A72A63A8F2816312AAEAB53"}]}` + "\n"),
			true,
		},
		{
			"bmdMjyk86N6+Bo",
			[]byte("1611501424\n7c6ee840478cacdcf25b8fde1bc492c0\n"),
			false,
		},
		{
			"bmdMjyk86N6+BoI8Sf6WEo0oEAgfbLqyQHop7asqdU8p8/RsnVXSoQzwQsyqSUl0mbOichQYpFHXl1Zk/jTNGclsJ49iLBN49pTlnc6bzFTR1qmMkdFMkZ4a0USLzrE9/m8UOSyEp5gT4X4oRtrYgFI0bMujUqGdIGNPgry8YRXvxAAUnE+9mwCY9LFNZxYk6rfbvmMdIjeQar321cmF63Iyq2Vo9Vb//j7wZB9LS/iGGAjOQ2hj9S79u7A9LIfuZDKG6ENIfCUbXabTpog/zFgksuwf821PH3Hy+/7oNepbDcOHrqJUQZ+lPx7h9jfK+yCTd2Rhf/U4w0z2hMBlzA==",
			[]byte("1611501424\n7c6ee840478cacdcf25b8fde1bc492c0\n"),
			false,
		},
	}

	for _, c := range cases {
		err := VerifySHA256WithRSA(publicKey, c.signature, c.plain)
		actual := err == nil
		if actual != c.expect {
			t.Fatalf("expect %v, got %v, %v", c.expect, actual, err)
		}
	}
}

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}
