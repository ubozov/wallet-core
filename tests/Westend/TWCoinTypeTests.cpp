// Copyright © 2017-2020 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.
//
// This is a GENERATED FILE, changes made here MAY BE LOST.
// Generated one-time (codegen/bin/cointests)
//

#include "../interface/TWTestUtilities.h"
#include <TrustWalletCore/TWCoinTypeConfiguration.h>
#include <gtest/gtest.h>


TEST(TWWestendCoinType, TWCoinType) {
    auto symbol = WRAPS(TWCoinTypeConfigurationGetSymbol(TWCoinTypeWestend));

    auto txId = WRAPS(TWStringCreateWithUTF8Bytes("0x6f71c6f13d6254fbc326f50301aff4fd1eebfc54f6e7e44a59c0349f0dc833a4"));
    auto txUrl = WRAPS(TWCoinTypeConfigurationGetTransactionURL(TWCoinTypeWestend, txId.get()));
    auto accId = WRAPS(TWStringCreateWithUTF8Bytes("5FmkdVBGFFH8Fy7BBTwfPHddgm3wAc7GEjqcEFLZgQeHWBCx"));
    auto accUrl = WRAPS(TWCoinTypeConfigurationGetAccountURL(TWCoinTypeWestend, accId.get()));
    auto id = WRAPS(TWCoinTypeConfigurationGetID(TWCoinTypeWestend));
    auto name = WRAPS(TWCoinTypeConfigurationGetName(TWCoinTypeWestend));

    ASSERT_EQ(TWCoinTypeConfigurationGetDecimals(TWCoinTypeWestend), 10);
    ASSERT_EQ(TWBlockchainPolkadot, TWCoinTypeBlockchain(TWCoinTypeWestend));
    ASSERT_EQ(0x0, TWCoinTypeP2shPrefix(TWCoinTypeWestend));
    ASSERT_EQ(0x0, TWCoinTypeStaticPrefix(TWCoinTypeWestend));
    assertStringsEqual(symbol, "WND");
    assertStringsEqual(txUrl, "https://westend.subscan.io/extrinsic/0x6f71c6f13d6254fbc326f50301aff4fd1eebfc54f6e7e44a59c0349f0dc833a4");
    assertStringsEqual(accUrl, "https://westend.subscan.io/account/5FmkdVBGFFH8Fy7BBTwfPHddgm3wAc7GEjqcEFLZgQeHWBCx");
    assertStringsEqual(id, "westend");
    assertStringsEqual(name, "Westend");
}
