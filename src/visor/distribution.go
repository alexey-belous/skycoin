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
	"2FMAnmCfWT6C3Nz1ugsJM1wvgVvnNg3fcTM",
	"xWURH7p1Gyg25Eja2xyBqBn13HY6ZNpYAu",
	"BxJSfyHgzwPt4gsDU2zkfavAdnp69yGNQ9",
	"K2irXY1G3bj9YBA1XwNRQZXEPk9qWfbRaL",
	"SiimJ8cis8xZfbst1yN5amkisbGuAYRsom",
	"26Yd2aytrvRHyJ9b36CcCkfnhZKV93SW6RY",
	"h4Te2G2ZWpeXJzzv7uwgyvRdqXt7z2Hv8d",
	"itXS5xTNz2hFECLSFLqvpdD1L37WmK8zrZ",
	"2bocnkRssxqKJ4K4P4gKErTDzAiLyDjBKfp",
	"2LWcfesczSi7aqpnqc1hLTBdr64Szp87Rzr",
	"2Tke7r9p9ZkGK9E6SKn7Qx6Gytdf2jBQUo3",
	"S99fBJkZGnScocoet1s9FQWsE28hu7Kzgk",
	"aaopiZpcqPgPpctP6XEH8S3DpbwocEHsV4",
	"pCXSzTwMiG6PGdhErhf9CkR8dgXVyEXLA5",
	"24GRv5KhPLb1VwgFFgDSx5rFiP59c9HVBs4",
	"S1hbbv8SJc2V998aWziPAqfmyQ92AccNbb",
	"KYQKH6wMayZ5CSGoyUtN8xcFGWDYUvCDfB",
	"2Pog4cZeDP6xcci1zJoKqwcXaaVp8HNVovd",
	"zZ4DbpumZGvRC11hWzsLnNk7LjJfFPbMW4",
	"2iXdrDqoKUFpuQfa47wVdsQaSvaW28A3Pis",
	"2Hz2nvbnpsTMnctYxLRxmrNXw9KCBSgh6kP",
	"tQ6N121TqGNWkdUBiSa533LLmvRwHJSZ8e",
	"UEfHBx4u7j1ZY26tWqq3VXYsCnM3Un8g4J",
	"YXYDWyaswkna8cvdbsY5jYSDnRLMG13rmu",
	"K6U1s9GURK9MzNdj86qmej6KUwsneQj5Tz",
	"kDa6bi7bp4UiWJaC6544c2SVXAdgPhtYSB",
	"t396ZwGkD1QjDryv1XxSbv2abytvGtoKCN",
	"22dSUitosU8uzG5Bn8DFYpNxJDKeQ3RXvwK",
	"2CHpG3GLMdvMNv6Z3e6BFS7YWfF2P9TDD1a",
	"hTF49TJQm3ND1znHHtnMKyc9wuAng58XRW",
	"2iE5QGc9HDnEwmu8amkRLWmS8UvwgTXHJc2",
	"iJfW9h8PmuKQp7UtoQZVQbhWULCpFVYZgA",
	"28kv43eTY2y9pJRBH4ykQidqo6Wt1nJXw26",
	"Fgne8xpUcjuvv8S4jvDksWdSYfdmbxGU9C",
	"2ei8FrQHSmRASfPfztPZgffDn4xjwdLuwP1",
	"2CB7MdJSpLXK2YnSZL3jyGWskyuR8tJ5sgY",
	"2gjx2qAUHcJ6n98pM3bPgEEbeLGwNaxpJ1z",
	"c5W16hcDxcCRE3zrgh5imddGV9hBH1q4MH",
	"VERz6cDAdR3jjUiGxFX4U7f1kuDEcJbt9Q",
	"rsgbTQ2rNDDJXx53godpynf2rJsiP591iz",
	"2juGKvQBxM3XnesSRy51ivTqutnTxWNk7jg",
	"2jp1gLLo73sFy1c3nFQS2EGz2yj9HPzH61o",
	"2WjQcjWXk4QeqTT6c7Gkg14UzNiPVJfWzFu",
	"DCYiANQxoRfZboGz9RCHG4RHRCDmRmzxXG",
	"2NSr9vhvdY3bjmBk1v9dB2EFsChN5EiDPDq",
	"Z5fpGDy4V43dvQ52Nmc1Jr9G6Z4xFGJSKj",
	"27s1sioP1HHh9hT1F7WXXgkvwYJ1qVf5Loo",
	"UnL18qGvQbrokU4LntEDFA8if9FZ1GtCNs",
	"ZfuirMBiMYLfA4HrM87VDkZHBVa89EjUGt",
	"2BALkePPQ1j1BcXRJtR9sMnLvxELzuGdgHc",
	"atnzZjNetMTtppmBqXqKivi2zK5BNcxjFc",
	"24vJJKFNx8kcSYPkHb1k4VAaMKZAF1JnTBW",
	"2ZcC9y2FxzVE6wmk3U1xwWATxxHG4vSnrU9",
	"ZGqhDxb78zoNCU1KATkxgBr5zK84LXYu3M",
	"2j6BEAcXJTcMz5y174KRXzPeMLUP6hw3YpB",
	"S12uE8hNDoTKRwyJb4sPH4okqe547CvVaP",
	"NLCNZ9yd1sFsmwrdpGqjqqB3Hia5JuGqRZ",
	"2fsRNVLd3KahUHC2bcEs3qPQTHnrP3LequJ",
	"2aFQoo7mDyrgYkoDuNfbJcsFbiPZNhNAgAW",
	"2RkVNmiZDg7wCkQmpQKTNC2C3RrAZDWhCRE",
	"2TCkcawz7AFtqhWNCyR4VBh1GwnAXQnDpsV",
	"neqcQtfuprmSV5eyJrCrManXsnNHD5xw4P",
	"qektXc2JnNJYD4cALYVuy57cEvDCVTVV2C",
	"QzxzvRvCc8FWCygTNrL8z5HTxuQxT8cC1D",
	"nPUkE62xAC16RiF4adtL5735N8yWN8qc6d",
	"55KmMPYgeVzAVJauj4Eu7n1NQdU9pbRp3m",
	"jbGhUFsDad3qL1HbBhwppi2BzAFVqVYnK8",
	"2UEqUpuJTcFGeN2REd1fAkhashmpyoEwFhy",
	"26K8acZE6Y9xRG93Ks1J8kcr8bbkV5qnXFm",
	"fLGv5TamodoT4psFjuf1npUq6tzUqEFjEE",
	"2gGVSskaixBzYsaYiNCT26kRP735PoyqjuU",
	"2Mj6oippmY9CaJHcLDHLiYV5YuT9NP48Lxr",
	"VuaCrcMfu4tGo8FMmWRk8NLQhMLogbQc8",
	"KXX2pu8ifAQNhB9KqHVKQNNW52jUvPKHRe",
	"22TMZLuLcJHDFgdqCTbNWPa1rz7V1vFgirr",
	"f5xQMPhdSYkqRFRgzrmSwCDGqa1WRPtgUr",
	"ereKARSFUui9ZUH82zk1dvNWEYW3T75rYs",
	"5pppbMfhSNJ7TLRbQMNwqTcQvDBavsQ1gv",
	"XaFwVLHQWKkcs5Fb1ehicZB56stnjhdrdB",
	"Vp9cC2uZPqwJriYJutuwkZmYz4uSZATyqC",
	"2jyGSptXPcEuidUmuzLwoLiTAG9Qpob2AFW",
	"sAVpR959xRAfx98LqHatuahJs3bA1xGBFr",
	"TdAP6LzTjHkJaDdJWQYWYqfxDkvz1AfYNX",
	"N2GugLK4U5TwRP9jrqXX3RZYNkbGoDMNPv",
	"EnvDi8ktHBe8H4y9DuBxfoxbqPtGwn3zZz",
	"nxncHmGZMG8Dx4SwsiGwhbZwFRUXFgqs3P",
	"gyHNeH78bvJBJfR1694RpVQrNah6PVRSoc",
	"PmS9kFWUxdhtPEYFDSFHLJYUnTHGRV1Po3",
	"rru4VSgQra3kTeK4emJqHH4BXroMyGonJF",
	"MzXcc9bC1hG9pvEM1QfHGKCUWJiS7wtSbu",
	"76sScDnrq4JUw7CKzaazKZqvzivpkCeBKZ",
	"2Xo3d7AU4ffQSC8VGX3bR4qxQA14Bg4KfLA",
	"2LvRjcCHjsvcWVk8rWb64ypBvcL6K9CTBtf",
	"Hh2Ljppp1X32AC3K7s6jGUM1DJJEJEMxHH",
	"VwEXfiB43aCFYLVuQPyM5MM4LXUNDpKn9r",
	"WNxD54DXBWALbNKEZYKpA1QzjBnTaCfKLg",
	"RtrqWBC6MeJrvYJLu9ubqU4noDgYhSEga2",
	"LAM8zS9nZVLcmi2ve6SgTCBB3U5VKMmMnY",
	"29WuwfJFMyvrjxdPaVS16DcbPvQ7rEt9mfU",
	"2maZ2vGBzcqLbjvuhH6xdJF6RRs2on8kPP7",
}
