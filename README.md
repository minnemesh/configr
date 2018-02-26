# configr
A high-reliability configuration management tool designed for mesh networking and distributed systems.

Currently under development.

## Planned features
* Ability to load configuration files from multiple sources including HTTP(S) and IPFS.
* Configurations are encrypted to their destination nodes, allowing secure storage of secrets.
* Configurations are signed to ensure they come from a trusted source.
* Configurations are versioned to prevent replay attacks.
* Flexible fallback behavior including using the last known config, retries, and timeouts.
* Flexible output options (stdout, environment variables, write to file)
