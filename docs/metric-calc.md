# Metrics calculation

Sgx Wallet Status: Returns SGX server connection status and SGX, got result by calling `/status/sgx` api,considered result field is `Status`.

Sgx Wallet version: Returns SGX wallet version, result got from `status/sgx` api, considered result field is `SgxWalletVersion`

BTRFS status: Returns information about btrfs kernel module whether it is enabled or disabled, result got by calling `/status/btrfs` considered result field is `KernelModule`.

Hardware: Returns node hardware information, result got by calling `/status/hardware` api endpoint, the info includes 
- cpu_total_cores - The number of logical CPUs in the system(the number of physical cores multiplied by the number of threads taht can run on each other)
- cpu_physical_cores - The number of physical CPUs in the system
- memory - Total physical memory
- swap - total swap memory
- system_release - system/OS name and system's release
- uname_version - system's release version
- attached_storae_size - attached storage size

IMA Status: Returns the status of the IMA container whether the IMA docker container is running or not, will get result from api `/status/ima`.

Block Number: Returns info on ethereum node endpoint, used by a given SKALE node, returns current block number and syncing status, result got from `/status/endpoint` api, considered result field is `BlockNumber`.
   
Containers Health: Returns health status of all  SKALE containers running on a given node. Here's a list of the base containers:
   - skale_transaction-manager
   - skale_admin
   - skale_api
   - skale_mysql
   - skale_sla
   - skale_bounty
   - skale_watchdog
   - skale_nginx

Wallet Address: Returns address of the wallet, will get result by executing node cli wallet command `skale wallet info`, result field is `Address`.

ETH Balance: Returns ETH account balance, will get result by executing node cli wallet command `skale wallet info`, result field is `EthBalanceWei`.

SKALE Balance: Returns Skale account balance, will get result by executing node cli wallet command `skale wallet info`, result field is `SkaleBalanceWei`.







