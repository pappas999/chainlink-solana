import { Result } from '@chainlink/gauntlet-core'
import { logger, prompt } from '@chainlink/gauntlet-core/dist/utils'
import { SolanaCommand, TransactionResponse, RawTransaction } from '@chainlink/gauntlet-solana'
import { AccountMeta, PublicKey } from '@solana/web3.js'
import { CONTRACT_LIST, getContract } from '../../../lib/contracts'
import { SolanaConstructor } from '../../../lib/types'
import { makeTx } from '../../../lib/utils'

export const makeTransferOwnershipCommand = (contractId: CONTRACT_LIST): SolanaConstructor => {
  return class TransferOwnership extends SolanaCommand {
    static id = `${contractId}:transfer_ownership`
    static category = contractId

    static examples = [
      `yarn gauntlet ${contractId}:transfer_ownership --network=devnet --state=[PROGRAM_STATE] --to=[PROPOSED_OWNER]`,
    ]

    constructor(flags, args) {
      super(flags, args)

      this.require(!!this.flags.state, 'Please provide flags with "state"')
      this.require(!!this.flags.to, 'Please provide flags with "to"')
    }

    makeRawTransaction = async (signer: PublicKey) => {
      const contract = getContract(contractId, '')
      const address = contract.programId.toString()
      const program = this.loadProgram(contract.idl, address)

      const state = new PublicKey(this.flags.state)
      const proposedOwner = new PublicKey(this.flags.to)

      const data = program.coder.instruction.encode('transfer_ownership', {
        proposedOwner,
      })

      const accounts: AccountMeta[] = [
        {
          pubkey: state,
          isSigner: false,
          isWritable: true,
        },
        {
          pubkey: signer,
          isSigner: true,
          isWritable: false,
        },
      ]

      const rawTx: RawTransaction = {
        data,
        accounts,
        programId: contract.programId,
      }

      return [rawTx]
    }

    execute = async () => {
      const contract = getContract(contractId, '')
      const rawTx = await this.makeRawTransaction(this.wallet.payer.publicKey)
      const tx = makeTx(rawTx)
      await prompt(
        `Transfering ownership of ${contractId} state (${this.flags.state.toString()}) to: (${new PublicKey(
          this.flags.to,
        ).toString()}). Continue?`,
      )
      logger.debug(tx)
      logger.loading('Sending tx...')
      const txhash = await this.sendTx(tx, [this.wallet.payer], contract.idl)
      logger.success(`Ownership transferred to ${new PublicKey(this.flags.to)} on tx ${txhash}`)
      return {
        responses: [
          {
            tx: this.wrapResponse(txhash, this.flags.state),
            contract: this.flags.state,
          },
        ],
      } as Result<TransactionResponse>
    }
  }
}
