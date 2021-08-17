// Copyright © 2017-2020 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#include "Polkadot/Signer.h"
#include "Polkadot/Extrinsic.h"
#include "Polkadot/Address.h"
#include "SS58Address.h"
#include "HexCoding.h"
#include "PrivateKey.h"
#include "PublicKey.h"
#include "proto/Polkadot.pb.h"
#include "uint256.h"

#include <TrustWalletCore/TWSS58AddressType.h>
#include <gtest/gtest.h>


namespace TW::Polkadot {
    extern PrivateKey privateKey;
    extern PublicKey toPublicKey;
    auto genesisHashWND = parse_hex("0xe143f23803ac50e8f6f8e62695d1ce9e4e1d68aa36c1cd2cfd15340213f3423e");

TEST(PolkadotSigner, SignTransferWND) {

    auto blockHash = parse_hex("0x343a3f4258fd92f5ca6ca5abdf473d86a78b0bcd0dc09c568ca594245cc8c642");
    auto toAddress = SS58Address(toPublicKey, TWSS58AddressTypeWestend);

    auto input = Proto::SigningInput();
    input.set_genesis_hash(genesisHashWND.data(), genesisHashWND.size());
    input.set_block_hash(blockHash.data(), blockHash.size());

    input.set_nonce(0);
    input.set_spec_version(17);
    input.set_private_key(privateKey.bytes.data(), privateKey.bytes.size());
    input.set_network(Proto::Network::WESTEND);
    input.set_transaction_version(3);

    auto &era = *input.mutable_era();
    era.set_block_number(927699);
    era.set_period(8);

    auto balanceCall = input.mutable_balance_call();
    auto &transfer = *balanceCall->mutable_transfer();
    auto value = store(uint256_t(12345));
    transfer.set_to_address(toAddress.string());
    transfer.set_value(value.data(), value.size());

    auto extrinsic = Extrinsic(input);
    auto preimage = extrinsic.encodePayload();
    auto output = Signer::sign(input);

    ASSERT_EQ(hex(preimage), "05008eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48e5c0320000001100000003000000e143f23803ac50e8f6f8e62695d1ce9e4e1d68aa36c1cd2cfd15340213f3423e343a3f4258fd92f5ca6ca5abdf473d86a78b0bcd0dc09c568ca594245cc8c642");
    ASSERT_EQ(hex(output.encoded()), "29028488dc3417d5058ec4b4503e0c12ea1a0a89be200fe98922423d4334014fa6b0ee0019d1e371382b0953d4c35b1a0d24a9a761950fdeb7967bd4de5b05dce0d65751a050c73b1f5fbc94ca8f0850cd84fe81eda435416396a32c9ccb043898c5d60c3200000005008eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48e5c0");
}

} // namespace
