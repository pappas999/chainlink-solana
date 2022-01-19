import { Middleware, Next } from '@chainlink/gauntlet-core'
import { assertions } from '@chainlink/gauntlet-core/dist/utils'
import { Provider } from '@project-serum/anchor'
import { Connection, ConnectionConfig, Keypair } from '@solana/web3.js'
import SolanaCommand from './internal/solana'
const { Wallet } = require('@project-serum/anchor') // module exported dynamically from anchor

const isValidURL = (url: string) => {
  var pattern = new RegExp('^(https?)://')
  return pattern.test(url)
}
export const withProvider: Middleware = (c: SolanaCommand, next: Next) => {
  const nodeURL = process.env.NODE_URL
  assertions.assert(
    nodeURL && isValidURL(nodeURL),
    `Invalid NODE_URL (${nodeURL}), please add an http:// or https:// prefix`,
  )

  let connectionConfig: ConnectionConfig = {}
  if (process.env.CONFIRM_TX_TIMEOUT_SECONDS) {
    connectionConfig.confirmTransactionInitialTimeout = parseInt(process.env.CONFIRM_TX_TIMEOUT_SECONDS) * 1000
  }
  c.provider = new Provider(new Connection(nodeURL, connectionConfig), c.wallet, {})
  return next()
}

export const withWallet: Middleware = (c: SolanaCommand, next: Next) => {
  const rawPK = process.env.PRIVATE_KEY
  assertions.assert(!!rawPK, `Missing PRIVATE_KEY, please add one`)

  c.wallet = new Wallet(Keypair.fromSecretKey(Uint8Array.from(JSON.parse(rawPK))))
  console.info(`Operator address is ${c.wallet.publicKey}`)
  return next()
}

export const withNetwork: Middleware = (c: SolanaCommand, next: Next) => {
  assertions.assert(
    !!c.flags.network,
    `Network required. Invalid network (${c.flags.network}), please specify a --network`,
  )
  return next()
}
