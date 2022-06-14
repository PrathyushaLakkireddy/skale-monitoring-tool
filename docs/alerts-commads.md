## Alerting (Telegram and Email)
A custom alerting module has been developed to alert on key validator health events.
The module use data from prometheus and triggers alerts based on user-configured thresholds.

Here are the list of Alerts
- Alert when container health is not **Running**, **Paused** or **Dead**
- Alert when ETH balance has dropped below **ethbalance_change_threshold** which is user configured in *config.toml*
- Alert when Skale balance has dropped below **sklbalance_change_threshold** which is user configured in *config.toml*
- Alert when ETH balance has dropped below or equal to 1.5ETH which is compilance alert to maintain minimum balance in wallet.
- Alert when node health is not OK
- Alert when SGX wallet is not connected 

## Telegram Commands
These Commands can be used to get quick information about your Skale node

Here is the list of available Telegram Commands.
  - **/list** - list out the available commands
  - **/eth_balance** - returns ETH account balance
  - **/skale_balance** - returns Skale balance
  - **/block_numner** - returns current block number
  - **/sgx_status** - status of the sgx server whether it is CONNECTED or not
  - **/container_status** - returns status of the container's state of health which are running in skale node.
  - **/stop** - which panics the running code and also alerts will be stopped.

