// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgPlayMove } from "./types/checkers/tx";
import { MsgRejectGame } from "./types/checkers/tx";
import { MsgCreateGame } from "./types/checkers/tx";
import { MsgCreatePost } from "./types/checkers/tx";


const types = [
  ["/alice.checkers.checkers.MsgPlayMove", MsgPlayMove],
  ["/alice.checkers.checkers.MsgRejectGame", MsgRejectGame],
  ["/alice.checkers.checkers.MsgCreateGame", MsgCreateGame],
  ["/alice.checkers.checkers.MsgCreatePost", MsgCreatePost],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgPlayMove: (data: MsgPlayMove): EncodeObject => ({ typeUrl: "/alice.checkers.checkers.MsgPlayMove", value: MsgPlayMove.fromPartial( data ) }),
    msgRejectGame: (data: MsgRejectGame): EncodeObject => ({ typeUrl: "/alice.checkers.checkers.MsgRejectGame", value: MsgRejectGame.fromPartial( data ) }),
    msgCreateGame: (data: MsgCreateGame): EncodeObject => ({ typeUrl: "/alice.checkers.checkers.MsgCreateGame", value: MsgCreateGame.fromPartial( data ) }),
    msgCreatePost: (data: MsgCreatePost): EncodeObject => ({ typeUrl: "/alice.checkers.checkers.MsgCreatePost", value: MsgCreatePost.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
