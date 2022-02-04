package api

import (
	"time"
)

const ProtocolConfigMethod = "EXPERIMENTAL_protocol_config"

type ProtocolConfigResult struct {
	ProtocolVersion                 int       `json:"protocol_version"`
	GenesisTime                     time.Time `json:"genesis_time"`
	ChainID                         string    `json:"chain_id"`
	GenesisHeight                   int       `json:"genesis_height"`
	NumBlockProducerSeats           int       `json:"num_block_producer_seats"`
	NumBlockProducerSeatsPerShard   []int     `json:"num_block_producer_seats_per_shard"`
	AvgHiddenValidatorSeatsPerShard []int     `json:"avg_hidden_validator_seats_per_shard"`
	DynamicResharding               bool      `json:"dynamic_resharding"`
	ProtocolUpgradeStakeThreshold   []int     `json:"protocol_upgrade_stake_threshold"`
	EpochLength                     int       `json:"epoch_length"`
	GasLimit                        int64     `json:"gas_limit"`
	MinGasPrice                     string    `json:"min_gas_price"`
	MaxGasPrice                     string    `json:"max_gas_price"`
	BlockProducerKickoutThreshold   int       `json:"block_producer_kickout_threshold"`
	ChunkProducerKickoutThreshold   int       `json:"chunk_producer_kickout_threshold"`
	OnlineMinThreshold              []int     `json:"online_min_threshold"`
	OnlineMaxThreshold              []int     `json:"online_max_threshold"`
	GasPriceAdjustmentRate          []int     `json:"gas_price_adjustment_rate"`
	RuntimeConfig                   struct {
		StorageAmountPerByte string `json:"storage_amount_per_byte"`
		TransactionCosts     struct {
			ActionReceiptCreationConfig struct {
				SendSir    int64 `json:"send_sir"`
				SendNotSir int64 `json:"send_not_sir"`
				Execution  int64 `json:"execution"`
			} `json:"action_receipt_creation_config"`
			DataReceiptCreationConfig struct {
				BaseCost struct {
					SendSir    int64 `json:"send_sir"`
					SendNotSir int64 `json:"send_not_sir"`
					Execution  int64 `json:"execution"`
				} `json:"base_cost"`
				CostPerByte struct {
					SendSir    int `json:"send_sir"`
					SendNotSir int `json:"send_not_sir"`
					Execution  int `json:"execution"`
				} `json:"cost_per_byte"`
			} `json:"data_receipt_creation_config"`
			ActionCreationConfig struct {
				CreateAccountCost struct {
					SendSir    int64 `json:"send_sir"`
					SendNotSir int64 `json:"send_not_sir"`
					Execution  int64 `json:"execution"`
				} `json:"create_account_cost"`
				DeployContractCost struct {
					SendSir    int64 `json:"send_sir"`
					SendNotSir int64 `json:"send_not_sir"`
					Execution  int64 `json:"execution"`
				} `json:"deploy_contract_cost"`
				DeployContractCostPerByte struct {
					SendSir    int `json:"send_sir"`
					SendNotSir int `json:"send_not_sir"`
					Execution  int `json:"execution"`
				} `json:"deploy_contract_cost_per_byte"`
				FunctionCallCost struct {
					SendSir    int64 `json:"send_sir"`
					SendNotSir int64 `json:"send_not_sir"`
					Execution  int64 `json:"execution"`
				} `json:"function_call_cost"`
				FunctionCallCostPerByte struct {
					SendSir    int `json:"send_sir"`
					SendNotSir int `json:"send_not_sir"`
					Execution  int `json:"execution"`
				} `json:"function_call_cost_per_byte"`
				TransferCost struct {
					SendSir    int64 `json:"send_sir"`
					SendNotSir int64 `json:"send_not_sir"`
					Execution  int64 `json:"execution"`
				} `json:"transfer_cost"`
				StakeCost struct {
					SendSir    int64 `json:"send_sir"`
					SendNotSir int64 `json:"send_not_sir"`
					Execution  int64 `json:"execution"`
				} `json:"stake_cost"`
				AddKeyCost struct {
					FullAccessCost struct {
						SendSir    int64 `json:"send_sir"`
						SendNotSir int64 `json:"send_not_sir"`
						Execution  int64 `json:"execution"`
					} `json:"full_access_cost"`
					FunctionCallCost struct {
						SendSir    int64 `json:"send_sir"`
						SendNotSir int64 `json:"send_not_sir"`
						Execution  int64 `json:"execution"`
					} `json:"function_call_cost"`
					FunctionCallCostPerByte struct {
						SendSir    int `json:"send_sir"`
						SendNotSir int `json:"send_not_sir"`
						Execution  int `json:"execution"`
					} `json:"function_call_cost_per_byte"`
				} `json:"add_key_cost"`
				DeleteKeyCost struct {
					SendSir    int64 `json:"send_sir"`
					SendNotSir int64 `json:"send_not_sir"`
					Execution  int64 `json:"execution"`
				} `json:"delete_key_cost"`
				DeleteAccountCost struct {
					SendSir    int64 `json:"send_sir"`
					SendNotSir int64 `json:"send_not_sir"`
					Execution  int64 `json:"execution"`
				} `json:"delete_account_cost"`
			} `json:"action_creation_config"`
			StorageUsageConfig struct {
				NumBytesAccount     int `json:"num_bytes_account"`
				NumExtraBytesRecord int `json:"num_extra_bytes_record"`
			} `json:"storage_usage_config"`
			BurntGasReward                    []int `json:"burnt_gas_reward"`
			PessimisticGasPriceInflationRatio []int `json:"pessimistic_gas_price_inflation_ratio"`
		} `json:"transaction_costs"`
		WasmConfig struct {
			ExtCosts struct {
				Base                        int   `json:"base"`
				ContractCompileBase         int   `json:"contract_compile_base"`
				ContractCompileBytes        int   `json:"contract_compile_bytes"`
				ReadMemoryBase              int64 `json:"read_memory_base"`
				ReadMemoryByte              int   `json:"read_memory_byte"`
				WriteMemoryBase             int64 `json:"write_memory_base"`
				WriteMemoryByte             int   `json:"write_memory_byte"`
				ReadRegisterBase            int64 `json:"read_register_base"`
				ReadRegisterByte            int   `json:"read_register_byte"`
				WriteRegisterBase           int64 `json:"write_register_base"`
				WriteRegisterByte           int   `json:"write_register_byte"`
				Utf8DecodingBase            int64 `json:"utf8_decoding_base"`
				Utf8DecodingByte            int   `json:"utf8_decoding_byte"`
				Utf16DecodingBase           int64 `json:"utf16_decoding_base"`
				Utf16DecodingByte           int   `json:"utf16_decoding_byte"`
				Sha256Base                  int64 `json:"sha256_base"`
				Sha256Byte                  int   `json:"sha256_byte"`
				Keccak256Base               int64 `json:"keccak256_base"`
				Keccak256Byte               int   `json:"keccak256_byte"`
				Keccak512Base               int64 `json:"keccak512_base"`
				Keccak512Byte               int   `json:"keccak512_byte"`
				Ripemd160Base               int   `json:"ripemd160_base"`
				Ripemd160Block              int   `json:"ripemd160_block"`
				EcrecoverBase               int64 `json:"ecrecover_base"`
				LogBase                     int64 `json:"log_base"`
				LogByte                     int   `json:"log_byte"`
				StorageWriteBase            int64 `json:"storage_write_base"`
				StorageWriteKeyByte         int   `json:"storage_write_key_byte"`
				StorageWriteValueByte       int   `json:"storage_write_value_byte"`
				StorageWriteEvictedByte     int   `json:"storage_write_evicted_byte"`
				StorageReadBase             int64 `json:"storage_read_base"`
				StorageReadKeyByte          int   `json:"storage_read_key_byte"`
				StorageReadValueByte        int   `json:"storage_read_value_byte"`
				StorageRemoveBase           int64 `json:"storage_remove_base"`
				StorageRemoveKeyByte        int   `json:"storage_remove_key_byte"`
				StorageRemoveRetValueByte   int   `json:"storage_remove_ret_value_byte"`
				StorageHasKeyBase           int64 `json:"storage_has_key_base"`
				StorageHasKeyByte           int   `json:"storage_has_key_byte"`
				StorageIterCreatePrefixBase int   `json:"storage_iter_create_prefix_base"`
				StorageIterCreatePrefixByte int   `json:"storage_iter_create_prefix_byte"`
				StorageIterCreateRangeBase  int   `json:"storage_iter_create_range_base"`
				StorageIterCreateFromByte   int   `json:"storage_iter_create_from_byte"`
				StorageIterCreateToByte     int   `json:"storage_iter_create_to_byte"`
				StorageIterNextBase         int   `json:"storage_iter_next_base"`
				StorageIterNextKeyByte      int   `json:"storage_iter_next_key_byte"`
				StorageIterNextValueByte    int   `json:"storage_iter_next_value_byte"`
				TouchingTrieNode            int64 `json:"touching_trie_node"`
				PromiseAndBase              int   `json:"promise_and_base"`
				PromiseAndPerPromise        int   `json:"promise_and_per_promise"`
				PromiseReturn               int   `json:"promise_return"`
				ValidatorStakeBase          int64 `json:"validator_stake_base"`
				ValidatorTotalStakeBase     int64 `json:"validator_total_stake_base"`
			} `json:"ext_costs"`
			GrowMemCost   int `json:"grow_mem_cost"`
			RegularOpCost int `json:"regular_op_cost"`
			LimitConfig   struct {
				MaxGasBurnt                      int64 `json:"max_gas_burnt"`
				MaxStackHeight                   int   `json:"max_stack_height"`
				StackLimiterVersion              int   `json:"stack_limiter_version"`
				InitialMemoryPages               int   `json:"initial_memory_pages"`
				MaxMemoryPages                   int   `json:"max_memory_pages"`
				RegistersMemoryLimit             int   `json:"registers_memory_limit"`
				MaxRegisterSize                  int   `json:"max_register_size"`
				MaxNumberRegisters               int   `json:"max_number_registers"`
				MaxNumberLogs                    int   `json:"max_number_logs"`
				MaxTotalLogLength                int   `json:"max_total_log_length"`
				MaxTotalPrepaidGas               int64 `json:"max_total_prepaid_gas"`
				MaxActionsPerReceipt             int   `json:"max_actions_per_receipt"`
				MaxNumberBytesMethodNames        int   `json:"max_number_bytes_method_names"`
				MaxLengthMethodName              int   `json:"max_length_method_name"`
				MaxArgumentsLength               int   `json:"max_arguments_length"`
				MaxLengthReturnedData            int   `json:"max_length_returned_data"`
				MaxContractSize                  int   `json:"max_contract_size"`
				MaxTransactionSize               int   `json:"max_transaction_size"`
				MaxLengthStorageKey              int   `json:"max_length_storage_key"`
				MaxLengthStorageValue            int   `json:"max_length_storage_value"`
				MaxPromisesPerFunctionCallAction int   `json:"max_promises_per_function_call_action"`
				MaxNumberInputDataDependencies   int   `json:"max_number_input_data_dependencies"`
				MaxFunctionsNumberPerContract    int   `json:"max_functions_number_per_contract"`
			} `json:"limit_config"`
		} `json:"wasm_config"`
		AccountCreationConfig struct {
			MinAllowedTopLevelAccountLength int    `json:"min_allowed_top_level_account_length"`
			RegistrarAccountID              string `json:"registrar_account_id"`
		} `json:"account_creation_config"`
	} `json:"runtime_config"`
	TransactionValidityPeriod int    `json:"transaction_validity_period"`
	ProtocolRewardRate        []int  `json:"protocol_reward_rate"`
	MaxInflationRate          []int  `json:"max_inflation_rate"`
	NumBlocksPerYear          int    `json:"num_blocks_per_year"`
	ProtocolTreasuryAccount   string `json:"protocol_treasury_account"`
	FishermenThreshold        string `json:"fishermen_threshold"`
	MinimumStakeDivisor       int    `json:"minimum_stake_divisor"`
}

func (c *Client) ProtocolConfig() (*ProtocolConfigResult, error) {
	var result *ProtocolConfigResult
	err := c.do(ProtocolConfigMethod, map[string]string{"finality": "final"}, &result)

	return result, err
}
