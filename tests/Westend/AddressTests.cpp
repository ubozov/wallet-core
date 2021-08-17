// Copyright © 2017-2020 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#include "HexCoding.h"
#include "Westend/Address.h"
#include "PublicKey.h"
#include "PrivateKey.h"
#include <gtest/gtest.h>
#include <vector>

using namespace TW;
using namespace TW::Westend;

TEST(WestendAddress, Validation) {
    // Bitcoin
    ASSERT_FALSE(Address::isValid("1ES14c7qLb5CYhLMUekctxLgc1FV2Ti9DA"));
    // Kusama ed25519
    ASSERT_FALSE(Address::isValid("FHKAe66mnbk8ke8zVWE9hFVFrJN1mprFPVmD5rrevotkcDZ"));
    // Kusama secp256k1
    ASSERT_FALSE(Address::isValid("FxQFyTorsjVsjjMyjdgq8w5vGx8LiA1qhWbRYcFijxKKchx"));
    // Kusama sr25519
    ASSERT_FALSE(Address::isValid("EJ5UJ12GShfh7EWrcNZFLiYU79oogdtXFUuDDZzk7Wb2vCe"));
    // Polkadot ed25519
    ASSERT_FALSE(Address::isValid("15KRsCq9LLNmCxNFhGk55s5bEyazKefunDxUH24GFZwsTxyu"));
    // Polkadot sr25519
    ASSERT_FALSE(Address::isValid("15AeCjMpcSt3Fwa47jJBd7JzQ395Kr2cuyF5Zp4UBf1g9ony"));
    
    // Westend ed25519
    ASSERT_TRUE(Address::isValid("5FmkdVBGFFH8Fy7BBTwfPHddgm3wAc7GEjqcEFLZgQeHWBCx"));
    // Westend sr25519
    ASSERT_TRUE(Address::isValid("5DtCbNMGwhnP5wJ25Zv59wc5aj5uo3wYdr8536qSRxbvmLdK"));
}

TEST(WestendAddress, FromPrivateKey) {
    auto privateKey = PrivateKey(parse_hex("0x6328b13f07b70107c70d94fd9cd6ed9d7394d94916362f82d90b9da9fad2e58a"));
    auto address = Address(privateKey.getPublicKey(TWPublicKeyTypeED25519));
    ASSERT_EQ(address.string(), "5DJZrQ3ZLAJATK9b8cb9ePRnaT6GJA1yEG6db9renscscA8b");
}

TEST(WestendAddress, FromPublicKey) {
    auto publicKey = PublicKey(parse_hex("0x182a8385912e206910a11dfcfbadcc018d73ff3febe36f9eb85774e77591a314"), TWPublicKeyTypeED25519);
    auto address = Address(publicKey);
    ASSERT_EQ(address.string(), "5CcPdpUAhRvGjyxcpnL7emi7SruQ98eP7mC8cDNkjCKfXNZ7");
}

TEST(WestendAddress, FromString) {
    auto address = Address("5DtCbNMGwhnP5wJ25Zv59wc5aj5uo3wYdr8536qSRxbvmLdK");
    ASSERT_EQ(address.string(), "5DtCbNMGwhnP5wJ25Zv59wc5aj5uo3wYdr8536qSRxbvmLdK");
}
