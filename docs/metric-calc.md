# Metrics calculation

Sgx Status: Returns SGX server info - connection status and SGX wallet version, got result by calling `/status/sgx` api considered result field is `Status`.

Hardware: Returns node hardware information result got from by calling `/status/hardware` api endpoint, the info includes 
- cpu_total_cores - The number of logical CPUs in the system(the number of physical cores multiplied by the number of threads taht can run on each other)
- cpu_physical_cores - The number of physical CPUs in the system
- memory - Total physical memory in bytes
- swap - total swap memory in bytes
- system_release - system/OS name and system's release
- uname_version - system's release version
- attached_storae_size - attached storage size in bytes 

IMA Status: Returns the status of the IMA container whether the IMA docker container is running or not, will get result from api `/status/ima`

Endpoint status: Returns info on ethereum node endpoint, used by a given SKALE node returns current block number and syncing status, result got from by calling `/status/endpoint` api. 
   
public IP : Returns Public IP of the node will get result from `/status/public-ip` api

schain status : Returns list of health checks for every sChain running on a given node, result got by calling `/status/schains` 

