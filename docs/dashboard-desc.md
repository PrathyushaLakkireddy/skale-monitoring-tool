## Validator Monitoring Metrics
The following list of metrics are displayed in this dashboard

*Block Number* : Current Block Number

*SGX Wallet Status* : Connection status of the SGX wallet whether it is conncected or not.

*SGX Wallet Version* : SGX wallet version 

*BTRFS kernal module status* : Status of the BTRFS kernal module whether it is enabled or disabled.

*Containers Health* : Displays status health of the all SKALE containers. Here's a list of the base containers:
   - skale_transaction-manager
   - skale_admin
   - skale_api
   - skale_mysql
   - skale_sla
   - skale_bounty
   - skale_watchdog
   - skale_nginx

*Hardware Info*: Displays node hardware information, which incudes
- cpu_total_cores - The number of logical CPUs in the system(the number of physical cores multiplied by the number of threads taht can run on each other)
- cpu_physical_cores - The number of physical CPUs in the system
- memory - Total physical memory
- swap - total swap memory
- system_release - system/OS name and system's release
- uname_version - system's release version
- attached_storae_size - attached storage size
