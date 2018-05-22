package visor

import (
	"github.com/skycoin/skycoin/src/coin"
)

const (
	// MaxCoinSupply is the maximum supply of skycoins
	MaxCoinSupply uint64 = 1e8 // 100,000,000 million

	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100

	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal

	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 25

	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5

	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 60 * 60 * 24 * 365 // 1 year
)

func init() {
	if MaxCoinSupply%DistributionAddressesTotal != 0 {
		panic("MaxCoinSupply should be perfectly divisible by DistributionAddressesTotal")
	}
}

// GetDistributionAddresses returns a copy of the hardcoded distribution addresses array.
// Each address has 1,000,000 coins. There are 100 addresses.
func GetDistributionAddresses() []string {
	addrs := make([]string, len(distributionAddresses))
	for i := range distributionAddresses {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// GetUnlockedDistributionAddresses returns distribution addresses that are unlocked, i.e. they have spendable outputs
func GetUnlockedDistributionAddresses() []string {
	// The first InitialUnlockedCount (25) addresses are unlocked by default.
	// Subsequent addresses will be unlocked at a rate of UnlockAddressRate (5) per year,
	// after the InitialUnlockedCount (25) addresses have no remaining balance.
	// The unlock timer will be enabled manually once the
	// InitialUnlockedCount (25) addresses are distributed.

	// NOTE: To have automatic unlocking, transaction verification would have
	// to be handled in visor rather than in coin.Transactions.Visor(), because
	// the coin package is agnostic to the state of the blockchain and cannot reference it.
	// Instead of automatic unlocking, we can hardcode the timestamp at which the first 30%
	// is distributed, then compute the unlocked addresses easily here.

	addrs := make([]string, InitialUnlockedCount)
	for i := range distributionAddresses[:InitialUnlockedCount] {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// GetLockedDistributionAddresses returns distribution addresses that are locked, i.e. they have unspendable outputs
func GetLockedDistributionAddresses() []string {
	// TODO -- once we reach 30% distribution, we can hardcode the
	// initial timestamp for releasing more coins
	addrs := make([]string, DistributionAddressesTotal-InitialUnlockedCount)
	for i := range distributionAddresses[InitialUnlockedCount:] {
		addrs[i] = distributionAddresses[InitialUnlockedCount+uint64(i)]
	}
	return addrs
}

// TransactionIsLocked returns true if the transaction spends locked outputs
func TransactionIsLocked(inUxs coin.UxArray) bool {
	lockedAddrs := GetLockedDistributionAddresses()
	lockedAddrsMap := make(map[string]struct{})
	for _, a := range lockedAddrs {
		lockedAddrsMap[a] = struct{}{}
	}

	for _, o := range inUxs {
		uxAddr := o.Body.Address.String()
		if _, ok := lockedAddrsMap[uxAddr]; ok {
			return true
		}
	}

	return false
}

var distributionAddresses = [DistributionAddressesTotal]string{
	"2Er4U9oGkXHjiDG7Qy7GzqEue1pFwCDVF29",
	"2HFdJaD5cymA7h6UMz4mSdMgAh7VWmH2J9",
	"2VUXRME4YPZvA77Vk3jEGada1t4tv9rYDHB",
	"2BR1fkewtk8rUziC16YKf2tFYFgpKECvUTo",
	"PRtJmxGpvZci2JqvHdqRk1tpM83yzNqM67",
	"7UYvSrKbmZpGr41pYroCNpEU9uPbPceyxg",
	"Rm71bdJrFhVeRDmuNyPMtwxnLsRePNWgK7",
	"2mFiEqBbKFdBzwfaewef4ha3T4c8ox7fDMg",
	"5dSsHrdGUAAJWUPc24HQonaHTz7oJTG2Wi",
	"6cuiVPTs3MggwpVucyJTo2PrgjHVSaGfHb",
	"fQjtswpNzYUCdkMLV6M2L9k175gxbfXeHV",
	"2EpsC92idYwcXq2ZSjHjjVvirSBpoCExHHm",
	"pTjvk2yY8KgvTRkyaPycWHk8YUmChmgWDq",
	"2cTpFcorfZnrP9skmjo1hj9DqEMuRgDwLi1",
	"2fE4YEwukzkVPj2BY3StudKCnQtbpwciMKP",
	"vbFEYA3PuLNqiNywYWx4YmVhmTAwmy3QcJ",
	"t2GVtYN9sZudpe1DuaGJtxEzV2MVPK9NKK",
	"hauFHDucYHHgE4Rs4AD4YtNeU3ynPA3gtW",
	"NNjVBuhuFrMzhtNvjkmb9osAMFmeExjWwq",
	"84qnnFVHE1YSENdVPxHnnsD2MYN8UZps42",
	"22paaZ4gb5VigZ3VaArPuAiwHBxUA2SyZag",
	"fwmVzBSCYa6qX7q8491FZetMrqoq4QdaM2",
	"2fncFU7tWArkrc2uo1DudjDqSRHq58HLgiX",
	"2Quz8vYgXJZ9dBHRheu5hTrGiTNiyy5zT1f",
	"26jfxnyj4tSFRLosciuxWHp4qxyQFWZ1jm6",
	"EaoFR1yFVgDxxQQTPyHMutu1D87FotcMqX",
	"TCoTfXqcNkSatLjBS8V7nTQbb8zp9znuPi",
	"mjQM3c3S3HpgAWktbDGyT6FvvuAUNEpJ1i",
	"4wBee84yRDbxyS1WdmCFu7gQQuyffdPpC1",
	"WXaJciFpmWTUcvuGBEgEXkV9DUFAWjt19R",
	"k2AisKVhdZhV9pMkqo8MjdRsQcyoCMqhwj",
	"EuVhtyKZUVxt9MUwhCoRUcE2TWnrFj18xt",
	"2QjK5Sa5AsLyze61jNZpgrtQkXqecF6eqUu",
	"2ey5BodW3Q9vcSRjpcshZLcNj1koE9t6ik",
	"2CSrFX45rrure66qxL1tySjWyPwWnJ3KtXE",
	"2U2uyJ6e2kEEKED6MJk8jujRKyZG5Z8xjN2",
	"zhCWa3hPQ1ZENLRiSiyqT3HSqM1JBZpypG",
	"2EtNyGYsDvHKzEkAZzZiU6WVMzsfw4etAeM",
	"2DXP9gQgY4WpfhWR87YWzweuDJso64nNwj9",
	"2HcfNHM3xZXuit3MdHJ6NTFXAWAuAQbwUPU",
	"2KHN9L13L25QKibw7zhgnkkugNsDH9ojFp1",
	"2dv6WF215oX7c3TevBjB4h5mtR8LznwTYyH",
	"2fsmGC9d6Pyz9TFXkc61xrRXRSvnEopEw6d",
	"27nTPrhvMitUH531wRUP5GdGRpCi14aeyks",
	"24ugZDZ6YvuR4QswxhvMRZBWj46obRApHmL",
	"EYGaGACHUku2gGLh9PJ5jGKTjBhNDrpLF8",
	"BYPcaxFsdWRpfxNEbsURwktCBWWLUNBr8c",
	"28fpPtv7zwPGy6TDBsEEQp4DLab9qMWinPH",
	"27LwR7LwwyjZAE86H1ELB9E5m1JLhnhhHCv",
	"riZMqzoW1kNGeqWgV7rjczDFtTfVuB4VpR",
	"kwQpv22ZyYyD7sqKtaJH3KDdTPZcDyRMdq",
	"CnthD3PWRGDhn2R2RqcM2P7D2ZAjhdDCZC",
	"VexSYY2HQoRcNKyuZ2pwavgSg2dtSK7K1e",
	"2ByP2cdgMA77eauDARF65jngdsnh6WGHX3k",
	"2emCAxEteKzh98QxFsDKQL5a8LP6CCHoHzL",
	"6tYJFXZt4wdrgT5TNEzjAJ4LvUVh9sMFax",
	"E2K7MPeG13LKetPp5jiTyHq4C7wZ8noBFn",
	"2Y5VkKz4cNkrSDNYbyGXhc94RYdrn3cjUqd",
	"tknovpmGY9ofRaMG9TEMsRb5YTLUwCT3g",
	"Y9MFpE6grZWgfYSX2ZQM9aeMArmefPHzoP",
	"MQKseDA9AWpgDH6Hfjm41MzCsq6fp94ahL",
	"2Ci7ky6oabh6aRQTqLLM6oFnhsBdfgwySwH",
	"W6qQ5K25kGKhJ1guEsZ7nGQVu2xBgHEiXG",
	"2bqtYBHyygudPapgEudnaPmq26mZb5dMqMD",
	"8w88mzHAhPTFYJoSUDV4YBkpJkxL99Q8o5",
	"kEqJHuDDDvcjJkHwY2AoKeRk5H1cc5qWPP",
	"2dbT96irJAK2A8CbKHCbsVdN6ZrXccB7s9m",
	"Eqfc8BjPPKHd4YFqxpYmnpS98ZbSBNtvKB",
	"27S6enbEVC6FrMpZSAikLXhuci4UsikJwBT",
	"D55faitDpLuHLKhsreFqsWhRZr7cXmtkSR",
	"2YRp8wcznqfMET1Z45g57edi1dSREw8jPNd",
	"DCMjqjdJJscnjdNYcDWWjsWLW5swcdJSqZ",
	"zhmx9Byh3vCFpnMLqXUJAuKnWdZEKvmhFL",
	"27h1iJaQo5tVWcxdeB437vKoN1RrCAhR2FD",
	"2kmT1EygE5DgsbXNiUxhQVMQjbrJKJHuH36",
	"bUxv6M8YzAZ9pCbbmkXLnfqu9TpunpJrpe",
	"YbUB28QSq6bUfe96HHEXgNmU8JWVoGo1oj",
	"38wCg6dBM7oMmQKMEHeNM2qPM1wtJKfRsS",
	"FTxppRPaBZcPtWZYfArrNJyWqcBHoL1Y1B",
	"2MJDmg9BWAxDPxMgJLpgQKGTPeAXJJ6v3c9",
	"28CAFDmMqmKSoxVLQjRvwcivL2QfMa6aXiz",
	"2Cw4ZpUwWguc9tyBwUH2LV8zZq3t59SU8bW",
	"2EypZrAnLKWd5KwC4LMdazYYKR1SWoaWTor",
	"x6p2yGvRpjHT9rJmpanHXdpTv4VmV6jc7Q",
	"2XDcmSZZqerTrj8LPTuFyTHV9VFVVYWts72",
	"5WPWqGUQPM2LJUZJu5f5ZSLSdYjZbGum6a",
	"5e8ziamKuhcd6i4BUzFPKAVbrQrW616bK9",
	"2mtbDQRzVoH3ymBpqw72xzd2Ez41N6PESNc",
	"2enACCZVAF5TFq87yBGJyFfNTEivKCZ8yC9",
	"hccUwHQmGCqNCfqAcPfi4j9kfoHsVAco2L",
	"EejGwSt3xqBndH4WYR6UHMFJAxLYvmemVx",
	"56FXtufSGQ2HY3SaS37XLKupFQAjunoo1d",
	"2cyg9Dnxx2uN3ABAco5p5kmazB2vo3by2Hk",
	"npn21JqxF54Ex3LmopiocnbkirCtFMvhfz",
	"2b3yDPM4ocboFY7DckzXAhY6MPeHYfxxgpG",
	"2bb1sC6ba1fMjGFDvD61Hqnmu46KjuRL5Ay",
	"25PdXP81WzCtZ2bFsiprRz8qpQS9S9V4Dzw",
	"qusHCLkCFxdMXKM7PPQrATsqZV2xfZ7Lkf",
	"vg8tua7XXFCRsu8yEKGR1yFJo7He5oupTf",
	"2cWR6UcfFNkktbDbpWVQY1csnmyS7cBQ7Ln",
}
